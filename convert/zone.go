package convert

import (
	"strings"
)

type Zone struct {
	Name        string   `json:"name'`
	Addresses   []string `json:"addresses"`
	Description string   `json:"description"`
	EntryCount  int      `json:"entry_count"`
}
type ZoneCSV struct {
	Name        string `csv:"name"`
	Addresses   string `csv:"addresses"`
	Description string `csv:"description"`
}

func (zoneCSV *ZoneCSV) Convert() (*Zone, error) {
	addresses := strings.Split(zoneCSV.Addresses, ";")
	return &Zone{
		Name:        zoneCSV.Name,
		Addresses:   addresses,
		Description: zoneCSV.Description,
		EntryCount:  0,
	}, nil
}
