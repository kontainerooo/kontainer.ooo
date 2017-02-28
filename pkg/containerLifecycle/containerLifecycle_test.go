package containerLifecycle_test

import (
	"github.com/ttdennis/kontainer.io/pkg/containerLifecycle"
	"github.com/ttdennis/kontainer.io/pkg/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Container Lifecycle", func() {
	Describe("Create Service", func() {
		It("Should create Service", func() {
			cls := containerLifecycle.NewService(testutils.NewMockDCli())
			Expect(cls).ToNot(BeZero())
		})
	})

	Describe("Start Container", func() {
		dcli := testutils.NewMockDCli()
		cls := containerLifecycle.NewService(dcli)
		It("Should start a container", func() {
			container := "test"
			err := cls.StartContainer(container)
			Ω(dcli.IsRunning(container)).Should(BeTrue())
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should return an error if the container won't start", func() {
			container := "test"
			dcli.SetError()
			err := cls.StartContainer(container)
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if the container is already started", func() {
			container := "test"
			dcli.SetError()
			cls.StartContainer(container)
			err := cls.StartContainer(container)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Execute Command", func() {
		dcli := testutils.NewMockDCli()
		cls := containerLifecycle.NewService(dcli)
		container := "test"
		cls.StartContainer(container)

		It("Should start a command", func() {
			command := "/bin/sh -c grep"
			cmd, err := cls.StartCommand(container, command)

			Ω(cmd).Should(Equal(command))
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should return an error if command won't execute", func() {
			command := "/bin/sh -c grep"
			dcli.SetError()
			cmd, err := cls.StartCommand(container, command)

			Ω(cmd).Should(BeEmpty())
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if container is not started", func() {
			container := "test2"
			command := "/bin/sh -c grep"
			cmd, err := cls.StartCommand(container, command)

			Ω(cmd).Should(BeEmpty())
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Stop Container", func() {
		dcli := testutils.NewMockDCli()
		cls := containerLifecycle.NewService(dcli)
		container := "test"
		It("Should stop a container", func() {
			cls.StartContainer(container)
			err := cls.StopContainer(container)

			Ω(dcli.IsRunning(container)).ShouldNot(BeTrue())
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should return an error if container is not running", func() {
			err := cls.StopContainer(container)
			Ω(err).Should(HaveOccurred())
		})
	})
})
