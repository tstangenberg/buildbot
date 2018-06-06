package main_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/tstangenberg/buildbot/src"
)

var _ = Describe("Main", func() {
	It("should exit normal", func() {
		main.Main()
	})
})
