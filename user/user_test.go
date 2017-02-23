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
		})

		PIt("Should not add address if user can't be created")
	})

	Describe("Delete User", func() {
		db := testutils.NewMockDB()
		userService, _ := user.NewService(db)
		It("Should remove User from DB", func() {
			id, _ := userService.CreateUser("username", &user.Config{}, &user.Address{})
			err := userService.DeleteUser(id)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should return error on db failure", func() {
			id, _ := userService.CreateUser("username", &user.Config{}, &user.Address{})
			db.SetError(1)
			err := userService.DeleteUser(id)
			Ω(err).Should(HaveOccurred())
		})

		It("Should return error if ID does not exist", func() {
			err := userService.DeleteUser(24)
			Ω(err).Should(BeEquivalentTo(testutils.ErrNotFound))
		})
	})

})
