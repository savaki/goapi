package goapi

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/savaki/httpctx"
	"golang.org/x/net/context"
)

func defaultContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

type Client struct {
	codebase string
	api      httpctx.HttpClient
}

func New(codebase string) *Client {
	if strings.HasSuffix(codebase, "/") {
		codebase = codebase[0 : len(codebase)-1]
	}
	return &Client{
		codebase: codebase,
		api:      httpctx.NewClient(),
	}
}

func WithAuth(c *Client, username, password string) *Client {
	if username != "" && password != "" {
		authFunc := func(req *http.Request) *http.Request {
			req.SetBasicAuth(username, password)
			return req
		}
		c.api = httpctx.WithAuthFunc(authFunc)
	}
	return c
}

func (c *Client) pathTo(format string, args ...interface{}) string {
	path := fmt.Sprintf(format, args...)
	return c.codebase + "/go/api" + path
}
