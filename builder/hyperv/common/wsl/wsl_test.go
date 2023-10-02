// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wsl

import (
	"os"
	"testing"
)

func TestGetWSlTemp(t *testing.T) {

	if !IsWSL() {
		t.Skipf("not running in WSL")
		return
	}
	tempDir, err := GetWSlTemp()
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if tempDir == "" {
		t.Fatalf("tempDir is not polulated correctly")
	}
}

func TestConvertWindowsPathToWSlPath(t *testing.T) {

	if !IsWSL() {
		t.Skipf("not running in WSL")
		return
	}
	wslPath, err := ConvertWindowsPathToWSlPath("C:\\Users\\User\\path with spaces")
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if wslPath == "" {
		t.Fatalf("wslPath is not polulated correctly")
	}
}

func TestConvertWSlPathToWindowsPath(t *testing.T) {

	if !IsWSL() {
		t.Skipf("not running in WSL")
		return
	}
	curDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd should not have error: %s", err)
	}
	winPath, err := ConvertWSlPathToWindowsPath(curDir)
	if err != nil {
		t.Fatalf("ConvertWSlPathToWindowsPath should not have error: %s", err)
	}
	if winPath == "" {
		t.Fatalf("wslPath is not polulated correctly")
	}
}
