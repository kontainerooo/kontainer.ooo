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
})
