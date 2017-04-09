package template_test

import (
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing/template"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Writingservice", func() {
	Describe("NewWritingService", func() {
		It("Should return a new writing service", func() {
			db := testutils.NewMockDB()
			s, _ := routing.NewService(db)
			w, err := template.NewWritingService(s, template.Nginx, "/tmp")
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
})
