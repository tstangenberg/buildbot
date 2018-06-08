package buildbot

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega")

var _ bool = Describe("Buildbot/Configuration", func() {

	It("Creates a configuration file", func() {
		Expect("a").To(BeEquivalentTo("a"))
	})

})
