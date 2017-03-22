package network_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kontainerooo/kontainer.ooo/pkg/customercontainer"
	"github.com/kontainerooo/kontainer.ooo/pkg/network"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"
)

var _ = Describe("Network", func() {
	Describe("Create service", func() {
		It("Should create service", func() {
			networkService, err := network.NewService(testutils.NewMockDCli(), testutils.NewMockDB())
			Ω(networkService).ShouldNot(BeNil())
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should return db error", func() {
			db := testutils.NewMockDB()
			db.SetError(1)
			_, err := network.NewService(testutils.NewMockDCli(), db)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Create Network", func() {
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		networkService, _ := network.NewService(dcli, db)
		It("Should create a new network", func() {
			name, id, err := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			Ω(err).ShouldNot(HaveOccurred())
			Ω(id).ShouldNot(BeEmpty())
			Ω(name).ShouldNot(BeEmpty())

			Ω(len(dcli.GetNetworks())).Should(Equal(1))
		})

		It("Should error on client network generation", func() {
			dcli.SetError()
			name, id, err := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			Ω(err).Should(HaveOccurred())
			Ω(id).Should(BeEmpty())
			Ω(name).Should(BeEmpty())

			Ω(len(dcli.GetNetworks())).Should(Equal(1))
		})

		It("Should receive a database error", func() {
			db.SetError(1)
			name, id, err := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			Ω(err).Should(HaveOccurred())
			Ω(id).Should(BeEmpty())
			Ω(name).Should(BeEmpty())

			Ω(len(dcli.GetNetworks())).Should(Equal(1))
		})
	})

	Describe("Remove Network", func() {
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		networkService, _ := network.NewService(dcli, db)
		It("Should remove a network", func() {
			name, _, _ := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			err := networkService.RemoveNetworkByName(123, name)

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error when network cannot be removed", func() {
			name, _, _ := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			dcli.SetError()

			err := networkService.RemoveNetworkByName(123, name)

			Ω(err).Should(HaveOccurred())
		})

		It("Should error when network is not found in db", func() {
			name, _, _ := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			db.SetError(1)

			err := networkService.RemoveNetworkByName(123, name)

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Add Container to network", func() {
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		ccs := customercontainer.NewService(dcli)
		networkService, _ := network.NewService(dcli, db)
		dcli.CreateMockImage("testimage")
		containerID, _, _ := ccs.CreateContainer(123, &customercontainer.ContainerConfig{
			ImageName: "testimage",
		})

		It("Should add a container to a network", func() {
			name, _, _ := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			err := networkService.AddContainerToNetwork(123, name, containerID)

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error when network does not exist in db", func() {
			err := networkService.AddContainerToNetwork(123, "idontexist", containerID)

			Ω(err).Should(HaveOccurred())
		})

		It("Should error when network does not exist on host", func() {
			name, _, _ := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			dcli.SetError()

			err := networkService.AddContainerToNetwork(123, name, containerID)

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Remove container from network", func() {
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		ccs := customercontainer.NewService(dcli)
		networkService, _ := network.NewService(dcli, db)
		dcli.CreateMockImage("testimage")
		containerID, _, _ := ccs.CreateContainer(123, &customercontainer.ContainerConfig{
			ImageName: "testimage",
		})

		It("Should remove a container from a network", func() {
			name, _, _ := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			_ = networkService.AddContainerToNetwork(123, name, containerID)

			err := networkService.RemoveContainerFromNetwork(123, name, containerID)

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error when container cannot be removed from network", func() {
			name, _, _ := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			_ = networkService.AddContainerToNetwork(123, name, containerID)

			dcli.SetError()

			err := networkService.RemoveContainerFromNetwork(123, name, containerID)

			Ω(err).Should(HaveOccurred())
		})

		It("Should error when network is not found in db", func() {
			name, _, _ := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			_ = networkService.AddContainerToNetwork(123, name, containerID)

			db.SetError(1)

			err := networkService.RemoveContainerFromNetwork(123, name, containerID)

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Check if user has a network", func() {
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		ccs := customercontainer.NewService(dcli)
		networkService, _ := network.NewService(dcli, db)
		dcli.CreateMockImage("testimage")
		_, _, _ = ccs.CreateContainer(123, &customercontainer.ContainerConfig{
			ImageName: "testimage",
		})
		nwame := ""

		It("Should return true when user has a network", func() {
			nwame, _, _ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
			})

			b := networkService.UserHasNetwork(123)

			Ω(b).Should(BeTrue())
		})

		It("Should return false when user does not have a network", func() {
			networkService.RemoveNetworkByName(123, nwame)

			b := networkService.UserHasNetwork(123)

			Ω(b).ShouldNot(BeTrue())
		})

		It("Should return false when DB errors", func() {
			db.SetError(1)

			b := networkService.UserHasNetwork(123)

			Ω(b).ShouldNot(BeTrue())
		})
	})
})
