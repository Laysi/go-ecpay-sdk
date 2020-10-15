package ecpay

import "context"

func WithReturnURL(url string) optionFunc {
	return func(c *Client) {
		c.returnURL = url
	}
}

func WithPeriodReturnURL(url string) optionFunc {
	return func(c *Client) {
		c.periodReturnURL = &url
	}
}

func WithClientBackURL(url string) optionFunc {
	return func(c *Client) {
		c.clientBackURL = &url
	}
}
func WithPaymentInfoURL(url string) optionFunc {
	return func(c *Client) {
		c.paymentInfoURL = &url
	}
}
func WithClientRedirectURL(url string) optionFunc {
	return func(c *Client) {
		c.clientRedirectURL = &url
	}
}
func WithOrderResultURL(url string) optionFunc {
	return func(c *Client) {
		c.orderResultURL = &url
	}
}

func WithPlatformID(id string) optionFunc {
	return func(c *Client) {
		c.platformID = &id
	}
}

func WithCtxFunc(f func(c context.Context) context.Context) optionFunc {
	return func(c *Client) {
		c.ctxFunc = f
	}
}
