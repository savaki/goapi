package goapi

import "encoding/xml"

func (c *Client) JobScheduled() ([]Job, error) {
	data := []byte{}
	path := c.pathTo("/jobs/scheduled.xml")
	err := c.api.Get(defaultContext(), path, nil, &data)
	if err != nil {
		return nil, err
	}

	// convert data to xml
	scheduled := struct {
		Jobs []Job `xml:"job"`
	}{}
	err = xml.Unmarshal(data, &scheduled)
	if err != nil {
		return nil, err
	}

	return scheduled.Jobs, nil
}

func (c *Client) JobHistory(pipelineName, stageName, jobName string, offset int) (JobHistory, error) {
	path := c.pathTo("/jobs/%s/%s/%s/history/%d", pipelineName, stageName, jobName, offset)
	history := JobHistory{}
	err := c.api.Get(defaultContext(), path, nil, &history)
	return history, err
}
