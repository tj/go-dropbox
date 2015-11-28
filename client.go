package dropbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Client implements a Dropbox client. You may use the Files and Users
// clients directly if preferred, however Client exposes them both.
type Client struct {
	*Config
	Users *Users
	Files *Files
}

// New client.
func New(config *Config) *Client {
	c := &Client{Config: config}
	c.Users = &Users{c}
	c.Files = &Files{c}
	return c
}

// call rpc style endpoint.
func (c *Client) call(path string, in interface{}) (io.ReadCloser, error) {
	url := "https://api.dropboxapi.com/2" + path

	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	return c.do(req)
}

// download style endpoint.
func (c *Client) download(path string, in interface{}, r io.Reader) (io.ReadCloser, error) {
	url := "https://content.dropboxapi.com/2" + path

	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Dropbox-API-Arg", string(body))

	if r != nil {
		req.Header.Set("Content-Type", "application/octet-stream")
	}

	return c.do(req)
}

// perform the request.
func (c *Client) do(req *http.Request) (io.ReadCloser, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		defer res.Body.Close()
		if b, err := ioutil.ReadAll(res.Body); err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf("response %s: %s", res.Status, b)
		}
	}

	return res.Body, err
}
