package gopen

import (
	"fmt"

	"github.com/pensando/go-psm/psm_ent"
)

type PsmEntClient struct {
	*PsmClient
}

func NewPsmEntClient(config *PsmClientConfig) (*PsmEntClient, error) {
	client, err := NewPsmClient(config)
	if err != nil {
		return &PsmEntClient{}, err
	}
	return &PsmEntClient{client}, nil
}
func (client *PsmEntClient) GetVersion() (*psm_ent.ClusterVersion, error) {
	var result psm_ent.ClusterVersion
	path := "/configs/cluster/v1/version"
	err := client.Get(path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *PsmEntClient) GetNetworkSecurityPolicyList() ([]psm_ent.SecurityNetworkSecurityPolicy, error) {
	var result ListResponse[psm_ent.SecurityNetworkSecurityPolicy]
	path := fmt.Sprintf("/configs/security/v1/tenant/%s/networksecuritypolicies", client.tenant)
	err := client.Get(path, nil, &result)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}
func (client *PsmEntClient) GetNetworkSecurityPolicy(policyName string) (psm_ent.SecurityNetworkSecurityPolicy, error) {
	var result psm_ent.SecurityNetworkSecurityPolicy
	path := fmt.Sprintf("/configs/security/v1/tenant/%s/networksecuritypolicies/%s", client.tenant, policyName)
	err := client.Get(path, nil, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
