package gopen

import (
	"fmt"
	"testing"
)

var (
	PSM_DSS_CLIENT_ADDRESSES = []string{
		"psm-t.p6o.net",
	}
	TEST_POLICY_DSS = "zone1-intra"
	PSM_DSS_CREDS   = map[string]string{
		"username": "admin",
		"password": "Pensando0$",
		"tenant":   "default",
	}
)

func TestPsmDssClientGetNetworkSecurityPolicyList(t *testing.T) {
	for _, clientAddress := range PSM_DSS_CLIENT_ADDRESSES {
		c, err := NewPsmDssClient(&PsmClientConfig{
			Username: PSM_DSS_CREDS["username"],
			Password: PSM_DSS_CREDS["password"],
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

func TestPsmDssClientGetDistributedServiceCardList(t *testing.T) {
	for _, clientAddress := range PSM_DSS_CLIENT_ADDRESSES {
		c, err := NewPsmDssClient(&PsmClientConfig{
			Username: PSM_DSS_CREDS["username"],
			Password: PSM_DSS_CREDS["password"],
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
		distributedServiceCards, err := c.GetDistributedServiceCardList()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("number of distributed service cards: %v\n", len(distributedServiceCards))
		c.Close()
	}
}

func TestPsmDssClientGetVirtaulRouterList(t *testing.T) {
	for _, clientAddress := range PSM_DSS_CLIENT_ADDRESSES {
		c, err := NewPsmDssClient(&PsmClientConfig{
			Username: PSM_DSS_CREDS["username"],
			Password: PSM_DSS_CREDS["password"],
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
		virtualRouters, err := c.GetVirtualRouterList()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("number of virtual routers: %v\n", len(virtualRouters))
		c.Close()
	}
}

func TestPsmDssClientGetNetworkSecurityPolicy(t *testing.T) {
	for _, clientAddress := range PSM_DSS_CLIENT_ADDRESSES {
		c, err := NewPsmDssClient(&PsmClientConfig{
			Username: PSM_DSS_CREDS["username"],
			Password: PSM_DSS_CREDS["password"],
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
		networkSecurityPolicy, err := c.GetNetworkSecurityPolicy(TEST_POLICY_DSS)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("policy: %s number of rules: %v\n", *networkSecurityPolicy.Meta.Name, len(*networkSecurityPolicy.Spec.Rules))
		c.Close()
	}
}

func TestPsmDssClientGetVersion(t *testing.T) {
	for _, clientAddress := range PSM_DSS_CLIENT_ADDRESSES {
		c, err := NewPsmDssClient(&PsmClientConfig{
			Username: PSM_DSS_CREDS["username"],
			Password: PSM_DSS_CREDS["password"],
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

func TestPsmDssLogin(t *testing.T) {
	for _, clientAddress := range PSM_DSS_CLIENT_ADDRESSES {
		psmClientConfig := PsmClientConfig{
			Username: PSM_DSS_CREDS["username"],
			Password: PSM_DSS_CREDS["password"],
			Hostname: clientAddress,
			Tenant:   PSM_DSS_CREDS["tenant"],
		}
		_, err := Login(&psmClientConfig)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("successful psm dss login")
	}
}
