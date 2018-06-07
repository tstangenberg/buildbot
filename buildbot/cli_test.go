package buildbot_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/tstangenberg/buildbot/buildbot"
)

var _ = Describe("Cli", func() {

	app := NewCli()

	Describe("has correct", func() {
		It("name", func() {
			Expect(app.Name).To(BeEquivalentTo("Build Bot"))
		})
		It("usage", func() {
			Expect(app.Usage).To(BeEquivalentTo("#1 Software Development Automation Robot"))
		})
	})

})
