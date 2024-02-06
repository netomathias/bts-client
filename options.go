package bts

import (
	"net/url"
)

type Option func(*Client) error

func WithURL(u string) Option {
	return func(c *Client) error {
		if _, err := url.ParseRequestURI(u); err != nil {
			return err
		}
		c.Url = u

		return nil
	}
}
