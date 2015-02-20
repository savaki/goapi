package goapi

import (
	"bytes"
	"encoding/xml"
)

func (c *Client) BuildStatus() ([]Project, error) {
	path := c.rawPathTo("/go/cctray.xml")
	data := []byte{}
	err := c.api.Get(defaultContext(), path, nil, &data)
	if err != nil {
		return nil, err
	}

	// decode the content
	ccTray := CCTray{}
	err = xml.NewDecoder(bytes.NewReader(data)).Decode(&ccTray)
	return ccTray.Projects, err
}
