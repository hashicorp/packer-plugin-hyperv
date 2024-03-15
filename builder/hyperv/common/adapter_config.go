// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type AdapterConfig

package common

import (
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type AdapterConfig struct {
	// The name of the network adapter.
	// By default, leaving this value unset will cause Packer to
	// try and determine the switch to use by looking for an external switch
	// that is up and running.
	Name string `mapstructure:"adapter_name" required:"false"`
	// The name of the switch for this adapter
	SwitchName string `mapstructure:"switch_name" required:"false"`
	// This is the VLAN of the virtual switch's
	// network card. By default, none is set. If none is set then a VLAN is not
	// set on the switch's network card. If this value is set it should match
	// the VLAN specified in by vlan_id.
	VlanId string `mapstructure:"vlan_id" required:"false"`
	// This allows a specific MAC address to be used on
	// the default main virtual network card. The MAC address must be a string with
	// no delimiters, for example "037777777777deadbeef".
	MacAddress string `mapstructure:"mac_address" required:"false"`
}

func (ac AdapterConfig) Prepare(ctx *interpolate.Context, pc *common.PackerConfig) ([]error, []string) {
	return nil, nil
}
