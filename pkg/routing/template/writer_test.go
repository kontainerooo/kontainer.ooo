package template_test

import (
	"io/ioutil"
	"os"

	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing/template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var testPath = "/tmp/test-template-kroo/"

var _ = Describe("Writer", func() {
	Describe("New Writer", func() {
		It("Should return new writer", func() {
			w, err := template.NewWriter(template.Nginx, "/tmp")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(w).ShouldNot(BeNil())
		})

		It("Should return an error if router does not exist", func() {
			_, err := template.NewWriter(1000, "/tmp")
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if path is no directory", func() {
			_, err := template.NewWriter(template.Nginx, "/etc/hosts")
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if path does not exist", func() {
			_, err := template.NewWriter(template.Nginx, "-")
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Create File", func() {
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should write Conf file", func() {
			w, _ := template.NewWriter(template.Nginx, testPath)

			refID, name := uint(1), "test"
			c := &routing.RouterConfig{
				RefID: refID,
				Name:  name,
			}

			err := w.CreateFile(c)
			Ω(err).ShouldNot(HaveOccurred())

			b, err := ioutil.ReadFile(w.CreatePath(refID, name))
			Ω(err).ShouldNot(HaveOccurred())
			Ω(b).ShouldNot(BeEmpty())
		})
	})

	Describe("RemoveFile", func() {
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should remove a file", func() {
			w, _ := template.NewWriter(template.Nginx, testPath)

			refID, name := uint(1), "test"
			c := &routing.RouterConfig{
				RefID: refID,
				Name:  name,
			}
			w.CreateFile(c)

			err := w.RemoveFile(refID, name)
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
