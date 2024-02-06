package bts

import (
	"github.com/netomathias/bts-client/log"
	"net/http"
)

const productionURL = "https://in.logs.betterstack.com"

type Client struct {
	Url        string
	HttpClient *http.Client

	Log log.Service
}

func NewClient(sourceToken string, options ...Option) (*Client, error) {
	c := Client{
		Url:        productionURL,
		HttpClient: &http.Client{},
	}

	for _, option := range options {
		if err := option(&c); err != nil {
			return nil, err
		}
	}

	c.Log = log.NewService(c.HttpClient, c.Url, sourceToken)

	return &c, nil
}
