package containerlifecycle_test

import (
	"context"

	"github.com/kontainerooo/kontainer.ooo/pkg/containerlifecycle"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Container Lifecycle", func() {
	Describe("Create Service", func() {
		It("Should create Service", func() {
			cls := containerlifecycle.NewService(testutils.NewMockDCli())
			Expect(cls).ToNot(BeZero())
		})
	})

	Describe("Start Container", func() {
		dcli := testutils.NewMockDCli()
		cls := containerlifecycle.NewService(dcli)
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
		cls := containerlifecycle.NewService(dcli)
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
		cls := containerlifecycle.NewService(dcli)
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

	Describe("Endpoints and Transport", func() {
		dcli := testutils.NewMockDCli()
		cls := containerlifecycle.NewService(dcli)
		es := &containerlifecycle.Endpoints{}
		ctx := context.Background()
		It("Should create valid Endpoints", func() {
			es.StartCommandEndpoint = containerlifecycle.MakeStartCommandEndpoint(cls)
			es.StartContainerEndpoint = containerlifecycle.MakeStartContainerEndpoint(cls)
			es.StopContainerEndpoint = containerlifecycle.MakeStopContainerEndpoint(cls)
		})

		Context("StartContainerEndpoint", func() {
			It("Should work with StartContainer request and response", func() {
				res, err := es.StartContainerEndpoint(ctx, containerlifecycle.StartContainerRequest{
					ID: "123",
				})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(res.(containerlifecycle.StartContainerResponse).Error).ShouldNot(HaveOccurred())
			})
		})

		Context("StartCommandEndpoint", func() {
			It("Should work with StartCommand request and response", func() {
				res, err := es.StartCommandEndpoint(ctx, containerlifecycle.StartCommandRequest{
					ID:  "123",
					Cmd: "/bin/sh",
				})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(res.(containerlifecycle.StartCommandResponse).ID).ShouldNot(BeZero())
				Ω(res.(containerlifecycle.StartCommandResponse).Error).ShouldNot(HaveOccurred())
			})
		})

		Context("StopContainerEndpoint", func() {
			It("Should work with StopContainer request and response", func() {
				res, err := es.StopContainerEndpoint(ctx, containerlifecycle.StopContainerRequest{
					ID: "123",
				})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(res.(containerlifecycle.StopContainerResponse).Error).ShouldNot(HaveOccurred())
			})
		})
	})
})
