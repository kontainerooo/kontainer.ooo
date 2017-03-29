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

	Describe("Edit Router Config", func() {
		db := testutils.NewMockDB()
		routingService, _ := routing.NewService(db)
		It("Should change Router Config", func() {
			refID, name := uint(1), "test"
			routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: refID,
				Name:  name,
			})

			newName := "test2"
			err := routingService.EditRouterConfig(refID, name, &routing.RouterConfig{
				Name: newName,
			})
			Ω(err).ShouldNot(HaveOccurred())

			conf := &routing.RouterConfig{}
			routingService.GetRouterConfig(refID, newName, conf)
			Expect(conf.Name).To(BeEquivalentTo(newName))
		})

		It("Should prevent from changing the refID", func() {
			err := routingService.EditRouterConfig(1, "test2", &routing.RouterConfig{
				RefID: 2,
			})
			Ω(err).Should(HaveOccurred())
		})

		It("Should return error on db failure", func() {
			db.SetError(1)
			err := routingService.EditRouterConfig(1, "", &routing.RouterConfig{})
			Ω(err).Should(HaveOccurred())

			db.SetError(2)
			err = routingService.EditRouterConfig(1, "", &routing.RouterConfig{})
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("GetRouterConfig", func() {
		db := testutils.NewMockDB()
		routingService, _ := routing.NewService(db)
		It("Should fill RouterConfig struct", func() {
			refID, name := uint(1), "test"
			routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: refID,
				Name:  name,
			})

			conf := &routing.RouterConfig{}
			err := routingService.GetRouterConfig(refID, name, conf)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(conf.Name).To(BeEquivalentTo(name))
		})

		It("Should return error if ID does not exist", func() {
			err := routingService.GetRouterConfig(28, "", &routing.RouterConfig{})
			Ω(err).Should(BeEquivalentTo(testutils.ErrNotFound))
		})

		It("Should return error on db failure", func() {
			db.SetError(1)
			err := routingService.GetRouterConfig(1, "", &routing.RouterConfig{})
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("DeleteRouterconfig", func() {
		db := testutils.NewMockDB()
		routingService, _ := routing.NewService(db)
		It("Should remove RouterConfig from DB", func() {
			refID, name := uint(1), "test"
			routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: refID,
				Name:  name,
			})
			err := routingService.RemoveRouterConfig(refID, name)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should return error on db failure", func() {
			refID, name := uint(1), "test"
			routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: refID,
				Name:  name,
			})
			db.SetError(1)
			err := routingService.RemoveRouterConfig(refID, name)
			Ω(err).Should(HaveOccurred())
		})

		It("Should return error if refID and/or name are not set", func() {
			err := routingService.RemoveRouterConfig(0, "")
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("AddLocationRule", func() {
		db := testutils.NewMockDB()
		routingService, _ := routing.NewService(db)
		It("Should add a LocationRule", func() {
			refID, name := uint(1), "test"
			routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: refID,
				Name:  name,
			})

			err := routingService.AddLocationRule(refID, name, &routing.LocationRule{
				Location: "/",
			})
			Ω(err).ShouldNot(HaveOccurred())

			conf := &routing.RouterConfig{}
			routingService.GetRouterConfig(refID, name, conf)
			Expect(conf.LocationRules).To(HaveLen(1))
		})

		It("Should return error if ID does not exist", func() {
			err := routingService.AddLocationRule(28, "", &routing.LocationRule{})
			Ω(err).Should(BeEquivalentTo(testutils.ErrNotFound))
		})

		It("Should return error on db failure", func() {
			db.SetError(1)
			err := routingService.AddLocationRule(1, "", &routing.LocationRule{})
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("RemoveLocationRule", func() {
		db := testutils.NewMockDB()
		routingService, _ := routing.NewService(db)
		It("Should add a LocationRule", func() {
			refID, name := uint(1), "test"
			routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: refID,
				Name:  name,
				LocationRules: routing.LocationRules{
					&routing.LocationRule{
						Location: "/",
					},
				},
			})

			err := routingService.RemoveLocationRule(refID, name, 0)
			Ω(err).ShouldNot(HaveOccurred())

			conf := &routing.RouterConfig{}
			routingService.GetRouterConfig(refID, name, conf)
			Expect(conf.LocationRules).To(HaveLen(0))
		})

		It("Should return error if ID does not exist", func() {
			err := routingService.RemoveLocationRule(28, "", 28)
			Ω(err).Should(BeEquivalentTo(testutils.ErrNotFound))
		})

		It("Should return error on db failure", func() {
			db.SetError(1)
			err := routingService.RemoveLocationRule(1, "", 0)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Change Listen Statement", func() {
		db := testutils.NewMockDB()
		routingService, _ := routing.NewService(db)
		It("Should change User Config", func() {
			refID, name, port := uint(1), "test", uint16(80)
			routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: refID,
				Name:  name,
			})

			err := routingService.ChangeListenStatement(refID, name, &routing.ListenStatement{
				Port: port,
			})
			Ω(err).ShouldNot(HaveOccurred())

			conf := &routing.RouterConfig{}
			routingService.GetRouterConfig(refID, name, conf)
			Expect(conf.ListenStatement.Port).To(BeEquivalentTo(port))
		})

		It("Should return error on db failure", func() {
			db.SetError(1)
			err := routingService.ChangeListenStatement(1, "", nil)
			Ω(err).Should(HaveOccurred())

			db.SetError(2)
			err = routingService.ChangeListenStatement(1, "", &routing.ListenStatement{})
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("AddServerName", func() {
		db := testutils.NewMockDB()
		routingService, _ := routing.NewService(db)
		It("Should add a ServerName", func() {
			refID, name := uint(1), "test"
			routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: refID,
				Name:  name,
			})

			err := routingService.AddServerName(refID, name, "test")
			Ω(err).ShouldNot(HaveOccurred())

			conf := &routing.RouterConfig{}
			routingService.GetRouterConfig(refID, name, conf)
			Expect(conf.ServerName).To(HaveLen(1))
		})

		It("Should return error if ID does not exist", func() {
			err := routingService.AddServerName(28, "", "")
			Ω(err).Should(BeEquivalentTo(testutils.ErrNotFound))
		})

		It("Should return error on db failure", func() {
			db.SetError(1)
			err := routingService.AddServerName(1, "", "")
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("RemoveServerName", func() {
		db := testutils.NewMockDB()
		routingService, _ := routing.NewService(db)
		It("Should add a LocationRule", func() {
			refID, name := uint(1), "test"
			routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID:      refID,
				Name:       name,
				ServerName: []string{"test"},
			})

			err := routingService.RemoveServerName(refID, name, 0)
			Ω(err).ShouldNot(HaveOccurred())

			conf := &routing.RouterConfig{}
			routingService.GetRouterConfig(refID, name, conf)
			Expect(conf.LocationRules).To(HaveLen(0))
		})

		It("Should return error if ID does not exist", func() {
			err := routingService.RemoveServerName(28, "", 28)
			Ω(err).Should(BeEquivalentTo(testutils.ErrNotFound))
		})

		It("Should return error on db failure", func() {
			db.SetError(1)
			err := routingService.RemoveServerName(1, "", 0)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Configurations", func() {
		It("Should return all configurations", func() {
			routingService, _ := routing.NewService(testutils.NewMockDB())
			routingService.CreateRouterConfig(&routing.RouterConfig{
				RefID: 1,
				Name:  "test",
			})

			conf := make([]routing.RouterConfig, 0)
			routingService.Configurations(&conf)
			Expect(conf).To(HaveLen(1))
		})
	})
})
