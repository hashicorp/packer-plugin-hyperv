// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type SwitchConfig

package common

import (
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type SwitchConfig struct {
	// The name of the main switch to connect the virtual
	// machine to. By default, leaving this value unset will cause Packer to
	// try and determine the switch to use by looking for an external switch
	// that is up and running.
	SwitchName string `mapstructure:"switch_name" required:"false"`
	// This allows the specification of the switch type.
	// the default will be Internal if unspecificed
	SwitchType string `mapstructure:"switch_type" required:"false"`
	// This is the VLAN of the virtual switch's
	// network card. By default, none is set. If none is set then a VLAN is not
	// set on the switch's network card. If this value is set it should match
	// the VLAN specified in by vlan_id.
	SwitchVlanId string `mapstructure:"switch_vlan_id" required:"false"`
}

func (sc SwitchConfig) Prepare(ctx *interpolate.Context, pc *common.PackerConfig) ([]error, []string) {
	return nil, nil
}
