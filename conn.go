package goinsta

import (
	"encoding/json"
	"github.com/ahmdrz/goinsta/endpoint"
	"github.com/valyala/fasthttp"
)

type Conn struct {
	response *fasthttp.Response
	request  *fasthttp.Request
	client   *fasthttp.Client
}

func NewConn() *Conn {
	conn := new(Conn)
	conn.Client = &fasthttp.Client{
		Name: APIUserAgent,
	}
	conn.requests = fasthttp.AcquireRequest()
	conn.response = fasthttp.AcquireResponse()
}

func (c *Conn) Do(ep int) ([]byte, error) {
	e := endpoint.endpoint[ep]
	if e == "" {
		return nil, ErrNoEndPoint
	}

	c.request.SetRequestURI(
		fmt.Sprintf("%s/%s", APIUrl, e),
	)
	err := c.client.Do(c.request, c.response)
	if err != nil {
		return nil, err
	}
	return c.response.Body(), err
}

func (c *Conn) Close() {
	c.response = nil
	c.request = nil
	c.client = nil
}
