// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

const (
	SwitchTypeExternal = "External"
	SwitchTypeInternal = "Internal"
	SwitchTypePrivate  = "Private"
	DefaultSwitchType  = SwitchTypeInternal
)

// This step creates switch for VM.
//
// Produces:
//
//	SwitchConfigs []SwitchConfigs - The new swichConfigs
type StepCreateSwitches struct {
	// Specifies the switches to be created.
	SwitchConfigs []SwitchConfig

	createdSwitch []bool
}

func (s *StepCreateSwitches) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packersdk.Ui)

	for index, sw := range s.SwitchConfigs {
		ui.Say(fmt.Sprintf("Creating switch '%v' if required...", sw.SwitchName))
		if len(sw.SwitchType) == 0 {
			sw.SwitchType = DefaultSwitchType
		}
		var createdSwitch bool
		var err error
		if sw.SwitchType == SwitchTypeExternal {
			createdSwitch, err = driver.CreateExternalVirtualSwitch(sw.SwitchName)
		} else {
			createdSwitch, err = driver.CreateVirtualSwitch(sw.SwitchName, sw.SwitchType)
		}
		if err != nil {
			verr := fmt.Errorf("Error creating switch: %s: %v", sw.SwitchName, err)
			state.Put("error", verr)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		if index == 0 {
			state.Put("swName", sw.SwitchName)
		}
		s.createdSwitch = append(s.createdSwitch, createdSwitch)

		if !s.createdSwitch[index] {
			ui.Say(fmt.Sprintf("    switch '%v' already exists. Will not delete on cleanup...", sw.SwitchName))
		}
	}

	return multistep.ActionContinue
}

func (s *StepCreateSwitches) Cleanup(state multistep.StateBag) {
	if len(s.SwitchConfigs) == 0 {
		return
	}

	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packersdk.Ui)
	ui.Say("Unregistering and deleting switch...")

	for index, sw := range s.SwitchConfigs {
		if s.createdSwitch[index] {
			err := driver.DeleteVirtualSwitch(sw.SwitchName)
			if err != nil {
				ui.Error(fmt.Sprintf("Error deleting switch: %s", err))
			}
		}
	}
}
