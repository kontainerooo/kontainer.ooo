package kmi_test

import (
	"github.com/ttdennis/kontainer.io/pkg/kmi"
	"github.com/ttdennis/kontainer.io/pkg/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KMI", func() {
	Describe("Create Service", func() {
		It("Should create service", func() {
			kmiService, err := kmi.NewService(testutils.NewMockDB())
			Ω(err).ShouldNot(HaveOccurred())
			Expect(kmiService).ToNot(BeZero())
		})

		It("Should return db error", func() {
			db := testutils.NewMockDB()
			db.SetError(1)
			_, err := kmi.NewService(db)
			Ω(err).Should(HaveOccurred())
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
