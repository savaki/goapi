package goapi

import (
	"regexp"
	"time"
)

type ProjectFilter struct {
	Within   time.Duration
	Activity []string
	Status   []string
	Name     *regexp.Regexp
}

func (filter ProjectFilter) MatchWithin(project Project) bool {
	if filter.Within == 0 {
		return true
	}

	buildTime, err := project.LastBuildTime()
	if err != nil {
		return false
	}

	return buildTime.Add(filter.Within).Before(time.Now()) == false
}

func (filter ProjectFilter) MatchActivity(project Project) bool {
	if filter.Activity == nil {
		return true
	}

	for _, activity := range filter.Activity {
		if activity == project.Activity {
			return true
		}
	}

	return false
}

func (filter ProjectFilter) MatchStatus(project Project) bool {
	if filter.Status == nil {
		return true
	}

	for _, status := range filter.Status {
		if status == project.LastBuildStatus {
			return true
		}
	}

	return false
}

func (filter ProjectFilter) MatchName(project Project) bool {
	if filter.Name == nil {
		return true
	}

	return filter.Name.MatchString(project.Name)
}

func (filter ProjectFilter) Match(project Project) bool {
	if !filter.MatchWithin(project) {
		return false
	}

	if !filter.MatchStatus(project) {
		return false
	}

	if !filter.MatchName(project) {
		return false
	}

	if !filter.MatchActivity(project) {
		return false
	}

	return true
}

func (filter ProjectFilter) Filter(projects []Project) []Project {
	results := []Project{}

	if projects != nil {
		for _, project := range projects {
			if filter.MatchWithin(project) {
				projects = append(projects, project)
			}
		}
	}

	return results
}
