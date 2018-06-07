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

// A build step that requires additional params, or platform specific steps for example
func Build() {
	mg.Deps(InstallDeps)
	mg.Deps(Generate)
	fmt.Println("Building...")
	startAndLog(exec.Command("go", "build", "-o", "./out/buildbot", "-v", "./cmd/buildbot"), "build")
}

func Generate() {
	mg.Deps(InstallDeps)
	fmt.Println("Generating...")
	startAndLog(exec.Command("go", "generate", "-v", "./cmd/..."), "build")
}

func Test() {
	mg.Deps(Generate)
	fmt.Println("Running Tests...")
	startAndLog(exec.Command("go", "test", "-v", "./cmd/..."), "build")
	startAndLog(exec.Command("go", "test", "-v", "./buildbot/..."), "build")
}


// Manage your deps, or running package managers.
func InstallDeps() {
	fmt.Println("Installing Deps...")
	startAndLog(exec.Command("dep", "ensure", "-update"), "deps")
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("out")
}

// Install required tools.
func Setup() {
	fmt.Println("Running Development Setup...")
	goGet("github.com/golang/dep/cmd/dep")
	goGet("github.com/onsi/ginkgo/ginkgo")
	goGet("github.com/onsi/gomega/...")
	goGet("github.com/matryer/moq")
}

func goGet(name string) {
	fmt.Printf("Getting %s ...\n", name)
	startAndLog(exec.Command("go", "get", "-u", "-v", name), "go get")
}

func startAndLog(cmd *exec.Cmd, logprefix string) {
	startAndLogWithOutAndErr(cmd, os.Stdout, os.Stderr, logprefix)
}

func startAndLogWithOutAndErr(cmd *exec.Cmd, stdout *os.File, stderr *os.File, logprefix string) {
	outReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}
	outScanner := bufio.NewScanner(outReader)
	go func() {
		for outScanner.Scan() {
			fmt.Fprintf(stdout, "%s | %s\n", logprefix, outScanner.Text())
		}
	}()
	errReader, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(stderr, "Error creating StderrPipe for Cmd", err)
		os.Exit(1)
	}
	errScanner := bufio.NewScanner(errReader)
	go func() {
		for errScanner.Scan() {
			fmt.Fprintf(stderr, "%s | %s\n", logprefix, errScanner.Text())
		}
	}()
	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(stderr, "Error starting Cmd", err)
		os.Exit(1)
	}
}

