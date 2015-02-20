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

func (filter ProjectFilter) matchWithin(project Project) bool {
	if filter.Within == 0 {
		return true
	}

	buildTime, err := project.LastBuildTime()
	if err != nil {
		return false
	}

	return buildTime.Add(filter.Within).Before(time.Now()) == false
}

func (filter ProjectFilter) matchActivity(project Project) bool {
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

func (filter ProjectFilter) matchStatus(project Project) bool {
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

func (filter ProjectFilter) matchName(project Project) bool {
	if filter.Name == nil {
		return true
	}

	return filter.Name.MatchString(project.Name)
}

func (filter ProjectFilter) match(project Project) bool {
	if !filter.matchWithin(project) {
		return false
	}

	if !filter.matchStatus(project) {
		return false
	}

	if !filter.matchName(project) {
		return false
	}

	if !filter.matchActivity(project) {
		return false
	}

	return true
}

func (filter ProjectFilter) Filter(projects []Project) []Project {
	results := []Project{}

	if projects != nil {
		for _, project := range projects {
			if filter.matchWithin(project) {
				projects = append(projects, project)
			}
		}
	}

	return results
}
