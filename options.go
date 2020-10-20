package ecpay

import "context"

func WithReturnURL(url string) OptionFunc {
	return func(c *Client) {
		c.returnURL = url
	}
}

func WithPeriodReturnURL(url string) OptionFunc {
	return func(c *Client) {
		c.periodReturnURL = &url
	}
}

func WithClientBackURL(url string) OptionFunc {
	return func(c *Client) {
		c.clientBackURL = &url
	}
}
func WithPaymentInfoURL(url string) OptionFunc {
	return func(c *Client) {
		c.paymentInfoURL = &url
	}
}
func WithClientRedirectURL(url string) OptionFunc {
	return func(c *Client) {
		c.clientRedirectURL = &url
	}
}
func WithOrderResultURL(url string) OptionFunc {
	return func(c *Client) {
		c.orderResultURL = &url
	}
}

func WithPlatformID(id string) OptionFunc {
	return func(c *Client) {
		c.platformID = &id
	}
}

func WithCtxFunc(f func(c context.Context) context.Context) OptionFunc {
	return func(c *Client) {
		c.ctxFunc = f
	}
}

func WithDebug(c *Client) {
	c.apiClient.GetConfig().Debug = true
}
