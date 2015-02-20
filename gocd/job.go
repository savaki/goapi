package main

import "github.com/codegangsta/cli"

var job = cli.Command{
	Name:  "job",
	Usage: "commands for jobs",
	Subcommands: []cli.Command{
		{
			Name:   "history",
			Usage:  "show job history",
			Flags:  append(flags, flagPipelineName, flagStageName, flagJobName, flagOffset),
			Action: jobHistory,
		},
		{
			Name:   "scheduled",
			Usage:  "show scheduled jobs",
			Flags:  flags,
			Action: jobScheduled,
		},
	},
}

func jobHistory(c *cli.Context) {
	client := newClient(c)
	pipeline := c.String("pipeline")
	stage := c.String("stage")
	job := c.String("job")
	offset := c.Int("offset")

	history, err := client.JobHistory(pipeline, stage, job, offset)
	assert(err)
	print(history)
}

func jobScheduled(c *cli.Context) {
	client := newClient(c)

	jobs, err := client.JobScheduled()
	assert(err)
	print(jobs)
}
