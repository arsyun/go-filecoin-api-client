package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type RequestBuilder interface {
	Arguments(args ...string) RequestBuilder
	BodyString(body string) RequestBuilder
	BodyBytes(body []byte) RequestBuilder
	Body(body io.Reader) RequestBuilder
	Option(key string, value interface{}) RequestBuilder
	Header(name, value string) RequestBuilder
	Send(ctx context.Context) (*Response, error)
	Exec(ctx context.Context, res interface{}) error
}

type requestBuilder struct {
	command string
	args    []string
	opts    map[string]string
	headers map[string]string
	body    io.Reader

	shell *HttpApi
}

func (r *requestBuilder) Arguments(args ...string) RequestBuilder {
	r.args = append(r.args, args...)
	return r
}

func (r *requestBuilder) BodyString(body string) RequestBuilder {
	return r.Body(strings.NewReader(body))
}

func (r *requestBuilder) BodyBytes(body []byte) RequestBuilder {
	return r.Body(bytes.NewReader(body))
}

func (r *requestBuilder) Body(body io.Reader) RequestBuilder {
	r.body = body
	return r
}

// Option sets the given option.
func (r *requestBuilder) Option(key string, value interface{}) RequestBuilder {
	var s string
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprint(value)
	}
	if r.opts == nil {
		r.opts = make(map[string]string, 1)
	}
	r.opts[key] = s
	return r
}

func (r *requestBuilder) Header(name, value string) RequestBuilder {
	if r.headers == nil {
		r.headers = make(map[string]string, 1)
	}
	r.headers[name] = value
	return r
}

func (r *requestBuilder) Send(ctx context.Context) (*Response, error) {
	req := NewRequest(ctx, r.shell.url, r.command, r.args...)
	req.Opts = r.opts
	req.Headers = r.headers
	req.Body = r.body

	return req.Send(&r.shell.httpcli)
}

func (r *requestBuilder) Exec(ctx context.Context, res interface{}) error {
	httpRes, err := r.Send(ctx)
	if err != nil {
		return err
	}

	if res == nil {
		lateErr := httpRes.Close()
		if httpRes.Error != nil {
			return httpRes.Error
		}
		return lateErr
	}

	return httpRes.decode(res)
}

var _ RequestBuilder = &requestBuilder{}
