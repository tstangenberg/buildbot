package buildbot

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Buildbot/Factory", func() {

	It("creates the configuration", func() {
		c := factory.NewConfiguration()
		Expect(c).NotTo(BeNil())
	})

	It("creates the bot", func() {
		b := factory.NewBot()
		Expect(b).NotTo(BeNil())
		Expect(b.config).NotTo(BeNil())
	})

})
