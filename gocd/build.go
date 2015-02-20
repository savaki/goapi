package main

import "github.com/codegangsta/cli"

var build = cli.Command{
	Name:  "build",
	Usage: "commands for build status",
	Subcommands: []cli.Command{
		{
			Name:   "status",
			Usage:  "current status of all builds",
			Flags:  flags,
			Action: buildStatus,
		},
	},
}

func buildStatus(c *cli.Context) {
	client := newClient(c)
	projects, err := client.BuildStatus()
	assert(err)
	print(projects)
}
