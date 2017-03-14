package customercontainer_test

import (
	"context"
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
		cli.CreateMockImage("testimage")

		It("Should return intances", func() {
			_, containerID, err := cc.CreateContainer(123, &customercontainer.ContainerConfig{
				ImageName: "testimage",
			})

			instances := cc.Instances(123)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(instances).ShouldNot(BeEmpty())
			Ω(instances[0]).Should(Equal(containerID))

			cc.RemoveContainer(containerID)
		})

		It("Should return no instances when none exist", func() {
			instances := cc.Instances(123)
			Ω(instances).Should(BeEmpty())
		})
	})

	Describe("Endpoints and Transport", func() {
		cli := testutils.NewMockDCli()
		cc := customercontainer.NewService(cli)
		cli.CreateMockImage("testimage")
		es := &customercontainer.Endpoints{}
		ctx := context.Background()
		gID := ""

		It("Should create valid Endpoints", func() {
			es.CreateContainerEndpoint = customercontainer.MakeCreateContainerEndpoint(cc)
			es.EditContainerEndpoint = customercontainer.MakeEditContainerEndpoint(cc)
			es.InstancesEndpoint = customercontainer.MakeInstancesEndpoint(cc)
			es.RemoveContainerEndpoint = customercontainer.MakeRemoveContainerEndpoint(cc)
		})

		Context("CreateContainerEndpoint", func() {
			It("Should work with CreateContainer request and response", func() {
				cfg := customercontainer.ContainerConfig{
					ImageName: "testimage",
				}
				res, err := es.CreateContainerEndpoint(ctx, customercontainer.CreateContainerRequest{
					Refid: 123,
					Cfg:   &cfg,
				})

				gID = res.(customercontainer.CreateContainerResponse).ID

				Ω(err).ShouldNot(HaveOccurred())
				Ω(res.(customercontainer.CreateContainerResponse).Error).ShouldNot(HaveOccurred())
				Ω(res.(customercontainer.CreateContainerResponse).ID).ShouldNot(BeEmpty())
				Ω(res.(customercontainer.CreateContainerResponse).Name).ShouldNot(BeEmpty())
			})
		})

		Context("EditContainerEndpoint", func() {
			It("Should work with EditContainer request and response", func() {
				cfg := customercontainer.ContainerConfig{
					ImageName: "testimage",
				}
				res, err := es.EditContainerEndpoint(ctx, customercontainer.EditContainerRequest{
					ID:  gID,
					Cfg: &cfg,
				})

				Ω(err).ShouldNot(HaveOccurred())
				Ω(res.(customercontainer.EditContainerResponse).Error).ShouldNot(HaveOccurred())
			})
		})

		Context("InstancesEndpoint", func() {
			It("Should work with Instances request and response", func() {
				res, err := es.InstancesEndpoint(ctx, customercontainer.InstancesRequest{
					Refid: 123,
				})

				Ω(err).ShouldNot(HaveOccurred())
				Ω(res.(customercontainer.InstancesResponse).Instances).ShouldNot(BeNil())
			})
		})

		Context("RemoveContainerEndpoint", func() {
			It("Should work with RemoveContainer request and response", func() {

				res, err := es.RemoveContainerEndpoint(ctx, customercontainer.RemoveContainerRequest{
					ID: gID,
				})

				Ω(err).ShouldNot(HaveOccurred())
				Ω(res.(customercontainer.RemoveContainerResponse).Error).ShouldNot(HaveOccurred())
			})
		})
	})
})
