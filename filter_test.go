package goapi

import (
	"regexp"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestProjectFilter(t *testing.T) {
	var filter ProjectFilter
	var project Project
	var match bool

	Convey("Given a Project and a ProjectFilter", t, func() {
		project = Project{
			Name:                "First-Pipeline :: defaultStage",
			Activity:            "Sleeping",
			LastBuildStatus:     "Success",
			LastBuildLabel:      "4",
			LastBuildTimeString: "2015-02-20T10:13:09",
		}
		filter = ProjectFilter{}

		// #MatchStatus --------------------------------------------------------

		Convey("When I call #MatchStatus with no #Status defined", func() {
			match = filter.MatchStatus(project)

			Convey("Then I expect a match", func() {
				So(match, ShouldBeTrue)
			})
		})

		Convey("When I call #MatchStatus with a matching #Status defined", func() {
			filter.Status = []string{project.LastBuildStatus}
			match = filter.MatchStatus(project)

			Convey("Then I expect a match", func() {
				So(match, ShouldBeTrue)
			})
		})

		Convey("When I call #MatchStatus with a NON-matching #Status defined", func() {
			filter.Status = []string{project.LastBuildStatus + "blah"}
			match = filter.MatchStatus(project)

			Convey("Then I expect NO match", func() {
				So(match, ShouldBeFalse)
			})
		})

		// #MatchActivity ------------------------------------------------------

		Convey("When I call #MatchActivity with no #Activity defined", func() {
			match = filter.MatchActivity(project)

			Convey("Then I expect a match", func() {
				So(match, ShouldBeTrue)
			})
		})

		Convey("When I call #MatchActivity with a matching #Activity defined", func() {
			filter.Activity = []string{project.Activity}
			match = filter.MatchActivity(project)

			Convey("Then I expect a match", func() {
				So(match, ShouldBeTrue)
			})
		})

		Convey("When I call #MatchActivity with a NON-matching #Activity defined", func() {
			filter.Activity = []string{project.Activity + "blah"}
			match = filter.MatchActivity(project)

			Convey("Then I expect NO match", func() {
				So(match, ShouldBeFalse)
			})
		})

		// #MatchWithin --------------------------------------------------------

		Convey("When I call #MatchWithin with no #Within defined", func() {
			match = filter.MatchWithin(project)

			Convey("Then I expect a match", func() {
				So(match, ShouldBeTrue)
			})
		})

		Convey("When I call #MatchWithin with a matching #Within defined", func() {
			filter.Within = time.Hour * 24 * 365
			match = filter.MatchWithin(project)

			Convey("Then I expect a match", func() {
				So(match, ShouldBeTrue)
			})
		})

		Convey("When I call #MatchWithin with a NON-matching #Within defined", func() {
			filter.Within = time.Hour
			match = filter.MatchWithin(project)

			Convey("Then I expect NO match", func() {
				So(match, ShouldBeFalse)
			})
		})

		// #MatchName --------------------------------------------------------

		Convey("When I call #MatchName with no #Name defined", func() {
			match = filter.MatchName(project)

			Convey("Then I expect a match", func() {
				So(match, ShouldBeTrue)
			})
		})

		Convey("When I call #MatchName with a matching #Name defined", func() {
			filter.Name = regexp.MustCompile(`\S`)
			match = filter.MatchName(project)

			Convey("Then I expect a match", func() {
				So(match, ShouldBeTrue)
			})
		})

		Convey("When I call #MatchName with a NON-matching #Name defined", func() {
			filter.Name = regexp.MustCompile(`this-wont-match`)
			match = filter.MatchName(project)

			Convey("Then I expect NO match", func() {
				So(match, ShouldBeFalse)
			})
		})

		// #Match ------------------------------------------------------------

		Convey("When I call #Match with the default filter", func() {
			match = filter.Match(project)

			Convey("Then I expect a match", func() {
				So(match, ShouldBeTrue)
			})
		})

		Convey("When I call #Match with a NON-matching name", func() {
			filter.Name = regexp.MustCompile(`this-wont-match`)
			match = filter.Match(project)

			Convey("Then I expect NO match", func() {
				So(match, ShouldBeFalse)
			})
		})

		Convey("When I call #Match with a NON-matching activity", func() {
			filter.Activity = []string{"this-wont-match"}
			match = filter.Match(project)

			Convey("Then I expect NO match", func() {
				So(match, ShouldBeFalse)
			})
		})

		Convey("When I call #Match with a NON-matching status", func() {
			filter.Status = []string{"this-wont-match"}
			match = filter.Match(project)

			Convey("Then I expect NO match", func() {
				So(match, ShouldBeFalse)
			})
		})

		Convey("When I call #Match with a NON-matching within", func() {
			filter.Within = time.Second
			match = filter.Match(project)

			Convey("Then I expect NO match", func() {
				So(match, ShouldBeFalse)
			})
		})
	})
}
