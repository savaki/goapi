package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var artifact = cli.Command{
	Name:  "artifact",
	Usage: "commands for artifacts",
	Subcommands: []cli.Command{
		{
			Name:   "list",
			Usage:  "list artifacts for a specific build",
			Flags:  append(flags, flagBuildIdentifier...),
			Action: artifactList,
		},
		{
			Name:   "download",
			Usage:  "download artifacts",
			Flags:  append(append(flags, flagBuildIdentifier...), flagDownload, flagPath),
			Action: artifactDownload,
		},
	},
}

func artifactList(c *cli.Context) {
	client := newClient(c)
	b := buildIdentifier(c)

	artifacts, err := client.ArtifactList(b)
	assert(err)
	print(artifacts)
}

func artifactDownload(c *cli.Context) {
	client := newClient(c)
	b := buildIdentifier(c)
	pathname := c.String("path")
	fmt.Println(b)

	switch c.String("download") {
	case "file":
		artifacts, err := client.ArtifactDownloadFile(b, pathname)
		assert(err)
		print(artifacts)
	default:
		artifacts, err := client.ArtifactDownloadZip(b, pathname)
		assert(err)
		print(artifacts)
	}
}
