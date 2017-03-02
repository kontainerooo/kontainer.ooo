package customercontainer_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ttdennis/kontainer.io/pkg/customercontainer"
	"github.com/ttdennis/kontainer.io/pkg/testutils"
)

var _ = Describe("Customercontainer", func() {

	Describe("Create service", func() {
		It("Should create a new service", func() {
			cc := customercontainer.NewService(testutils.NewMockDCli())
			Ω(cc).ShouldNot(BeZero())
		})
	})

	Describe("Create Container", func() {
		It("Should create a new container", func() {
			cli := testutils.NewMockDCli()
			cc := customercontainer.NewService(cli)
			cli.CreateMockImage("testimage")
			containerName, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
				ImageName: "testimage",
			})

			Ω(err).ShouldNot(HaveOccurred())
			Ω(strings.HasPrefix(containerName, "123")).Should(BeTrue())
		})

		It("Should fail with missing container image", func() {
			cli := testutils.NewMockDCli()
			cc := customercontainer.NewService(cli)

			containerName, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
				ImageName: "testimage",
			})

			Ω(err).Should(HaveOccurred())
			Ω(containerName).Should(BeZero())
		})

		It("Should fail creating the docker container", func() {
			cli := testutils.NewMockDCli()
			cc := customercontainer.NewService(cli)
			cli.CreateMockImage("testimage")

			cli.SetDockerOffline()
			containerName, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
				ImageName: "testimage",
			})

			Ω(err).Should(HaveOccurred())
			Ω(containerName).Should(BeZero())
		})

		It("Should fail on renaming the container", func() {
			cli := testutils.NewMockDCli()
			cc := customercontainer.NewService(cli)
			cli.CreateMockImage("testimage")

			cli.SetIDNotExisting()
			containerName, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
				ImageName: "testimage",
			})

			Ω(err).Should(HaveOccurred())
			Ω(containerName).Should(BeZero())
		})

		It("Should fail decoding seccomp profile", func() {
			cli := testutils.NewMockDCli()
			cc := customercontainer.NewService(cli)
			cli.CreateMockImage("testimage")

			// Save seccomp and remove
			tmpSeccomp := customercontainer.SeccompProfile
			customercontainer.SeccompProfile = ``

			containerName, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
				ImageName: "testimage",
			})
			fmt.Println(err)

			Ω(err).Should(HaveOccurred())
			Ω(containerName).Should(BeZero())

			customercontainer.SeccompProfile = tmpSeccomp
		})
	})
})
