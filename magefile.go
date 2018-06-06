// +build mage

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

const app = "buildbot"

// A build step that requires additional params, or platform specific steps for example
func Build() {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	startAndLog(exec.Command("go", "build", "-o", app, "-v", "./src/..."), "build")
}

// Manage your deps, or running package managers.
func InstallDeps() {
	fmt.Println("Installing Deps...")
	startAndLog(exec.Command("dep", "ensure", "-update"), "deps")
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("MyApp")
}

// Install required tools.
func Setup() {
	fmt.Println("Running Development Setup...")
	startAndLog(exec.Command("dep", "ensure", "-update"), "setup")
}

func startAndLog(cmd *exec.Cmd, logprefix string) {
	startAndLogWithOutAndErr(cmd, os.Stdout, os.Stderr, logprefix)
}

func startAndLogWithOutAndErr(cmd *exec.Cmd, stdout *os.File, stderr *os.File, logprefix string) {
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Fprintf(stdout, "%s | %s\n", logprefix, scanner.Text())
		}
	}()
	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(stderr, "Error starting Cmd", err)
		os.Exit(1)
	}
}

