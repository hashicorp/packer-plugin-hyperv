// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// This step shuts down the machine. It first attempts to do so gracefully,
// but ultimately forcefully shuts it down if that fails.
//
// Uses:
//   communicator packersdk.Communicator
//   driver       Driver
//   ui           packersdk.Ui
//   vmName       string
//
// Produces:
//   <nothing>
type StepShutdown struct {
	Command         string
	Timeout         time.Duration
	DisableShutdown bool
}

func (s *StepShutdown) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {

	comm := state.Get("communicator").(packersdk.Communicator)
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packersdk.Ui)
	vmName := state.Get("vmName").(string)

	if !s.DisableShutdown {
		if s.Command != "" {
			ui.Say("Gracefully halting virtual machine...")
			log.Printf("Executing shutdown command: %s", s.Command)

			var stdout, stderr bytes.Buffer
			cmd := &packersdk.RemoteCmd{
				Command: s.Command,
				Stdout:  &stdout,
				Stderr:  &stderr,
			}
			if err := comm.Start(ctx, cmd); err != nil {
				err := fmt.Errorf("Failed to send shutdown command: %s", err)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}

		} else {
			ui.Say("Forcibly halting virtual machine...")
			if err := driver.Stop(vmName); err != nil {
				err := fmt.Errorf("Error stopping VM: %s", err)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}
	} else {
		ui.Say("Automatic shutdown disabled.")
	}

	// Wait for the machine to actually shut down
	log.Printf("Waiting max %s for shutdown to complete", s.Timeout)
	shutdownTimer := time.After(s.Timeout)

	waitRunning := make(chan bool, 1)
	go func() {
		// loop until the VM has shut down.
		for {
			running, _ := driver.IsRunning(vmName)
			if !running {
				waitRunning <- true
				return
			}

			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		// Wait for the shutdown timeout to elapse, an interrupt, or the VM to be shut down.
		select {
		case <-shutdownTimer:
			err := errors.New("Timeout while waiting for machine to shut down.")
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		case <-ctx.Done():
			// The step sequence was cancelled, so cancel the halt wait and just exit.
			log.Println("[WARN] Interrupt detected, quitting waiting for shutdown.")
			return multistep.ActionHalt
		case <-waitRunning:
			log.Println("VM shut down.")
			return multistep.ActionContinue
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (s *StepShutdown) Cleanup(state multistep.StateBag) {}
