package goapi

func (c *Client) AgentList() ([]Agent, error) {
	v := []Agent{}
	err := c.api.Get(defaultContext(), c.pathTo("/agents"), nil, &v)
	return v, err
}

func (c *Client) AgentEnable(uuid string) error {
	path := c.pathTo("/agents/%s/enable", uuid)
	return c.api.Post(defaultContext(), path, nil, nil)
}

func (c *Client) AgentDisable(uuid string) error {
	path := c.pathTo("/agents/%s/disable", uuid)
	return c.api.Post(defaultContext(), path, nil, nil)
}

func (c *Client) AgentDelete(uuid string) error {
	path := c.pathTo("/agents/%s/delete", uuid)
	return c.api.Post(defaultContext(), path, nil, nil)
}
