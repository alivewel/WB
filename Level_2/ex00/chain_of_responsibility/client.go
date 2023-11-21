package main

type Client struct {
	handler Handler
}

func (c *Client) MakeRequest(request int) {
	c.handler.HandleRequest(request)
}
