package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserLogin struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type LoggedInUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserLoginHandler struct {
	Path           string
	UserRepository UserRepository
}

func (u *UserLoginHandler) Login(writer http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	userLoginRequest := User{}
	_ = json.Unmarshal(requestBody, &userLoginRequest)

	requestUser := userLoginRequest
	println("requestUser.Username")
	println(requestUser.Username)
	println("requestUser.Email")
	println(requestUser.Email)
	println("requestUser.Password")
	println(requestUser.Password)
	user, _ := u.UserRepository.FindByEmailAndPassword(
		requestUser.Email,
		requestUser.Password)

	if user == nil {
		writer.WriteHeader(401)
		return
	}

	token, _ := CreateToken(user)
	loggedInUser := &LoggedInUser{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
	writer.WriteHeader(200)
	bytes, _ := json.Marshal(&loggedInUser)
	_, _ = writer.Write(bytes)
}

func CreateToken(user *User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_ACCESS")))
}
