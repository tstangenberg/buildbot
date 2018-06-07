package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
)

var _ = Describe("Buildbot", func() {

	It("init should create a new cli", func() {
		Expect(app).ToNot(BeNil())
	})

	It("main should run app with os args", func() {
		// arrange
		mockedApplication := &ApplicationMock{
			RunFunc: func(args []string) error { return nil },
		}
		os.Args = []string{"buildbot", "--version"}
		app = mockedApplication

		// act
		main()

		// assert
		Expect(len(mockedApplication.RunCalls())).To(BeEquivalentTo(1))
		Expect(mockedApplication.RunCalls()[0].Args).To(BeEquivalentTo(os.Args))
	})

})
