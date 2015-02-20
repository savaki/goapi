package main

import "github.com/codegangsta/cli"

var pipeline = cli.Command{
	Name:  "pipeline",
	Usage: "commands for pipelines",
	Subcommands: []cli.Command{
		{
			Name:   "schedule",
			Usage:  "schedule a pipeline with specific materials",
			Flags:  append(flags, flagPipelineName),
			Action: pipelineSchedule,
		},
		{
			Name:   "release",
			Usage:  "release a pipeline lock",
			Flags:  append(flags, flagPipelineName),
			Action: pipelineRelease,
		},
		{
			Name:   "pause",
			Usage:  "pause a pipeline",
			Flags:  append(flags, flagPipelineName),
			Action: pipelinePause,
		},
		{
			Name:   "unpause",
			Usage:  "unpause a pipeline",
			Flags:  append(flags, flagPipelineName),
			Action: pipelineUnpause,
		},
		{
			Name:   "status",
			Usage:  "check the status of a pipeline",
			Flags:  append(flags, flagPipelineName),
			Action: pipelineStatus,
		},
		{
			Name:   "history",
			Usage:  "retrieve the history of a pipeline",
			Flags:  append(flags, flagPipelineName, flagOffset),
			Action: pipelineHistory,
		},
	},
}

func pipelineSchedule(c *cli.Context) {
	client := newClient(c)
	pipeline := c.String("pipeline")
	err := client.PipelineSchedule(pipeline)
	assert(err)
}

func pipelineRelease(c *cli.Context) {
	client := newClient(c)
	pipeline := c.String("pipeline")
	err := client.PipelineReleaseLock(pipeline)
	assert(err)
}

func pipelinePause(c *cli.Context) {
	client := newClient(c)
	pipeline := c.String("pipeline")
	err := client.PipelinePause(pipeline)
	assert(err)
}

func pipelineUnpause(c *cli.Context) {
	client := newClient(c)
	pipeline := c.String("pipeline")
	err := client.PipelineUnpause(pipeline)
	assert(err)
}

func pipelineStatus(c *cli.Context) {
	client := newClient(c)
	pipeline := c.String("pipeline")
	status, err := client.PipelineStatus(pipeline)
	assert(err)
	print(status)
}

func pipelineHistory(c *cli.Context) {
	client := newClient(c)
	pipeline := c.String("pipeline")
	offset := c.Int("offset")
	history, err := client.PipelineHistory(pipeline, offset)
	assert(err)
	print(history)
}
