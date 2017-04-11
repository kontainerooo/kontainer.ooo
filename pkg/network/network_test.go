package network_test

import (
	"github.com/go-kit/kit/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kontainerooo/kontainer.ooo/pkg/customercontainer"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
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
		ccs, _ := customercontainer.NewService(dcli, db)
		mockKMI := testutils.NewMockKMIClient()
		mockKMIEndpoints := testutils.NewMockKMIEndpoints(log.NewNopLogger(), *mockKMI)
		mockKMI.AddMockKMI(0, kmi.KMI{
			KMDI: kmi.KMDI{
				ID:          1,
				Name:        "node",
				Version:     "",
				Description: "",
				Type:        3,
			},
			Dockerfile:  "FROM FROM node:7-wheezy",
			Context:     "./container-test",
			Commands:    nil,
			Environment: nil,
			Frontend:    nil,
			Imports:     nil,
			Interfaces:  nil,
			Mounts:      nil,
			Variables:   nil,
			Resources: map[string]interface{}{
				"cpus": 1,
				"mem":  500,
				"swap": 500,
			},
		})

		ccs.AddLogger(log.NewNopLogger())
		ccs.AddKMIClient(mockKMIEndpoints)
		_, containerID, _ := ccs.CreateContainer(123, 0)
		It("Should remove a network", func() {
			networkService.CreatePrimaryNetworkForContainer(123, &network.Config{
				Driver: "bridge",
				Name:   "network1",
			}, containerID)
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
		ccs, _ := customercontainer.NewService(dcli, db)
		networkService, _ := network.NewService(dcli, db, mockFwEndpoints)
		mockKMI := testutils.NewMockKMIClient()
		mockKMIEndpoints := testutils.NewMockKMIEndpoints(log.NewNopLogger(), *mockKMI)
		mockKMI.AddMockKMI(0, kmi.KMI{
			KMDI: kmi.KMDI{
				ID:          1,
				Name:        "node",
				Version:     "",
				Description: "",
				Type:        3,
			},
			Dockerfile:  "FROM FROM node:7-wheezy",
			Context:     "./container-test",
			Commands:    nil,
			Environment: nil,
			Frontend:    nil,
			Imports:     nil,
			Interfaces:  nil,
			Mounts:      nil,
			Variables:   nil,
			Resources: map[string]interface{}{
				"cpus": 1,
				"mem":  500,
				"swap": 500,
			},
		})

		ccs.AddLogger(log.NewNopLogger())
		ccs.AddKMIClient(mockKMIEndpoints)
		_, containerID, _ := ccs.CreateContainer(123, 0)

		It("Should add a container to a network", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network1",
			})

			err := networkService.AddContainerToNetwork(123, "network1", containerID)

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error when primary network already has a container", func() {
			_, containerTwo, _ := ccs.CreateContainer(123, 0)

			networkService.CreatePrimaryNetworkForContainer(123, &network.Config{
				Driver: "bridge",
				Name:   "network2",
			}, containerTwo)

			err := networkService.AddContainerToNetwork(123, "network2", containerID)

			Ω(err).Should(HaveOccurred())
		})

		It("Should error when network does not exist in db", func() {
			err := networkService.AddContainerToNetwork(123, "idontexist", containerID)

			Ω(err).Should(HaveOccurred())
		})

		It("Should error when network does not exist on host", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network3",
			})

			dcli.SetError()

			err := networkService.AddContainerToNetwork(123, "network3", containerID)

			Ω(err).Should(HaveOccurred())
		})

		It("Should fail on db error", func() {
			_ = networkService.CreateNetwork(123, &network.Config{
				Driver: "bridge",
				Name:   "network4",
			})

			db.SetError(1)

			err := networkService.AddContainerToNetwork(123, "network4", containerID)

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Remove container from network", func() {
		mockFw := testutils.NewMockFirewallClient()
		mockFwEndpoints := testutils.NewMockFirewallEndpoints(log.NewNopLogger(), *mockFw)
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		ccs, _ := customercontainer.NewService(dcli, db)
		networkService, _ := network.NewService(dcli, db, mockFwEndpoints)
		mockKMI := testutils.NewMockKMIClient()
		mockKMIEndpoints := testutils.NewMockKMIEndpoints(log.NewNopLogger(), *mockKMI)
		mockKMI.AddMockKMI(0, kmi.KMI{
			KMDI: kmi.KMDI{
				ID:          1,
				Name:        "node",
				Version:     "",
				Description: "",
				Type:        3,
			},
			Dockerfile:  "FROM FROM node:7-wheezy",
			Context:     "./container-test",
			Commands:    nil,
			Environment: nil,
			Frontend:    nil,
			Imports:     nil,
			Interfaces:  nil,
			Mounts:      nil,
			Variables:   nil,
			Resources: map[string]interface{}{
				"cpus": 1,
				"mem":  500,
				"swap": 500,
			},
		})

		ccs.AddLogger(log.NewNopLogger())
		ccs.AddKMIClient(mockKMIEndpoints)
		_, containerID, _ := ccs.CreateContainer(123, 0)

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

	Describe("Port exposure", func() {
		mockFw := testutils.NewMockFirewallClient()
		mockFwEndpoints := testutils.NewMockFirewallEndpoints(log.NewNopLogger(), *mockFw)
		dcli := testutils.NewMockDCli()
		db := testutils.NewMockDB()
		networkService, _ := network.NewService(dcli, db, mockFwEndpoints)
		ccs, _ := customercontainer.NewService(dcli, db)
		mockKMI := testutils.NewMockKMIClient()
		mockKMIEndpoints := testutils.NewMockKMIEndpoints(log.NewNopLogger(), *mockKMI)
		mockKMI.AddMockKMI(0, kmi.KMI{
			KMDI: kmi.KMDI{
				ID:          1,
				Name:        "node",
				Version:     "",
				Description: "",
				Type:        3,
			},
			Dockerfile:  "FROM FROM node:7-wheezy",
			Context:     "./container-test",
			Commands:    nil,
			Environment: nil,
			Frontend:    nil,
			Imports:     nil,
			Interfaces:  nil,
			Mounts:      nil,
			Variables:   nil,
			Resources: map[string]interface{}{
				"cpus": 1,
				"mem":  500,
				"swap": 500,
			},
		})

		ccs.AddLogger(log.NewNopLogger())
		ccs.AddKMIClient(mockKMIEndpoints)

		_, containerIDOne, _ := ccs.CreateContainer(123, 0)
		_, containerIDTwo, _ := ccs.CreateContainer(123, 0)

		It("Should expose port from container to container", func() {
			networkService.CreatePrimaryNetworkForContainer(123, &network.Config{
				Driver: "bridge",
				Name:   "network1",
			}, containerIDOne)

			networkService.CreatePrimaryNetworkForContainer(123, &network.Config{
				Driver: "bridge",
				Name:   "network2",
			}, containerIDTwo)
			err := networkService.ExposePortToContainer(123, containerIDOne, 8080, "tcp", containerIDTwo)

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should remove a port exposure", func() {
			err := networkService.RemovePortFromContainer(123, containerIDOne, 8080, "tcp", containerIDTwo)
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
