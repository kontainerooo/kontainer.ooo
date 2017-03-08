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
			containerName, _, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
				ImageName: "testimage",
			})

			Ω(err).ShouldNot(HaveOccurred())
			Ω(strings.HasPrefix(containerName, "123")).Should(BeTrue())
		})

		It("Should fail with missing container image", func() {
			cli := testutils.NewMockDCli()
			cc := customercontainer.NewService(cli)

			containerName, _, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
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
			containerName, _, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
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
			containerName, _, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
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

			containerName, _, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
				ImageName: "testimage",
			})
			fmt.Println(err)

			Ω(err).Should(HaveOccurred())
			Ω(containerName).Should(BeZero())

			customercontainer.SeccompProfile = tmpSeccomp
		})
	})

	Describe("Edit Container", func() {
		cli := testutils.NewMockDCli()
		cc := customercontainer.NewService(cli)
		It("Should edit container", func() {
			err := cc.EditContainer("123", &customercontainer.ContainerConfig{})

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("Remove container", func() {
		cli := testutils.NewMockDCli()
		cc := customercontainer.NewService(cli)
		cli.CreateMockImage("testimage")
		_, containerID, _ := cc.CreateContainer(123, &customercontainer.ContainerConfig{
			ImageName: "testimage",
		})
		It("Should remove container", func() {
			err := cc.RemoveContainer(containerID)

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should fail when container does not exist", func() {
			err := cc.RemoveContainer(containerID)

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Get instances", func() {
		cli := testutils.NewMockDCli()
		cc := customercontainer.NewService(cli)
		It("Should return intances", func() {
			instances := cc.Instances(123)

			Ω(instances).Should(BeEmpty())
		})
	})
})
