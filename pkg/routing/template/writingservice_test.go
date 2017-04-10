package template_test

import (
	"os"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing/template"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"
	"github.com/lib/pq"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Writingservice", func() {
	var (
		refID          = uint(1)
		name           = "name"
		completeConfig = &routing.RouterConfig{
			RefID: refID,
			Name:  name,
			ListenStatement: &routing.ListenStatement{
				IPAddress: abstraction.Inet("127.0.0.1"),
				Port:      1337,
				Keyword:   "ssl",
			},
			ServerName: pq.StringArray{"domain.com"},
		}
	)

	Describe("NewWritingService", func() {
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should return a new writing service", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, err := template.NewWritingService(s, template.Nginx, testPath)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(w).ShouldNot(BeNil())
		})

		It("Should return an error if the path is unaccessable", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			_, err := template.NewWritingService(s, template.Nginx, "bububububu")
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Create Router Config", func() {
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create RouterConfig with new RefID Name Pair", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			err := w.CreateRouterConfig(completeConfig)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should not create RouterConfig if RefID Name Pair already exists", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			err := w.CreateRouterConfig(completeConfig)
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if the config is incomplete", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			err := w.CreateRouterConfig(&routing.RouterConfig{})
			Ω(err).Should(HaveOccurred())
		})

		It("Should return error if the file could not be updated", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			os.RemoveAll(testPath)
			err := w.CreateRouterConfig(completeConfig)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Edit Router Config", func() {
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
			completeConfig.Name = name
		})

		It("Should change Router Config", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			newName := "test2"
			err := w.EditRouterConfig(refID, name, &routing.RouterConfig{
				Name: newName,
			})
			Ω(err).ShouldNot(HaveOccurred())

			conf := &routing.RouterConfig{}
			w.GetRouterConfig(refID, newName, conf)
			Expect(conf.Name).To(BeEquivalentTo(newName))
		})

		It("Should return an error if the underlying service returns one", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			err := w.EditRouterConfig(refID, name, &routing.RouterConfig{
				RefID: refID + 2,
			})
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if the config is falsy", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			err := w.EditRouterConfig(refID, name, &routing.RouterConfig{
				ListenStatement: &routing.ListenStatement{
					Port: 12,
				},
			})
			Ω(err).Should(HaveOccurred())
		})

		It("Should return error if the file could not be updated", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)
			os.RemoveAll(testPath)
			err := w.EditRouterConfig(refID, name, &routing.RouterConfig{
				Name: "test",
			})
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("GetRouterConfig", func() {
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should fill RouterConfig struct", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			conf := &routing.RouterConfig{}
			err := w.GetRouterConfig(refID, name, conf)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(conf.Name).To(BeEquivalentTo(name))
		})

		It("Should return error if ID does not exist", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			err := w.GetRouterConfig(28, "", &routing.RouterConfig{})
			Ω(err).Should(BeEquivalentTo(testutils.ErrNotFound))
		})
	})

	Describe("RemoveRouterConfig", func() {
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should remove RouterConfig from DB", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			err := w.RemoveRouterConfig(refID, name)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should return error if config does not exist", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			err := w.RemoveRouterConfig(0, "")
			Ω(err).Should(HaveOccurred())
		})

		It("Should return error if config does not exist", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			db.SetError(1)
			err := w.RemoveRouterConfig(refID, name)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Change Listen Statement", func() {
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should change Listen Statement", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			port := uint16(1025)
			err := w.ChangeListenStatement(refID, name, &routing.ListenStatement{
				IPAddress: abstraction.Inet("127.0.0.1"),
				Port:      port,
			})
			Ω(err).ShouldNot(HaveOccurred())

			conf := &routing.RouterConfig{}
			w.GetRouterConfig(refID, name, conf)
			Expect(conf.ListenStatement.Port).To(BeEquivalentTo(port))
		})

		It("Should return an error if the ListenStatement is falsy", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			err := w.ChangeListenStatement(refID, name, &routing.ListenStatement{
				Port: 12,
			})
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if the underlying service returns an error", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)
			db.SetError(1)
			err := w.ChangeListenStatement(refID, name, &routing.ListenStatement{
				IPAddress: abstraction.Inet("127.0.0.1"),
				Port:      1026,
			})
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if the file couldn't be updated", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			os.RemoveAll(testPath)
			err := w.ChangeListenStatement(refID, name, &routing.ListenStatement{
				IPAddress: abstraction.Inet("127.0.0.1"),
				Port:      1026,
			})
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("AddServerName", func() {
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should change the ServerName", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			err := w.AddServerName(refID, name, "domain2.com")
			Ω(err).ShouldNot(HaveOccurred())

			conf := &routing.RouterConfig{}
			w.GetRouterConfig(refID, name, conf)
			Expect(conf.ServerName).To(HaveLen(2))
		})

		It("Should change return an error if the server name is falsy", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			err := w.AddServerName(refID, name, "domain2")
			Ω(err).Should(HaveOccurred())
		})

		It("Should change return an error if the underlying service returns an error", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			db.SetError(1)
			err := w.AddServerName(refID, name, "domain2.com")
			Ω(err).Should(HaveOccurred())
		})

		It("Should change return an error if the underlying service returns an error", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			os.RemoveAll(testPath)
			err := w.AddServerName(refID, name, "domain2.com")
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("RemoveServerName", func() {
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should remove a ServerName", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			l := len(completeConfig.ServerName)

			err := w.RemoveServerName(refID, name, l-1)
			Ω(err).ShouldNot(HaveOccurred())

			conf := &routing.RouterConfig{}
			w.GetRouterConfig(refID, name, conf)
			Expect(conf.ServerName).To(HaveLen(l - 1))
		})

		It("Should return an error if the server name is falsy", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			err := w.AddServerName(refID, name, "domain2")
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if the underlying service returns an error", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			db.SetError(1)
			err := w.RemoveServerName(refID, name, 1)
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if the underlying service returns an error", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, _ := template.NewWritingService(s, template.Nginx, testPath)
			w.CreateRouterConfig(completeConfig)

			os.RemoveAll(testPath)
			err := w.RemoveServerName(refID, name, 0)
			Ω(err).Should(HaveOccurred())
		})
	})
})
