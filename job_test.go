package goapi

import (
	"testing"

	"github.com/savaki/httpctx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestJob(t *testing.T) {
	var client *Client

	Convey("Given a Client", t, func() {
		client = &Client{}

		Convey("When I call #JobScheduled", func() {
			content := `<scheduledJobs>
          <job name="fresh.install.go" id="186225">
            <link rel="self" href="http://go-server:8153/go/tab/build/detail/auto-deploy-testing-open-solaris/11/fresh-install/1/fresh.install.go"/>
            <buildLocator>
              auto-deploy-testing-open-solaris/11/fresh-install/1/fresh.install.go
            </buildLocator>
            <environment>AutoDeploy-OpenSolaris</environment>
            <resources>
              <resource>autodeploy</resource>
            </resources>
            <environmentVariables>
              <variable name="TWIST_SERVER_PATH">/etc/go</variable>
              <variable name="TWIST_SERVER_CONFIG_PATH">/etc/go</variable>
              <variable name="TWIST_AGENT_PATH">/var/lib/go-agent</variable>
            </environmentVariables>
          </job>
        </scheduledJobs>`
			client.api = httpctx.Mock{Body: content}
			jobs, err := client.JobScheduled()

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect the jobs back", func() {
				So(jobs, ShouldNotBeNil)
			})
		})

		Convey("When I call #JobHistory", func() {
			content := `{
  "jobs": [
    {
      "agent_uuid": "0794793b-5c1a-443f-b860-df480986586b",
      "name": "UnitTest",
      "job_state_transitions": [],
      "scheduled_date": 1411456876262,
      "pipeline_counter": 2,
      "rerun": true,
      "pipeline_name": "foo",
      "result": "Failed",
      "state": "Completed",
      "id": 3,
      "stage_counter": "1",
      "stage_name": "DEV"
    }
  ],
  "pagination": {
    "offset": 1,
    "total": 2,
    "page_size": 10
  }
}`
			client.api = httpctx.Mock{Body: content}
			history, _ := client.JobHistory("pipeline", "stage", "job", 0)

			Convey("Then I expect a valid JobHistory", func() {
				So(len(history.Jobs), ShouldEqual, 1)

				job := history.Jobs[0]
				So(job.AgentUuid, ShouldEqual, "0794793b-5c1a-443f-b860-df480986586b")
				So(job.Name, ShouldEqual, "UnitTest")
				So(job.ScheduledDate, ShouldEqual, 1411456876262)
				So(job.PipelineCounter, ShouldEqual, 2)
				So(job.Rerun, ShouldBeTrue)
				So(job.PipelineName, ShouldEqual, "foo")
				So(job.Result, ShouldEqual, "Failed")
				So(job.State, ShouldEqual, "Completed")
				So(job.Id, ShouldEqual, 3)
				So(job.StageCounter, ShouldEqual, "1")
				So(job.StageName, ShouldEqual, "DEV")

				pagination := history.Pagination
				So(pagination.Offset, ShouldEqual, 1)
				So(pagination.Total, ShouldEqual, 2)
				So(pagination.PageSize, ShouldEqual, 10)
			})
		})
	})
}
