package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBuildbot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Buildbot Suite")
}
