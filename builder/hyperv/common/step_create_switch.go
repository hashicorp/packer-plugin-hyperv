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
	SwitchTypeInternal = "Internal"
	SwitchTypePrivate  = "Private"
	DefaultSwitchType  = SwitchTypeInternal
)

// This step creates switch for VM.
//
// Produces:
//
//	SwitchName string - The name of the Switch
type StepCreateSwitches struct {
	// Specifies the name of the switch to be created.
	MainSwitchName string
	SwitchesNames  []string
	// Specifies the type of the switch to be created. Allowed values are Internal and Private. To create an External
	// virtual switch, specify either the NetAdapterInterfaceDescription or the NetAdapterName parameter, which
	// implicitly set the type of the virtual switch to External.
	SwitchType string
	// Specifies the name of the network adapter to be bound to the switch to be created.
	NetAdapterName string
	// Specifies the interface description of the network adapter to be bound to the switch to be created.
	NetAdapterInterfaceDescription string

	createdSwitch []bool
}

func (s *StepCreateSwitches) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packersdk.Ui)

	if len(s.SwitchType) == 0 {
		s.SwitchType = DefaultSwitchType
	}

	switches := append([]string{s.MainSwitchName}, s.SwitchesNames...)

	for index, switchName := range switches {
		ui.Say(fmt.Sprintf("Creating switch '%v' if required...", switchName))
		createdSwitch, err := driver.CreateVirtualSwitch(switchName, s.SwitchType)
		if err != nil {
			err := fmt.Errorf("Error creating switch: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		s.createdSwitch = append(s.createdSwitch, createdSwitch)
		fmt.Sprintf("switches: '%v'", s.createdSwitch)

		if !s.createdSwitch[index] {
			ui.Say(fmt.Sprintf("    switch '%v' already exists. Will not delete on cleanup...", switchName))
		}
	}

	// Set the final name in the state bag so others can use it
	state.Put("SwitchName", s.MainSwitchName)
	state.Put("SwitchesNames", s.SwitchesNames)

	return multistep.ActionContinue
}

func (s *StepCreateSwitches) Cleanup(state multistep.StateBag) {
	if s.MainSwitchName == "" && len(s.SwitchesNames) == 0 {
		return
	}

	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packersdk.Ui)
	ui.Say("Unregistering and deleting switch...")

	switches := append([]string{s.MainSwitchName}, s.SwitchesNames...)

	for index, switchName := range switches {
		if s.createdSwitch[index] {
			err := driver.DeleteVirtualSwitch(switchName)
			if err != nil {
				ui.Error(fmt.Sprintf("Error deleting switch: %s", err))
			}
		}
	}
}
