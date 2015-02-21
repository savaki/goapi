package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/savaki/goapi"
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
			Flags:  append(append(flags, flagBuildIdentifier...), flagDownload, flagPath, flagOutput),
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
	path := c.String(flagPath.Name)
	dir := c.String(flagOutput.Name)

	// retrieve all the artifacts
	artifacts, err := client.ArtifactList(b)
	assert(err)

	// find artifacts that match the provided path
	debug("searching for artifact, %s\n", path)
	artifact, err := artifacts.Find(path)
	assert(err)

	if artifact.Type == "file" {
		err = saveArtifact(client, artifact)
		assert(err)

	} else {
		basedir := fmt.Sprintf("%s/%s/", dir, filepath.Base(path))
		basedir = filepath.Clean(basedir)
		err = client.Walk(artifact.Files, newVisitor(basedir))
		assert(err)
	}
}

func newVisitor(basedir string) goapi.Visitor {
	return func(path string, r io.Reader) error {
		filename := filepath.Clean(basedir + "/" + path)
		dirname := filepath.Dir(filename)
		if _, err := os.Stat(dirname); err != nil {
			if os.IsNotExist(err) {
				debug("creating directory, %s\n", dirname)
				os.MkdirAll(dirname, 0755)
			} else {
				return err
			}
		}

		return writeFile(filename, r)
	}
}

func saveArtifact(client *goapi.Client, artifact *goapi.Artifact) error {
	r, err := client.Download(artifact.Url)
	if err != nil {
		return err
	}
	defer r.Close()

	return writeFile(artifact.Name, r)
}

func writeFile(filename string, r io.Reader) error {
	debug("writing file, %s\n", filename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}
