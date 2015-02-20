package main

import "github.com/codegangsta/cli"

var pipelineGroups = cli.Command{
	Name:  "pipeline-groups",
	Usage: "commands for pipeline groups",
	Subcommands: []cli.Command{
		{
			Name:   "list",
			Usage:  "list pipeline groups",
			Flags:  flags,
			Action: pipelineGroupsList,
		},
	},
}

func pipelineGroupsList(c *cli.Context) {
	client := newClient(c)
	groups, err := client.PipelineGroups()
	assert(err)

	print(groups)
}
