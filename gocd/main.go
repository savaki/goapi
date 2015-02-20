package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/savaki/goapi"
)

var flags = []cli.Flag{
	cli.StringFlag{"username", "", "Go CD username", "GOCD_USERNAME"},
	cli.StringFlag{"password", "", "Go CD password", "GOCD_PASSWORD"},
	cli.StringFlag{"url", "http://localhost:8153", "Go CD server codebase", "GOCD_URL"},
}

var (
	flagPipelineName = cli.StringFlag{"pipeline", "", "the name of the pipeline", "GOCD_PIPELINE"}
	flagStageName    = cli.StringFlag{"stage", "", "the name of the stage", "GOCD_STAGE"}
	flagAgentUuid    = cli.StringFlag{"agent-uuid", "", "the uuid of the agent", "GOCD_AGENT_UUID"}
	flagOffset       = cli.IntFlag{"offset", 0, "the offset for pagination", "GOCD_OFFSET"}
)

func main() {
	app := cli.NewApp()
	app.Name = "gocd"
	app.Usage = "command line interface to Go CD"
	app.Author = "Matt"
	app.Commands = []cli.Command{
		agent,
		pipeline,
		pipelineGroups,
		stage,
		build,
	}
	app.Run(os.Args)
}

func newClient(c *cli.Context) *goapi.Client {
	username := c.String("username")
	password := c.String("password")
	url := c.String("url")

	return goapi.WithAuth(goapi.New(url), username, password)
}

func assert(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func marshalIndent(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func print(v interface{}) {
	fmt.Println(marshalIndent(v))
}
