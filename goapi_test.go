package goapi

import (
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWalk(t *testing.T) {
	Convey("Given a set of artifacts", t, func() {
		artifacts := []Artifact{
			{
				Name: "file1.txt",
				Url:  "http://blah",
				Type: "file",
			},
			{
				Name: "file2.txt",
				Url:  "http://blah",
				Type: "file",
			},
			{
				Name: "dir",
				Type: "folder",
				Files: []Artifact{
					{
						Name: "file3.txt",
						Url:  "http://blah",
						Type: "file",
					},
					{
						Name: "file4.txt",
						Url:  "http://blah",
						Type: "file",
					},
				},
			},
		}

		Convey("When I #Walk the artifacts", func() {
			visitor := &MockVisitor{}
			client := &Client{
				log:      noOpLog,
				Download: MockDownload{}.HandlerFunc,
			}

			err := client.Walk(artifacts, visitor.HandlerFunc)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I expect the files to have been walked", func() {
				sort.Strings(visitor.Paths)
				So(visitor.Paths, ShouldResemble, []string{
					"dir/file3.txt",
					"dir/file4.txt",
					"file1.txt",
					"file2.txt",
				})
			})
		})
	})
}

type MockVisitor struct {
	Err   error
	Paths []string
}

func (m *MockVisitor) HandlerFunc(path string, r io.Reader) error {
	if m.Paths == nil {
		m.Paths = []string{}
	}
	m.Paths = append(m.Paths, path)
	return m.Err
}

type MockDownload struct {
	Err     error
	Content string
}

func (m MockDownload) HandlerFunc(url string) (io.ReadCloser, error) {
	content := m.Content
	if content == "" {
		content = "hello world"
	}
	r := strings.NewReader(m.Content)
	return ioutil.NopCloser(r), m.Err
}
