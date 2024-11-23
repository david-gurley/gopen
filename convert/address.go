package convert

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type AddressGroup struct {
	Name         string   `json:"name"`
	Location     string   `json:"location"`
	MembersCount int      `json:"members_count"`
	EntryCount   int      `json:"entry_count"`
	Addresses    []string `json:"addresses"`
	Tags         []string `json:"tags"`
}

type AddressGroupCSV struct {
	Name         string `csv:"Name"`
	Location     string `csv:"Location"`
	MembersCount int    `csv:"Members Count"`
	Addresses    string `csv:"Addresses"`
	Tags         string `csv:"Tags"`
}

func (addressGroupCSV *AddressGroupCSV) Convert() (*AddressGroup, error) {
	addresses := strings.Split(addressGroupCSV.Addresses, ";")
	tags := strings.Split(addressGroupCSV.Tags, ";")
	name := addressGroupCSV.Name
	if name == "" {
		rand.Seed(time.Now().UnixNano())
		name = fmt.Sprintf("noname-%v", rand.Intn(10000))
	}
	return &AddressGroup{
		Name:         name,
		Location:     addressGroupCSV.Location,
		MembersCount: addressGroupCSV.MembersCount,
		Addresses:    addresses,
		Tags:         tags,
	}, nil
}
