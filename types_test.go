package goapi

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCCTray(t *testing.T) {
	Convey("Given a cctray xml content", t, func() {
		content := `<?xml version="1.0" encoding="utf-8"?>
<Projects>
  <Project name="First-Pipeline :: defaultStage" activity="Sleeping" lastBuildStatus="Success" lastBuildLabel="4" lastBuildTime="2015-02-20T10:13:09" webUrl="http://192.168.99.101:8153/go/pipelines/First-Pipeline/4/defaultStage/1" />
  <Project name="First-Pipeline :: defaultStage :: defaultJob" activity="Sleeping" lastBuildStatus="Success" lastBuildLabel="4" lastBuildTime="2015-02-20T10:13:09" webUrl="http://192.168.99.101:8153/go/tab/build/detail/First-Pipeline/4/defaultStage/1/defaultJob" />
</Projects>`

		Convey("When I decode the content to an instance of CCTray", func() {
			ccTray := CCTray{}
			err := xml.NewDecoder(strings.NewReader(content)).Decode(&ccTray)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect the ccTray to have the correct values assigned", func() {
				So(ccTray.Projects, ShouldNotBeNil)
				So(len(ccTray.Projects), ShouldEqual, 2)

				var project Project

				project = ccTray.Projects[0]
				So(project.Name, ShouldEqual, "First-Pipeline :: defaultStage")
				So(project.Activity, ShouldEqual, "Sleeping")
				So(project.LastBuildStatus, ShouldEqual, "Success")
				So(project.LastBuildLabel, ShouldEqual, "4")
				So(project.LastBuildTimeString, ShouldEqual, "2015-02-20T10:13:09")
				So(project.WebUrl, ShouldEqual, "http://192.168.99.101:8153/go/pipelines/First-Pipeline/4/defaultStage/1")

				buildTime, err := project.LastBuildTime()
				So(err, ShouldBeNil)
				So(buildTime.Year(), ShouldEqual, 2015)
				So(buildTime.Month(), ShouldEqual, time.February)
				So(buildTime.Day(), ShouldEqual, 20)
				So(buildTime.Hour(), ShouldEqual, 10)
				So(buildTime.Minute(), ShouldEqual, 13)
				So(buildTime.Second(), ShouldEqual, 9)
			})
		})
	})

	Convey("Given a scheduled.xml", t, func() {
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
          <job name="publish" id="285717">
            <link rel="self" href="http://go-server:8153/go/tab/build/detail/go-ec2-plugin/26/dist/1/publish"/>
            <buildLocator>go-ec2-plugin/26/dist/1/publish</buildLocator>
            <environment>performance-ec2</environment>
            <resources>
              <resource>deploy-agent</resource>
            </resources>
          </job>
          <job name="upgrade" id="297092">
            <link rel="self" href="http://go-server:8153/go/tab/build/detail/upgrade_qa_server/15/upgrade/1/upgrade"/>
            <buildLocator>upgrade_qa_server/15/upgrade/1/upgrade</buildLocator>
            <environment>UAT</environment>
            <resources>
              <resource>UAT-Server</resource>
            </resources>
          </job>
        </scheduledJobs>`

		jobs := struct {
			Jobs []ScheduledJob `xml:"job"`
		}{}
		err := xml.NewDecoder(strings.NewReader(content)).Decode(&jobs)

		Convey("Then I expect no errors", func() {
			So(err, ShouldBeNil)
		})

		Convey("And I expect the jobs to be parsed correctly", func() {
			fmt.Println(jobs)
			So(len(jobs.Jobs), ShouldBeGreaterThan, 0)

			var job ScheduledJob

			// job 1
			job = jobs.Jobs[0].Trim()
			So(job.Id, ShouldEqual, "186225")
			So(job.Name, ShouldEqual, "fresh.install.go")
			So(job.Link.Rel, ShouldEqual, "self")
			So(job.Link.Href, ShouldEqual, "http://go-server:8153/go/tab/build/detail/auto-deploy-testing-open-solaris/11/fresh-install/1/fresh.install.go")
			So(job.BuildLocator, ShouldEqual, "auto-deploy-testing-open-solaris/11/fresh-install/1/fresh.install.go")

			So(len(job.EnvironmentVariables), ShouldEqual, 3)
			So(job.EnvironmentVariables[0].Name, ShouldEqual, "TWIST_SERVER_PATH")
			So(job.EnvironmentVariables[0].Value, ShouldEqual, "/etc/go")
			So(job.EnvironmentVariables[1].Name, ShouldEqual, "TWIST_SERVER_CONFIG_PATH")
			So(job.EnvironmentVariables[1].Value, ShouldEqual, "/etc/go")
			So(job.EnvironmentVariables[2].Name, ShouldEqual, "TWIST_AGENT_PATH")
			So(job.EnvironmentVariables[2].Value, ShouldEqual, "/var/lib/go-agent")

			So(job.Resources, ShouldResemble, []string{"autodeploy"})
		})
	})
}
