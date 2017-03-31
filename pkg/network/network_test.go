package network_test

import (
	"fmt"

	"github.com/go-kit/kit/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kontainerooo/kontainer.ooo/pkg/customercontainer"
	"github.com/kontainerooo/kontainer.ooo/pkg/network"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"
)

var _ = Describe("Network", func() {
	Describe("Create service", func() {
		mockFw := testutils.NewMockFirewallClient()
		mockFwEndpoints := testutils.NewMockFirewallEndpoints(log.NewNopLogger(), *mockFw)

		It("Should create service", func() {

			networkService, err := network.NewService(testutils.NewMockDCli(), testutils.NewMockDB(), mockFwEndpoints)
			Ω(networkService).ShouldNot(BeNil())
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should return db error", func() {
			db := testutils.NewMockDB()
			db.SetError(1)
			_, err := network.NewService(testutils.NewMockDCli(), db, mockFwEndpoints)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Create Network", func() {
		mockFw := testutils.NewMockFirewallClient()
		mockFwEndpoints := testutils.NewMockFirewallEndpoints(log.NewNopLogger(), *mockFw)
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		networkService, _ := network.NewService(dcli, db, mockFwEndpoints)

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

		It("Should receive a database error on retrieving", func() {
			db.SetError(1)
			err := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network3",
			})

			Ω(err).Should(HaveOccurred())
			Ω(len(dcli.GetNetworks())).Should(Equal(1))
		})

		It("Should fail when network already exists", func() {
			err := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network1",
			})

			Ω(err).Should(HaveOccurred())
			Ω(len(dcli.GetNetworks())).Should(Equal(1))
		})
	})

	Describe("Remove Network", func() {
		mockFw := testutils.NewMockFirewallClient()
		mockFwEndpoints := testutils.NewMockFirewallEndpoints(log.NewNopLogger(), *mockFw)
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		networkService, _ := network.NewService(dcli, db, mockFwEndpoints)
		It("Should remove a network", func() {
			err := networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network1",
			})

			fmt.Println(err)

			err = networkService.RemoveNetworkByName(123, "network1")

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

		It("Should error when network does not exist", func() {
			err := networkService.RemoveNetworkByName(123, "sdf")

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Add Container to network", func() {
		mockFw := testutils.NewMockFirewallClient()
		mockFwEndpoints := testutils.NewMockFirewallEndpoints(log.NewNopLogger(), *mockFw)
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		ccs := customercontainer.NewService(dcli)
		networkService, _ := network.NewService(dcli, db, mockFwEndpoints)
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

		It("Should fail on db error", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network3",
			})

			db.SetError(1)

			err := networkService.AddContainerToNetwork(123, "network3", containerID)

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Remove container from network", func() {
		mockFw := testutils.NewMockFirewallClient()
		mockFwEndpoints := testutils.NewMockFirewallEndpoints(log.NewNopLogger(), *mockFw)
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		ccs := customercontainer.NewService(dcli)
		networkService, _ := network.NewService(dcli, db, mockFwEndpoints)
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

		It("Should error when network does not exist", func() {
			err := networkService.RemoveContainerFromNetwork(123, "network4", containerID)

			Ω(err).Should(HaveOccurred())
		})
	})
})
