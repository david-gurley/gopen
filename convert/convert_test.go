package convert

import (
	"fmt"
	"testing"
)

const IMPORT_ENTRIES_FILE = "importentries.csv"
const IMPORT_ADDRESS_GROUPS_FILE = "importaddressgroups.csv"
const IMPORT_ZONES_FILE = "importzones.csv"

func TestOpenFile(t *testing.T) {
	config := ConfigFromCSVFiles{
		Entries:       IMPORT_ENTRIES_FILE,
		AddressGroups: IMPORT_ADDRESS_GROUPS_FILE,
		Zones:         IMPORT_ZONES_FILE,
	}

	converted, err := NewFromCSVFiles(&config)
	if err != nil {
		t.Fatal(err)
	}
	for _, zone := range converted.Zones {
		fmt.Printf("zone: %s used: %v\n", zone.Name, zone.EntryCount)
	}
	for _, addressGroup := range converted.AddressGroups {
		fmt.Printf("addressGroup: %s used: %v\n", addressGroup.Name, addressGroup.EntryCount)
	}
	converted.ToDss()
	for _, rule := range *converted.DssPolicy[DEFAULT_POLICY].Spec.Rules {
		fmt.Printf("rule: %#v\n", rule)
	}
}
