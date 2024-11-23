package gopen

import (
	"regexp"
)

type PsmBuild struct {
	BuildMap map[string][]*regexp.Regexp
}

func NewPsmBuild() (*PsmBuild, error) {
	psmBuildMap := map[string][]*regexp.Regexp{
		"dss": make([]*regexp.Regexp, 0),
		"ent": make([]*regexp.Regexp, 0),
		"sdn": make([]*regexp.Regexp, 0),
	}
	r, err := regexp.Compile("grafana-ui")
	if err != nil {
		return &PsmBuild{}, err
	}
	psmBuildMap["dss"] = append(psmBuildMap["dss"], r)
	r, err = regexp.Compile("^1\\.29.*")
	if err != nil {
		return &PsmBuild{}, err
	}
	psmBuildMap["dss"] = append(psmBuildMap["dss"], r)
	r, err = regexp.Compile("^1\\.49.*")
	if err != nil {
		return &PsmBuild{}, err
	}
	psmBuildMap["dss"] = append(psmBuildMap["dss"], r)
	r, err = regexp.Compile("^1\\.51.*")
	if err != nil {
		return &PsmBuild{}, err
	}
	psmBuildMap["dss"] = append(psmBuildMap["dss"], r)
	r, err = regexp.Compile("^1\\.48.*")
	if err != nil {
		return &PsmBuild{}, err
	}
	psmBuildMap["dss"] = append(psmBuildMap["dss"], r)
	r, err = regexp.Compile("^1\\.28.*")
	if err != nil {
		return &PsmBuild{}, err
	}
	psmBuildMap["dss"] = append(psmBuildMap["dss"], r)
	r, err = regexp.Compile("^1\\.50.*")
	if err != nil {
		return &PsmBuild{}, err
	}
	psmBuildMap["ent"] = append(psmBuildMap["ent"], r)
	r, err = regexp.Compile("^1\\.16.*")
	if err != nil {
		return &PsmBuild{}, err
	}
	psmBuildMap["sdn"] = append(psmBuildMap["sdn"], r)
	return &PsmBuild{psmBuildMap}, nil
}

func (psmBuild *PsmBuild) Kind(version string) string {
	for kind, regexps := range psmBuild.BuildMap {
		for _, regexp := range regexps {
			if regexp.MatchString(version) {
				return kind
			}
		}
	}
	return ""
}
