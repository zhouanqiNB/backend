package users_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zhouanqiNB/backend/pkg/users"
)

var _ = Describe("Users", func() {
	userLoginRequest := users.UserLogin{
		// User: users.User{
		// 	Email:    "florent@example.org",
		// 	Password: "very-secure",
		// },
	}

	It("should log in", func() {
		email := "florent@example.org"
		handler := users.UserLoginHandler{
			Path: "/users/login",
			UserRepository: &FakeUserRepository{
				LoginResult: &users.User{
					Username: "flo",
					Email:    email,
				},
			},
		}
		testResponseWriter := httptest.NewRecorder()

		handler.Login(
			testResponseWriter,
			httptest.NewRequest(
				"POST",
				"/users/login",
				strings.NewReader(marshalLogin(&userLoginRequest))))

		Expect(testResponseWriter.Code).To(Equal(200))
		login := unmarshalLogin(testResponseWriter.Body)
		Expect(login.Email).To(Equal(email))
		Expect(login.Username).To(Equal("flo"))
		Expect(login.Token).NotTo(Equal(""), "token should be set")
	})

	It("should returns unauthorized response if login fails", func() {
		handler := users.UserLoginHandler{
			Path: "/users/login",
			UserRepository: &FakeUserRepository{
				LoginResult: nil,
			},
		}
		testResponseWriter := httptest.NewRecorder()

		handler.Login(
			testResponseWriter,
			httptest.NewRequest(
				"POST",
				"/users/login",
				strings.NewReader(marshalLogin(&userLoginRequest))))

		Expect(testResponseWriter.Code).To(Equal(401))
	})
})

func unmarshalLogin(payload *bytes.Buffer) *users.LoggedInUser {
	var result users.LoggedInUser
	err := json.Unmarshal(payload.Bytes(), &result)
	Expect(err).To(BeNil(), "JSON unmarshalling should work")
	return &result
}

func marshalLogin(login *users.UserLogin) string {
	payload, err := json.Marshal(login)
	Expect(err).To(BeNil(), "JSON marshalling should work")
	return string(payload)
}
