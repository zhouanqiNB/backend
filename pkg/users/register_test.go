package users_test

import (
	"bytes"
	"encoding/json"

	// . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zhouanqiNB/backend/pkg/users"
)

type FakeUserRepository struct {
	LoginResult *users.User
}

func (fur FakeUserRepository) RegisterUser(user *users.User) error {
	return nil
}

func (fur FakeUserRepository) FindByEmailAndPassword(email, password string) (*users.User, error) {
	return fur.LoginResult, nil
}

// var _ = Describe("Users", func() {

// 	var userRequest = users.UserRegistration{
// 		User: users.User{
// 			Username: "user",
// 			Email:    "user@example.com",
// 			Password: "s3cr3t",
// 		}}

// 	var expectedUserResponse = users.UserRegistration{
// 		User: users.User{
// 			Username: "user",
// 			Email:    "user@example.com",
// 		}}

// 	It("should register", func() {
// 		handler := users.UserRegistrationHandler{
// 			Path:           "/users",
// 			UserRepository: &FakeUserRepository{},
// 		}
// 		testResponseWriter := httptest.NewRecorder()

// 		handler.Register(
// 			testResponseWriter,
// 			httptest.NewRequest("POST", "/users", strings.NewReader(marshalRegistration(userRequest))))

// 		Expect(testResponseWriter.Code).To(Equal(201))
// 		Expect(unmarshalRegistration(testResponseWriter.Body)).To(Equal(&expectedUserResponse))
// 	})

// })

func marshalRegistration(registration users.UserRegistrationRequest) string {
	payload, err := json.Marshal(registration)
	Expect(err).To(BeNil(), "JSON marshalling should work")
	return string(payload)
}

func unmarshalRegistration(payload *bytes.Buffer) *users.UserRegistrationRequest {
	var result users.UserRegistrationRequest
	err := json.Unmarshal(payload.Bytes(), &result)
	Expect(err).To(BeNil(), "JSON unmarshalling should work")
	return &result
}
