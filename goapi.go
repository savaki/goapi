package goapi

import (
	"fmt"
	"io"
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
	download func(url string) (io.ReadCloser, error)
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
		c.download = downloadFunc(authFunc)
	}
	return c
}

func (c *Client) pathTo(format string, args ...interface{}) string {
	path := fmt.Sprintf(format, args...)
	return c.codebase + "/go/api" + path
}

func (c *Client) rawPathTo(format string, args ...interface{}) string {
	path := fmt.Sprintf(format, args...)
	return c.codebase + path
}

// ------------------------------------------------------------------

func (c *Client) walkFile(path string, artifact Artifact, visitor Visitor) error {
	r, err := c.download(artifact.Url)
	if err != nil {
		return err
	}
	defer r.Close()

	return visitor(path, r)
}

func (c *Client) walkFolders(path string, artifacts Artifacts, visitor Visitor) error {
	var err error

	for _, artifact := range artifacts {
		switch artifact.Type {
		case "file":
			err = c.walkFile(path+artifact.Name, artifact, visitor)
		default:
			err = c.walkFolders(path+artifact.Name+"/", artifact.Files, visitor)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Walk(artifacts Artifacts, visitor Visitor) error {
	return c.walkFolders("", artifacts, visitor)
}

// ------------------------------------------------------------------

func downloadFunc(authFunc httpctx.AuthFunc) func(url string) (io.ReadCloser, error) {
	return func(url string) (io.ReadCloser, error) {
		// construct the request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req = authFunc(req)

		// retrieve the content
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		return resp.Body, nil
	}
}
