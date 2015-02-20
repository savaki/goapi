package goapi

func (c *Client) PipelineGroups() ([]PipelineGroup, error) {
	v := []PipelineGroup{}
	err := c.api.Get(defaultContext(), c.pathTo("/config/pipeline_groups"), nil, &v)
	return v, err
}
