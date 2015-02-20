package main

import "github.com/codegangsta/cli"

var pipeline = cli.Command{
	Name:  "pipeline",
	Usage: "commands for pipelines",
	Subcommands: []cli.Command{
		{
			Name:   "schedule",
			Usage:  "schedule a pipeline with specific materials",
			Flags:  flags,
			Action: pipelineSchedule,
		},
		{
			Name:   "release",
			Usage:  "release a pipeline lock",
			Flags:  flags,
			Action: pipelineRelease,
		},
		{
			Name:   "pause",
			Usage:  "pause a pipeline",
			Flags:  flags,
			Action: pipelinePause,
		},
		{
			Name:   "unpause",
			Usage:  "unpause a pipeline",
			Flags:  flags,
			Action: pipelineUnpause,
		},
		{
			Name:   "status",
			Usage:  "check the status of a pipeline",
			Flags:  flags,
			Action: pipelineStatus,
		},
	},
}

func pipelineSchedule(c *cli.Context) {
	client := newClient(c)
	print(client)
}

func pipelineRelease(c *cli.Context) {
	client := newClient(c)
	print(client)
}

func pipelinePause(c *cli.Context) {
	client := newClient(c)
	print(client)
}

func pipelineUnpause(c *cli.Context) {
	client := newClient(c)
	print(client)
}

func pipelineStatus(c *cli.Context) {
	client := newClient(c)
	print(client)
}

func pipelineHistory(c *cli.Context) {
	client := newClient(c)
	print(client)
}
