package goapi

func (c *Client) PipelineGroups() ([]PipelineGroup, error) {
	v := []PipelineGroup{}
	err := c.api.Get(defaultContext(), c.pathTo("/config/pipeline_groups"), nil, &v)
	return v, err
}

func (c *Client) PipelineSchedule(pipelineName string) error {
	path := c.pathTo("/pipelines/%s/schedule", pipelineName)
	return c.api.Post(defaultContext(), path, nil, nil)
}

func (c *Client) PipelineReleaseLock(pipelineName string) error {
	path := c.pathTo("/pipelines/%s/releaseLock", pipelineName)
	return c.api.Post(defaultContext(), path, nil, nil)
}

func (c *Client) PipelinePause(pipelineName string) error {
	path := c.pathTo("/pipelines/%s/pause", pipelineName)
	return c.api.Post(defaultContext(), path, nil, nil)
}

func (c *Client) PipelineUnpause(pipelineName string) error {
	path := c.pathTo("/pipelines/%s/unpause", pipelineName)
	return c.api.Post(defaultContext(), path, nil, nil)
}

func (c *Client) PipelineStatus(pipelineName string) (PipelineStatus, error) {
	path := c.pathTo("/pipelines/%s/status", pipelineName)
	status := PipelineStatus{}
	err := c.api.Post(defaultContext(), path, nil, &status)
	return status, err
}

func (c *Client) PipelineHistory(pipelineName string, offset int) (PipelineHistory, error) {
	path := c.pathTo("/pipelines/%s/history/%d", pipelineName, offset)
	history := PipelineHistory{}
	err := c.api.Get(defaultContext(), path, nil, &history)
	return history, err
}
