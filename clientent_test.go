package gopen

import (
	"fmt"
	"testing"
)

var (
	PSM_ENT_CLIENT_ADDRESSES = []string{
		"psm-ent.p6o.net",
	}
	PSM_ENT_CREDS = map[string]string{
		"username": "admin",
		"password": "Pensando0$",
		"tenant":   "default",
	}
	TEST_POLICY_ENT = "any"
)

func TestPsmEntClientGetNetworkSecurityPolicyList(t *testing.T) {
	for _, clientAddress := range PSM_ENT_CLIENT_ADDRESSES {
		c, err := NewPsmEntClient(&PsmClientConfig{
			Username: PSM_ENT_CREDS["username"],
			Password: PSM_ENT_CREDS["password"],
			Hostname: clientAddress,
		})
		if err != nil {
			t.Fatal(err)
		}
		err = c.Login()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("successfully logged into Psm: %s\n", clientAddress)
		networkSecurityPolicies, err := c.GetNetworkSecurityPolicyList()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("number of network security policies: %v\n", len(networkSecurityPolicies))
		for _, policy := range networkSecurityPolicies {
			fmt.Printf("policy: %s number of rules: %v\n", *policy.Meta.Name, len(*policy.Spec.Rules))
		}
		c.Close()
	}
}
func TestPsmEntClientGetNetworkSecurityPolicy(t *testing.T) {
	for _, clientAddress := range PSM_ENT_CLIENT_ADDRESSES {
		c, err := NewPsmEntClient(&PsmClientConfig{
			Username: PSM_ENT_CREDS["username"],
			Password: PSM_ENT_CREDS["password"],
			Hostname: clientAddress,
		})
		if err != nil {
			t.Fatal(err)
		}
		err = c.Login()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("successfully logged into Psm: %s\n", clientAddress)
		networkSecurityPolicy, err := c.GetNetworkSecurityPolicy(TEST_POLICY_ENT)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("policy: %s number of rules: %v\n", *networkSecurityPolicy.Meta.Name, len(*networkSecurityPolicy.Spec.Rules))
		c.Close()
	}
}

func TestPsmEntClientGetVersion(t *testing.T) {
	for _, clientAddress := range PSM_ENT_CLIENT_ADDRESSES {
		c, err := NewPsmEntClient(&PsmClientConfig{
			Username: PSM_ENT_CREDS["username"],
			Password: PSM_ENT_CREDS["password"],
			Hostname: clientAddress,
		})
		if err != nil {
			t.Fatal(err)
		}
		err = c.Login()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("successfully logged into Psm: %s\n", clientAddress)
		version, err := c.GetVersion()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("version: %s\n", *version.Status.BuildVersion)
		c.Close()
	}
}

func TestPsmEntLogin(t *testing.T) {
	for _, clientAddress := range PSM_ENT_CLIENT_ADDRESSES {
		psmClientConfig := PsmClientConfig{
			Username: PSM_ENT_CREDS["username"],
			Password: PSM_ENT_CREDS["password"],
			Hostname: clientAddress,
			Tenant:   PSM_ENT_CREDS["tenant"],
		}
		_, err := Login(&psmClientConfig)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("successful psm ent login")
	}
}
