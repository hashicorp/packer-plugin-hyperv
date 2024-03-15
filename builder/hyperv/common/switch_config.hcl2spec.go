// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatSwitchConfig is an auto-generated flat version of SwitchConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatSwitchConfig struct {
	SwitchName   *string `mapstructure:"switch_name" required:"false" cty:"switch_name" hcl:"switch_name"`
	SwitchType   *string `mapstructure:"switch_type" required:"false" cty:"switch_type" hcl:"switch_type"`
	SwitchVlanId *string `mapstructure:"switch_vlan_id" required:"false" cty:"switch_vlan_id" hcl:"switch_vlan_id"`
}

// FlatMapstructure returns a new FlatSwitchConfig.
// FlatSwitchConfig is an auto-generated flat version of SwitchConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*SwitchConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatSwitchConfig)
}

// HCL2Spec returns the hcl spec of a SwitchConfig.
// This spec is used by HCL to read the fields of SwitchConfig.
// The decoded values from this spec will then be applied to a FlatSwitchConfig.
func (*FlatSwitchConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"switch_name":    &hcldec.AttrSpec{Name: "switch_name", Type: cty.String, Required: false},
		"switch_type":    &hcldec.AttrSpec{Name: "switch_type", Type: cty.String, Required: false},
		"switch_vlan_id": &hcldec.AttrSpec{Name: "switch_vlan_id", Type: cty.String, Required: false},
	}
	return s
}
