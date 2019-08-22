package api

/*简单示例*/
import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/filecoin-project/go-filecoin/commands"
)

type API struct {
	host string
}

func New(url string) *API {
	if url == "" {
		url = "127.0.0.1:3453"
	}
	return &API{
		host: url,
	}
}

func (a *API) Request(args ...string) ([]byte, error) {
	url := fmt.Sprintf("http://%s/api", a.host)
	for _, arg := range args {
		url = url + "/" + arg
	}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, nil
	}
	defer res.Body.Close()
	s := commands.IDDetails{}
	fmt.Println(s)
	body, err := ioutil.ReadAll(res.Body)
	return body, nil
}
