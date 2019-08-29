package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type trailerReader struct {
	resp *http.Response
}

func (r *trailerReader) Read(b []byte) (int, error) {
	n, err := r.resp.Body.Read(b)
	if err != nil {
		return n, err
	}
	return n, nil
}

func (r *trailerReader) Close() error {
	return r.resp.Body.Close()
}

type Response struct {
	Output io.ReadCloser
	Error  *Error
}

func (r *Response) Close() error {
	if r.Output != nil {

		_, err1 := io.Copy(ioutil.Discard, r.Output)
		err2 := r.Output.Close()
		if err1 != nil {
			return err1
		}
		return err2
	}
	return nil
}

func (r *Response) Cancel() error {
	if r.Output != nil {
		return r.Output.Close()
	}

	return nil
}

func (r *Response) decode(dec interface{}) error {
	if r.Error != nil {
		return r.Error
	}
	err := json.NewDecoder(r.Output).Decode(dec)
	err2 := r.Close()
	if err != nil {
		return err
	}
	return err2
}

func (r *Request) Send(c *http.Client) (*Response, error) {
	url := r.getURL()
	req, err := http.NewRequest("POST", url, r.Body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(r.Ctx)
	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	contentType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	nresp := new(Response)
	nresp.Output = &trailerReader{resp}
	if resp.StatusCode >= http.StatusBadRequest {
		e := new(Error)
		switch {
		case resp.StatusCode == http.StatusNotFound:
			e.Message = "command not found"
		case contentType == "text/plain":
			out, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "filecoin: warning! response (%d) read error: %s\n", resp.StatusCode, err)
			}
			e.Message = string(out)
			switch resp.StatusCode {
			case http.StatusNotFound, http.StatusBadRequest:
				e.Code = ErrClient
			case http.StatusTooManyRequests:
				e.Code = ErrRateLimited
			case http.StatusForbidden:
				e.Code = ErrForbidden
			}
		case contentType == "application/json":
			if err = json.NewDecoder(resp.Body).Decode(e); err != nil {
				fmt.Fprintf(os.Stderr, "filecoin: warning! response (%d) unmarshall error: %s\n", resp.StatusCode, err)
			}
		default:
			e.Code = ErrImplementation
			fmt.Fprintf(os.Stderr, "filecoin: warning! unhandled response (%d) encoding: %s", resp.StatusCode, contentType)
			out, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "filecoin: response (%d) read error: %s\n", resp.StatusCode, err)
			}
			e.Message = fmt.Sprintf("unknown error encoding: %q - %q", contentType, out)
		}

		nresp.Error = e
		nresp.Output = nil

		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}
	return nresp, nil
}

func (r *Request) getURL() string {
	values := make(url.Values)
	for _, arg := range r.Args {
		values.Add("arg", arg)
	}
	for k, v := range r.Opts {
		values.Add(k, v)
	}

	return strings.TrimRight(fmt.Sprintf("%s/%s?%s", r.ApiBase, r.Command, values.Encode()), "?")
}
