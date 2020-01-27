package kea

import "errors"

type Hosts []*Host

type Host struct {
	Address        string             `json:"address"`
	Hostname       string             `json:"hostname"`
	ID             int64              `json:"id"`
	Identifier     string             `json:"identifier"`
	identifierType HostIdentifierType `json:"identifierTypeId"`
	subnetID       int64              `json:"subnetId"`
}

type SaveHost struct {
	Address        string             `json:"address"`
	Hostname       string             `json:"hostname"`
	Identifier     string             `json:"identifier"`
	IdentifierType HostIdentifierType `json:"identifierType"`
	SubnetID       int64              `json:"subnetId"`
}

type Subnets []*Subnet

type Subnet struct {
	ID       int32       `json:"id"`
	Pools    SubnetPools `json:"pools"`
	Prefix   string      `json:"prefix"`
	ServerID int32       `json:"serverId"`
}

type SubnetPools []*SubnetPool

type SubnetPool struct {
	ID           int32  `json:"id"`
	endAddress   string `json:"endAddress"`
	hosts        Hosts  `json:"hosts"`
	startAddress string `json:"startAddress"`
}

type HostIdentifierType string

func (lt HostIdentifierType) IsValid() error {
	switch lt {
	case HwAddress, DUID, CircuitId, ClientId, FlexId:
		return nil
	}
	return errors.New("Inalid leave type")
}

const (
	HwAddress = "HW_ADDRESS"
	DUID      = "DUID"
	CircuitId = "CIRCUIT_ID"
	ClientId  = "CLIENT_ID"
	FlexId    = "FLEX_ID"
)
