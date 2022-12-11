package http

import (
	"io/ioutil"
	"net/http"
)

type client struct {
	hc *http.Client
}

type Client interface {
	Get(url string) ([]byte, error)
}

func NewClient(hc *http.Client) Client {
	return &client{
		hc: hc,
	}
}

func (c client) Get(url string) ([]byte, error) {
	r, err := c.hc.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
