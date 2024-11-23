package gopen

import (
	"github.com/pensando/go-psm/psm_dss"
	"github.com/pensando/go-psm/psm_ent"
)

func IsDscRuleCatchall(rule *psm_ent.SecuritySGRule) bool {
	match := false
	if *rule.Action != "permit" {
		return match
	}
	for _, i := range *rule.FromIpAddresses {
		if i == "any" {
			for _, ii := range *rule.ToIpAddresses {
				if ii == "any" {
					for _, p := range *rule.ProtoPorts {
						if *p.Protocol == "any" {
							match = true
							return match
						}
					}
				}
			}
		}
	}
	return match
}

func IsDssRuleCatchall(rule *psm_dss.SecuritySGRule) bool {
	match := false
	if *rule.Action != "permit" {
		return match
	}
	for _, i := range *rule.FromIpAddresses {
		if i == "any" {
			for _, ii := range *rule.ToIpAddresses {
				if ii == "any" {
					for _, p := range *rule.ProtoPorts {
						if *p.Protocol == "any" {
							match = true
							return match
						}
					}
				}
			}
		}
	}
	return match
}
