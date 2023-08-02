// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// This resizes the first vhd on the virtual machine to the specified DiskSize (in MB)
type StepResizeVhd struct {
	DiskSize *uint
}

func (s *StepResizeVhd) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	if s.DiskSize == nil {
		return multistep.ActionContinue
	}

	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packersdk.Ui)
	ui.Say("Resizing vhd...")

	vmName := state.Get("vmName").(string)

	// convert the MB to bytes
	newDiskSizeInBytes := uint64(*s.DiskSize) * 1024 * 1024

	err := driver.ResizeVirtualMachineVhd(vmName, newDiskSizeInBytes)
	if err != nil {
		err := fmt.Errorf("Error resizing VHD: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *StepResizeVhd) Cleanup(state multistep.StateBag) {
	// do nothing
}
