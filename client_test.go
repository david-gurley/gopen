package gopen

import (
	"fmt"
	"testing"
)

var PSM_VERSIONS = []string{
	"1.29.1-T-2",
	"1.28.1-E-6",
}

func TestPsmKind(t *testing.T) {
	psmBuild, err := NewPsmBuild()
	if err != nil {
		t.Fatal(err)
	}
	for _, version := range PSM_VERSIONS {
		kind := psmBuild.Kind(version)
		fmt.Printf("version: %s kind: %s\n", version, kind)
	}
}
