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
	var completeConfig = &routing.RouterConfig{
		RefID: 1,
		Name:  "name",
		ListenStatement: &routing.ListenStatement{
			IPAddress: abstraction.Inet("127.0.0.1"),
			Port:      1337,
			Keyword:   "ssl",
		},
		ServerName: pq.StringArray{"domain.com"},
	}

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
})
