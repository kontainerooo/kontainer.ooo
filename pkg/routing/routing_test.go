package routing_test

import (
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Routing", func() {
	Describe("Create Service", func() {
		It("Should create service", func() {
			routingService, err := routing.NewService(testutils.NewMockDB())
			Ω(err).ShouldNot(HaveOccurred())
			Expect(routingService).ToNot(BeZero())
		})

		It("Should return db error", func() {
			db := testutils.NewMockDB()
			db.SetError(1)
			_, err := routing.NewService(db)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Create Router Config", func() {
		db := testutils.NewMockDB()
		routingService, _ := routing.NewService(db)
		It("Should create RouterConfig with new RefID Name Pair", func() {
			err := routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: 0,
				Name:  "test",
			})
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should not create RouterConfig if RefID Name Pair already exists", func() {
			err := routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: 0,
				Name:  "test",
			})
			Ω(err).Should(HaveOccurred())
		})

		It("Should return error on db failure", func() {
			db.SetError(2)
			err := routingService.CreateRouterConfig(&routing.RouterConfig{})
			Ω(err).Should(HaveOccurred())
		})
	})
})
