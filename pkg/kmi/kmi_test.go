package kmi_test

import (
	"github.com/ttdennis/kontainer.io/pkg/kmi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KMI", func() {
	Describe("Create Service", func() {
		It("Should create service", func() {
			kmiService := kmi.NewService()
			Expect(kmiService).ToNot(BeZero())
		})
	})

	XDescribe("Add KMI", func() {

	})

	XDescribe("Remove KMI", func() {

	})

	XDescribe("Get KMI", func() {

	})

	XDescribe("KMI", func() {

	})
})
