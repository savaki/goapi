package goapi

import (
	"encoding/xml"
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
}
