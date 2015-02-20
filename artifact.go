package goapi

import "fmt"

func (c *Client) ArtifactList(b BuildIdentifier) ([]Artifact, error) {
	path := c.rawPathTo("/go/files/%s/%d/%s/%d/%s.json", b.PipelineName, b.PipelineCounter, b.StageName, b.StageCounter, b.JobName)

	artifacts := []Artifact{}
	err := c.api.Get(defaultContext(), path, nil, &artifacts)
	return artifacts, err
}

func (c *Client) ArtifactDownloadFile(b BuildIdentifier, pathname string) (Artifacts, error) {
	artifacts, err := c.ArtifactList(b)
	if err != nil {
		return nil, err
	}

	return artifacts, nil
}

func (c *Client) ArtifactDownloadZip(b BuildIdentifier, pathname string) (Artifacts, error) {
	path := c.rawPathTo("/go/files/%s/%d/%s/%d/%s/%s.zip", b.PipelineName, b.PipelineCounter, b.StageName, b.StageCounter, b.JobName, pathname)
	fmt.Println(path)

	data := []byte{}
	err := c.api.Get(defaultContext(), path, nil, &data)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))

	return nil, nil
}
