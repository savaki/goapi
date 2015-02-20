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
	flagPipelineName    = cli.StringFlag{"pipeline", "", "the name of the pipeline", "GOCD_PIPELINE"}
	flagPipelineCounter = cli.StringFlag{"pipeline-counter", "", "the counter for the pipeline", "GOCD_PIPELINE_COUNTER"}
	flagStageName       = cli.StringFlag{"stage", "", "the name of the stage", "GOCD_STAGE"}
	flagStageCounter    = cli.StringFlag{"stage-counter", "", "the counter for the stage", "GOCD_STAGE_COUNTER"}
	flagJobName         = cli.StringFlag{"job", "", "the name of the job", "GOCD_JOB"}
	flagAgentUuid       = cli.StringFlag{"agent-uuid", "", "the uuid of the agent", "GOCD_AGENT_UUID"}
	flagOffset          = cli.IntFlag{"offset", 0, "the offset for pagination", "GOCD_OFFSET"}
	flagDownload        = cli.StringFlag{"download", "file", "what format to download => zip or file", "GOCD_DOWNLOAD"}
	flagPath            = cli.StringFlag{"path", "", "the path of the artifact to retrieve", "GOCD_PATH"}
)

var (
	flagBuildIdentifier = []cli.Flag{
		flagPipelineName,
		flagPipelineCounter,
		flagStageName,
		flagStageCounter,
		flagJobName,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "gocd"
	app.Usage = "command line interface to Go CD"
	app.Author = "Matt"
	app.Commands = []cli.Command{
		agent,
		artifact,
		build,
		job,
		pipeline,
		pipelineGroups,
		stage,
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

func buildIdentifier(c *cli.Context) goapi.BuildIdentifier {
	return goapi.BuildIdentifier{
		PipelineName:    c.String(flagPipelineName.Name),
		PipelineCounter: c.Int(flagPipelineCounter.Name),
		StageName:       c.String(flagStageName.Name),
		StageCounter:    c.Int(flagStageCounter.Name),
		JobName:         c.String(flagJobName.Name),
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
	if v == nil {
		fmt.Println("{}")
	} else {
		text := marshalIndent(v)
		if text == "null" {
			text = "[]"
		}
		fmt.Println(text)
	}
}
