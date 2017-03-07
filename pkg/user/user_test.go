package user_test

import (
	"context"

	"github.com/ttdennis/kontainer.io/pkg/testutils"
	"github.com/ttdennis/kontainer.io/pkg/user"

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
		userService = user.NewTransactionBasedService(userService)
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

		It("Should not add address if user can't be created", func() {
			db.SetError(3)
			city := "city"
			_, err := userService.CreateUser("username3", &user.Config{}, &user.Address{City: city})
			Ω(err).Should(HaveOccurred())
			db.Where("city = ?", city)
			Ω(db.GetValue()).Should(BeNil())
		})
	})

	Describe("Edit User", func() {
		db := testutils.NewMockDB()
		userService, _ := user.NewService(db)
		userService = user.NewTransactionBasedService(userService)
		It("Should change User Config", func() {
			id, _ := userService.CreateUser("foo", &user.Config{Email: "test@abc.com"}, &user.Address{})
			email := "new@abc.com"
			err := userService.EditUser(id, &user.Config{Email: email})
			Ω(err).ShouldNot(HaveOccurred())
			user := &user.User{}
			userService.GetUser(id, user)
			Expect(user.Email).To(BeEquivalentTo(email))
		})

		It("Should return error on db failure", func() {
			id, _ := userService.CreateUser("foo", &user.Config{}, &user.Address{})
			db.SetError(1)
			err := userService.EditUser(id, &user.Config{Email: "email"})
			Ω(err).Should(HaveOccurred())

			db.SetError(2)
			err = userService.EditUser(id, &user.Config{Email: "email"})
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Change Username", func() {
		db := testutils.NewMockDB()
		userService, _ := user.NewService(db)
		userService = user.NewTransactionBasedService(userService)
		It("Should rename User", func() {
			id, _ := userService.CreateUser("foo", &user.Config{}, &user.Address{})
			username := "bar"
			err := userService.ChangeUsername(id, username)
			Ω(err).ShouldNot(HaveOccurred())
			user := &user.User{}
			userService.GetUser(id, user)
			Expect(user.Username).To(BeEquivalentTo(username))
		})

		It("Should return error on db failure", func() {
			id, _ := userService.CreateUser("foo", &user.Config{}, &user.Address{})
			db.SetError(1)
			err := userService.ChangeUsername(id, "bar")
			Ω(err).Should(HaveOccurred())

			db.SetError(2)
			err = userService.ChangeUsername(id, "bar")
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Delete User", func() {
		db := testutils.NewMockDB()
		userService, _ := user.NewService(db)
		userService = user.NewTransactionBasedService(userService)
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

	Describe("GetUser", func() {
		db := testutils.NewMockDB()
		userService, _ := user.NewService(db)
		userService = user.NewTransactionBasedService(userService)
		It("Should fill user struct", func() {
			username := "username"
			id, _ := userService.CreateUser(username, &user.Config{}, &user.Address{})
			user := &user.User{}
			err := userService.GetUser(id, user)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(user.Username).To(BeEquivalentTo(username))
		})

		It("Should return error if ID does not exist", func() {
			user := &user.User{}
			err := userService.GetUser(28, user)
			Ω(err).Should(BeEquivalentTo(testutils.ErrNotFound))
		})

		It("Should return error on db failure", func() {
			user := &user.User{}
			db.SetError(1)
			err := userService.GetUser(1, user)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Endpoints and Transport", func() {
		db := testutils.NewMockDB()
		s, _ := user.NewService(db)
		s = user.NewTransactionBasedService(s)
		es := &user.Endpoints{}
		ctx := context.Background()
		It("Should create valid Endpoints", func() {
			es.CreateUserEndpoint = user.MakeCreateUserEndpoint(s)
			es.EditUserEndpoint = user.MakeEditUserEndpoint(s)
			es.ChangeUsernameEndpoint = user.MakeChangeUsernameEndpoint(s)
			es.DeleteUserEndpoint = user.MakeDeleteUserEndpoint(s)
			es.ResetPasswordEndpoint = user.MakeResetPasswordEndpoint(s)
			es.GetUserEndpoint = user.MakeGetUserEndpoint(s)
		})

		Context("CreateUserEndpoint", func() {
			It("Should work with CreateUser request and response", func() {
				res, err := es.CreateUserEndpoint(ctx, user.CreateUserRequest{
					Username: "foo",
					Cfg:      &user.Config{},
					Adr:      &user.Address{},
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(res.(user.CreateUserResponse).ID).To(BeEquivalentTo(1))
			})
		})

		Context("ChangeUsernameEndpoint", func() {
			It("Should work with ChangeUsername request and response", func() {
				res, err := es.ChangeUsernameEndpoint(ctx, user.ChangeUsernameRequest{
					ID:       1,
					Username: "bar",
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(res.(user.ChangeUsernameResponse).Error).NotTo(HaveOccurred())
			})
		})

		Context("EditUserEndpoint", func() {
			It("Should work with EditUser request and response", func() {
				res, err := es.EditUserEndpoint(ctx, user.EditUserRequest{
					ID: 1,
					Cfg: &user.Config{
						Email: "test",
					},
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(res.(user.EditUserResponse).Error).NotTo(HaveOccurred())
			})
		})

		Context("GetUserEndpoint", func() {
			It("Should work with GetUser request and response", func() {
				res, err := es.GetUserEndpoint(ctx, user.GetUserRequest{
					ID: 1,
				})
				r := res.(user.GetUserResponse)
				Expect(err).NotTo(HaveOccurred())
				Expect(r.Error).NotTo(HaveOccurred())
			})
		})

		Context("DeleteUserEndpoint", func() {
			It("Should work with DeleteUser request and response", func() {
				res, err := es.DeleteUserEndpoint(ctx, user.DeleteUserRequest{
					ID: 1,
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(res.(user.DeleteUserResponse).Error).NotTo(HaveOccurred())
			})
		})

	})

})
