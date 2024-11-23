package convert

import (
	"github.com/PaloAltoNetworks/pango/poli/security"
	"regexp"
	"strings"
)

// https://github.com/PaloAltoNetworks/pango/blob/master/poli/security/entry.go
// does the entry use any any for source and destination zone?
func IsEntryZoneAnyAny(entry *security.Entry) bool {
	for _, sourceZone := range entry.SourceZones {
		for _, destinationZone := range entry.DestinationZones {
			if sourceZone == "any" && destinationZone == "any" {
				return true
			}
		}
	}
	return false
}

// does the entry use AppID?
func IsEntryAppID(entry *security.Entry) bool {
	for _, application := range entry.Applications {
		if application == "any" {
			return false
		}
	}
	return true
}

// does the entry use a user?
func IsEntryUser(entry *security.Entry) bool {
	isUser := true
	for _, user := range entry.SourceUsers {
		match, _ := regexp.MatchString("[aA]ny", user)
		if match {
			return false
		}
	}
	return isUser
}

// does the entry use a device?
func IsEntryDevice(entry *security.Entry) bool {
	isDevice := true
	for _, device := range entry.SourceDevices {
		match, _ := regexp.MatchString("[aA]ny", device)
		if match {
			return false
		}
	}
	return isDevice
}

func IsEntryDisabled(entry *security.Entry) bool {
	match, _ := regexp.MatchString("[dD]isabled", entry.Name)
	return match
}

func EntrySourceAddressGroupNames(entry *security.Entry) []string {
	addressGroupNames := make([]string, 0)
	for _, sourceAddress := range entry.SourceAddresses {
		match, _ := regexp.MatchString("[Nn]egate", sourceAddress)
		if match {
			continue
		}
		match, _ = regexp.MatchString("[Aa]ny", sourceAddress)
		if match {
			continue
		}
		match, _ = regexp.MatchString("[a-zA-Z]", sourceAddress)
		if match {
			addressGroupNames = append(addressGroupNames, sourceAddress)
		}
	}
	return addressGroupNames
}

func EntryDestinationAddressGroupNames(entry *security.Entry) []string {
	addressGroupNames := make([]string, 0)
	for _, destinationAddress := range entry.DestinationAddresses {
		match, _ := regexp.MatchString("[Nn]egate", destinationAddress)
		if match {
			continue
		}
		match, _ = regexp.MatchString("[Aa]ny", destinationAddress)
		if match {
			continue
		}
		match, _ = regexp.MatchString("[a-zA-Z]", destinationAddress)
		if match {
			addressGroupNames = append(addressGroupNames, destinationAddress)
		}
	}
	return addressGroupNames
}

type EntryCSV struct {
	Name               string `csv:"Name"`
	Tags               string `csv:"Tags"`
	Type               string `csv:"Type"`
	SourceZone         string `csv:"Source Zone"`
	SourceAddress      string `csv:"Source Address"`
	SourceUser         string `csv:"Source User"`
	SourceDevice       string `csv:"Source Device"`
	DestinationZone    string `csv:"Destination Zone"`
	DestinationAddress string `csv:"Destination Device"`
	Application        string `csv:"Application"`
	Service            string `csv:"Service"`
	Action             string `csv:"Action"`
	Profile            string `csv:"Profile"`
	Options            string `csv:"Options"`
	RuleUsageHitCount  string `csv:"Rule Usage Hit Count"`
	RuleUsageLastHit   string `csv:"Rule Usage Last Hit"`
	RuleUsageFirstHit  string `csv:"Rule Usage First Hit"`
	RuleUsageAppsSeen  string `csv:"Rule Usage Apps Seen"`
	DaysWithNoNewApps  string `csv:"Days With No New Apps"`
	Modified           string `csv:"Modified"`
	Created            string `csv:"Created"`
}

func (entryCSV *EntryCSV) Convert() (*security.Entry, error) {
	var entry security.Entry
	entry.Name = entryCSV.Name
	if entryCSV.Tags != "" {
		entry.Tags = strings.Split(entryCSV.Tags, ";")
	}
	entry.Type = entryCSV.Type
	if entryCSV.SourceZone != "" {
		entry.SourceZones = strings.Split(entryCSV.SourceZone, ";")
	}
	if entryCSV.SourceAddress != "" {
		entry.SourceAddresses = strings.Split(entryCSV.SourceAddress, ";")
	}
	if entryCSV.SourceUser != "" {
		entry.SourceUsers = strings.Split(entryCSV.SourceUser, ";")
	}
	if entryCSV.SourceDevice != "" {
		entry.SourceDevices = strings.Split(entryCSV.SourceDevice, ";")
	}
	if entryCSV.DestinationZone != "" {
		entry.DestinationZones = strings.Split(entryCSV.DestinationZone, ";")
	}
	if entryCSV.DestinationAddress != "" {
		entry.DestinationAddresses = strings.Split(entryCSV.DestinationAddress, ";")
	}
	if entryCSV.Application != "" {
		entry.Applications = strings.Split(entryCSV.Application, ";")
	}
	if entryCSV.Service != "" {
		entry.Services = strings.Split(entryCSV.Service, ";")
	}
	entry.Action = entryCSV.Action
	if entryCSV.Profile != "" {
		entry.HipProfiles = strings.Split(entryCSV.Profile, ";")
	}
	entry.LogSetting = entryCSV.Options
	return &entry, nil
}
