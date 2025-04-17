// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepConfigureAdapters struct {
	UseLegacyNetworkAdapter bool
	AdapterConfigs          []AdapterConfig
}

func (s *StepConfigureAdapters) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	ui.Say("Configuring network adapters...")

	driver := state.Get("driver").(Driver)
	vmName := state.Get("vmName").(string)

	errorMsg := "Error configuring adapter: %s(%d): %v"

	for ii, adp := range s.AdapterConfigs {
		ui.Say(fmt.Sprintf("Building network adapter %s(%d)...", adp.Name, ii))

		if ii == 0 && s.UseLegacyNetworkAdapter {
			err := driver.ReplaceVirtualMachineNetworkAdapter(adp.Name, true)
			if err != nil {
				err := fmt.Errorf("Error creating legacy network adapter: %s", err)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		} else if ii != 0 {
			err := driver.CreateVirtualMachineNetworkAdapter(vmName, adp.Name, adp.SwitchName, s.UseLegacyNetworkAdapter)
			if err != nil {
				err := fmt.Errorf("Error creating network adapter: %s(%d): %v", adp.Name, ii, err)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}

		if adp.MacAddress != "" {
			aerr := driver.SetVmNetworkAdapterMacAddress(adp.Name, adp.MacAddress)
			if aerr != nil {
				err := fmt.Errorf("Error setting MAC address: %s(%d): %v", adp.Name, ii, aerr)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}

		if adp.VlanId != "" {
			err := driver.SetNetworkAdapterVlanId(adp.Name, adp.VlanId)
			if err != nil {
				err := fmt.Errorf(errorMsg, adp.Name, ii, err)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}

		if ii == 0 && adp.VlanId != "" {
			err := driver.SetVirtualMachineVlanId(vmName, adp.VlanId)
			if err != nil {
				err := fmt.Errorf(errorMsg, adp.Name, ii, err)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}
	}

	return multistep.ActionContinue
}

func (s *StepConfigureAdapters) Cleanup(state multistep.StateBag) {
	//do nothing
}
