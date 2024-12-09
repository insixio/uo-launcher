//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// --------------------------------------------------------------------------
// Build namespace
type Build mg.Namespace

// Build for this platform.
func (Build) Current() error {
	mg.Deps(Install.Godeps)
	platform := runtime.GOOS + "/" + runtime.GOARCH
	fmt.Printf("Building for %s...\n", platform)
	cmd := exec.Command("wails", "build", "--clean")
	return cmd.Run()
}

// Build for Mac. Accepts optional arch argument: intel, arm.
func (Build) Mac(arch ...string) error {
	mg.Deps(Install.Godeps)

	if len(arch) == 0 {
		fmt.Println("Building for Mac Universal...")
		cmd := exec.Command("wails", "build", "--clean", "--platform", "darwin/universal")
		return cmd.Run()
	}

	if arch[0] == "intel" {
		fmt.Println("Building for Mac (amd64)...")
		cmd := exec.Command("wails", "build", "--clean", "--platform", "darwin")
		return cmd.Run()
	}

	if arch[0] == "arm" {
		fmt.Println("Building for Mac (arm64)...")
		cmd := exec.Command("wails", "build", "--clean", "--platform", "darwin/arm64")
		return cmd.Run()
	}

	return fmt.Errorf("unknown architecture: %s", arch[0])
}

// Build for Windows.
// mage build:windows
func (Build) Windows() error {
	mg.Deps(Install.Godeps)
	fmt.Println("Building for Windows (amd64)...")
	cmd := exec.Command("wails", "build", "--clean", "--platform", "windows/amd64")
	return cmd.Run()
}

// Build for Linux.
// mage build:linux
func (Build) Linux() error {
	mg.Deps(Install.Godeps)
	fmt.Println("Building for Linux (amd64)...")
	cmd := exec.Command("wails", "build", "--clean", "--platform", "linux/amd64")
	return cmd.Run()
}

// --------------------------------------------------------------------------
// Install namespace
type Install mg.Namespace

// Install Wails if not installed and setup Go dependencies.
func (Install) Godeps() error {
	// check if Wails is installed
	_, err := exec.LookPath("wails")
	if err != nil {
		fmt.Println("Wails is not installed.")
		mg.Deps(Install.Wails)
	}

	// install dependencies
	cmd := exec.Command("go", "mod", "tidy")
	return cmd.Run()
}

// Install latest version of Wails.
func (Install) Wails() error {
	fmt.Println("Installing Wails...")
	cmd := exec.Command("go", "install", "github.com/wailsapp/wails/v2/cmd/wails@latest")
	return cmd.Run()
}

// --------------------------------------------------------------------------
// Dev namespace
type Dev mg.Namespace

func (Dev) Wails() error {
	fmt.Println("Running Wails in dev mode...")
	cmd := exec.Command("wails", "dev")
	return cmd.Run()
}
