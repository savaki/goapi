package goapi

func (c *Client) StageCancel(pipelineName, stageName string) error {
	path := c.pathTo("/stages/%s/%s/cancel", pipelineName, stageName)
	return c.api.Post(defaultContext(), path, nil, nil)
}
