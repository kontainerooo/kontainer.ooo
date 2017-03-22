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
			err := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network1",
			})

			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(dcli.GetNetworks())).Should(Equal(1))
		})

		It("Should error on client network generation", func() {
			dcli.SetError()
			err := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network2",
			})

			Ω(err).Should(HaveOccurred())
			Ω(len(dcli.GetNetworks())).Should(Equal(1))
		})

		It("Should receive a database error", func() {
			db.SetError(1)
			err := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network3",
			})

			Ω(err).Should(HaveOccurred())
			Ω(len(dcli.GetNetworks())).Should(Equal(1))
		})
	})

	Describe("Remove Network", func() {
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		networkService, _ := network.NewService(dcli, db)
		It("Should remove a network", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network1",
			})

			err := networkService.RemoveNetworkByName(123, "network1")

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error when network cannot be removed", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network2",
			})

			dcli.SetError()

			err := networkService.RemoveNetworkByName(123, "network2")

			Ω(err).Should(HaveOccurred())
		})

		It("Should error when network is not found in db", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network3",
			})

			db.SetError(1)

			err := networkService.RemoveNetworkByName(123, "network3")

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
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network1",
			})

			err := networkService.AddContainerToNetwork(123, "network1", containerID)

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error when network does not exist in db", func() {
			err := networkService.AddContainerToNetwork(123, "idontexist", containerID)

			Ω(err).Should(HaveOccurred())
		})

		It("Should error when network does not exist on host", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network2",
			})

			dcli.SetError()

			err := networkService.AddContainerToNetwork(123, "network2", containerID)

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
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network1",
			})

			_ = networkService.AddContainerToNetwork(123, "network1", containerID)

			err := networkService.RemoveContainerFromNetwork(123, "network1", containerID)

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error when container cannot be removed from network", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network2",
			})

			_ = networkService.AddContainerToNetwork(123, "network2", containerID)

			dcli.SetError()

			err := networkService.RemoveContainerFromNetwork(123, "network2", containerID)

			Ω(err).Should(HaveOccurred())
		})

		It("Should error when network is not found in db", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network3",
			})

			_ = networkService.AddContainerToNetwork(123, "network3", containerID)

			db.SetError(1)

			err := networkService.RemoveContainerFromNetwork(123, "network3", containerID)

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

		It("Should return true when user has a network", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network1",
			})

			b := networkService.UserHasNetwork(123)

			Ω(b).Should(BeTrue())
		})

		It("Should return false when user does not have a network", func() {
			networkService.RemoveNetworkByName(123, "network1")

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
