package user_test

import (
	"github.com/ttdennis/kontainer.io/testutils"
	"github.com/ttdennis/kontainer.io/user"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Describe("Create Service", func() {
		It("Should create service", func() {
			userService, err := user.NewService(testutils.NewMockDB())
			Ω(err).ShouldNot(HaveOccurred())
			Expect(userService).ToNot(BeZero())
		})

		It("Should return db error", func() {
			db := testutils.NewMockDB()
			db.SetError(1)
			_, err := user.NewService(db)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Create User", func() {
		db := testutils.NewMockDB()
		userService, _ := user.NewService(db)
		It("Should create user with new username", func() {
			id, err := userService.CreateUser("username", &user.Config{}, &user.Address{})
			Ω(err).ShouldNot(HaveOccurred())
			Expect(id).Should(BeEquivalentTo(1))
		})

		It("Should not create user with already used username", func() {
			_, err := userService.CreateUser("username", &user.Config{}, &user.Address{})
			Ω(err).Should(HaveOccurred())
		})

		It("Should return error on db failure", func() {
			db.SetError(2)
			_, err := userService.CreateUser("username2", &user.Config{}, &user.Address{})
			Ω(err).Should(HaveOccurred())
			db.SetError(3)
			_, err = userService.CreateUser("username2", &user.Config{}, &user.Address{})
			Ω(err).Should(HaveOccurred())
		})
	})
})
