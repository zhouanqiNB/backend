package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//func(writer http.ResponseWriter, request *http.Request) {
//
//}

type UserRegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserRegistrationHandler struct {
	Path           string
	UserRepository UserRepository
}

// 这里requestBody是默认post的
func (u *UserRegistrationHandler) Register(writer http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	requestUser := User{}
	_ = json.Unmarshal(requestBody, &requestUser)
	println("requestUser.Username")
	println(requestUser.Username)
	println("requestUser.Email")
	println(requestUser.Email)
	println("requestUser.Password")
	println(requestUser.Password)

	_ = u.UserRepository.RegisterUser(&requestUser)

	writer.WriteHeader(201)
	writer.Header().Add("Content-Type", "application/json")
	userRegistrationResponse := UserRegistrationRequest{
		Username: requestUser.Username,
		Email:    requestUser.Email,
	}
	bytes, _ := json.Marshal(&userRegistrationResponse)
	_, _ = writer.Write(bytes)
}
