package convert

import (
	"github.com/pensando/go-psm/psm_dss"
	"regexp"
)

type ServiceGroup struct {
}

type ServiceGroupCSV struct {
}

func IsEntryServiceDisabled(service string) bool {
	match, err := regexp.MatchString("[Dd]isabled", service)
	if err != nil {
		return false
	}
	return match
}

func EntryServiceGroup(service string) string {
	if IsEntryServiceDisabled(service) {
		return ""
	}
	_, dssProtoPort := EntryServiceToDssProtoPort(service)
	if dssProtoPort {
		return ""
	}
	re, err := regexp.Compile(`[0-9a-zA-Z\-]+`)
	if err != nil {
	} else {
		matches := re.FindStringSubmatch(service)
		if len(matches) > 0 {
			return matches[0]
		}
	}
	return ""
}

// parse proto/port from
func EntryServiceToDssProtoPort(service string) (psm_dss.SecurityProtoPort, bool) {
	if IsEntryServiceDisabled(service) {
		return psm_dss.SecurityProtoPort{}, false
	}
	re, err := regexp.Compile(`([TtUu][CcDd][Pp]) (\d+)`)
	matches := re.FindStringSubmatch(service)
	if err != nil {
		return psm_dss.SecurityProtoPort{}, false
	}
	if len(matches) > 2 {
		return psm_dss.SecurityProtoPort{
			Ports:    &matches[1],
			Protocol: &matches[2],
		}, true
	}
	return psm_dss.SecurityProtoPort{}, false
}
