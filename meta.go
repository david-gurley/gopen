package gopen

import (
	"sync"

	"github.com/pensando/go-psm/psm_dss"
	"github.com/pensando/go-psm/psm_ent"
)

type PsmEntMetaMap struct {
	Mu *sync.Mutex
	M  map[string]*PsmEntMeta
}

func NewPsmEntMetaMap() *PsmEntMetaMap {
	return &PsmEntMetaMap{Mu: new(sync.Mutex), M: make(map[string]*PsmEntMeta)}
}

type PsmEntMeta struct {
	NetworkSecurityPolicies []psm_ent.SecurityNetworkSecurityPolicy
	DistributedServiceCards []psm_ent.ClusterDistributedServiceCard
}

type PsmDssMetaMap struct {
	Mu *sync.Mutex
	M  map[string]*PsmDssMeta
}

func NewPsmDssMetaMap() *PsmDssMetaMap {
	return &PsmDssMetaMap{Mu: new(sync.Mutex), M: make(map[string]*PsmDssMeta)}
}

type PsmDssMeta struct {
	NetworkSecurityPolicies []psm_dss.SecurityNetworkSecurityPolicy
	DistributedServiceCards []psm_dss.ClusterDistributedServiceCard
	VirtualRouters          []psm_dss.NetworkVirtualRouter
}
