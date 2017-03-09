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

	Describe("Add KMI", func() {
		db := testutils.NewMockDB()
		kmiService, _ := kmi.NewService(db)
		It("Should Add KMI", func() {
			id, err := kmiService.AddKMI("test.kmi")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(id).Should(BeEquivalentTo(1))
		})

		XIt("error handling test")
	})

	Describe("Remove KMI", func() {
		db := testutils.NewMockDB()
		kmiService, _ := kmi.NewService(db)
		It("Should Remove KMI", func() {
			id, _ := kmiService.AddKMI("test.kmi")
			err := kmiService.RemoveKMI(id)
			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	XDescribe("Get KMI", func() {

	})

	XDescribe("KMI", func() {

	})
})
