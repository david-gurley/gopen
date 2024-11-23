package gopen

import (
	"fmt"

	"github.com/pensando/go-psm/psm_dss"
)

type PsmDssClient struct {
	*PsmClient
}

func NewPsmDssClient(config *PsmClientConfig) (*PsmDssClient, error) {
	client, err := NewPsmClient(config)
	if err != nil {
		return &PsmDssClient{client}, err
	}
	return &PsmDssClient{client}, nil
}

func (client *PsmDssClient) GetVersion() (*psm_dss.ClusterVersion, error) {
	var result psm_dss.ClusterVersion
	path := "/configs/cluster/v1/version"
	err := client.Get(path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *PsmDssClient) GetNetworkSecurityPolicyList() ([]psm_dss.SecurityNetworkSecurityPolicy, error) {
	var result ListResponse[psm_dss.SecurityNetworkSecurityPolicy]
	path := fmt.Sprintf("/configs/security/v1/tenant/%s/networksecuritypolicies", client.tenant)
	err := client.Get(path, nil, &result)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

func (client *PsmDssClient) GetNetworkSecurityPolicy(policyName string) (psm_dss.SecurityNetworkSecurityPolicy, error) {
	var result psm_dss.SecurityNetworkSecurityPolicy
	path := fmt.Sprintf("/configs/security/v1/tenant/%s/networksecuritypolicies/%s", client.tenant, policyName)
	err := client.Get(path, nil, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (client *PsmDssClient) GetDistributedServiceCardList() ([]psm_dss.ClusterDistributedServiceCard, error) {
	var result ListResponse[psm_dss.ClusterDistributedServiceCard]
	path := "/configs/cluster/v1/distributedservicecards"
	err := client.Get(path, nil, &result)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

func (client *PsmDssClient) GetVirtualRouterList() ([]psm_dss.NetworkVirtualRouter, error) {
	var result ListResponse[psm_dss.NetworkVirtualRouter]
	path := fmt.Sprintf("/configs/network/v1/tenant/%s/virtualrouters", client.tenant)
	err := client.Get(path, nil, &result)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}
