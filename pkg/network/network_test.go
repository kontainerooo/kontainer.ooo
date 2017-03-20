package network_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ttdennis/kontainer.io/pkg/network"
	"github.com/ttdennis/kontainer.io/pkg/testutils"
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
	})
})
