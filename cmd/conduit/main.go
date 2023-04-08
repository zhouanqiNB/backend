package main

import (
	"net/http"
	"os"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/zhouanqiNB/backend/pkg/stackoverflow"
	"github.com/zhouanqiNB/backend/pkg/users"
)

func main() {
	neo4jUri, found := os.LookupEnv("NEO4J_URI")
	if !found {
		panic("NEO4J_URI not set")
	}
	neo4jUsername, found := os.LookupEnv("NEO4J_USERNAME")
	if !found {
		panic("NEO4J_USERNAME not set")
	}
	neo4jPassword, found := os.LookupEnv("NEO4J_PASSWORD")
	if !found {
		panic("NEO4J_PASSWORD not set")
	}

	usersRepository := users.UserNeo4jRepository{
		Driver: driver(neo4jUri, neo4j.BasicAuth(neo4jUsername, neo4jPassword, "")),
	}
	// println(neo4jUri)
	// println(neo4jUsername)
	// println(neo4jPassword)

	// this one is working but somehow I just can't transfer my args(by post,json)
	registrationHandler := &users.UserRegistrationHandler{
		Path:           "/users",
		UserRepository: &usersRepository,
	}
	loginHandler := &users.UserLoginHandler{
		Path:           "/users/login",
		UserRepository: &usersRepository,
	}

	server := http.NewServeMux()
	server.HandleFunc(registrationHandler.Path, registrationHandler.Register)
	server.HandleFunc(loginHandler.Path, loginHandler.Login)

	// ======

	// use driver to build a connection with db.
	stackOverflowRepository := stackoverflow.StackOverflowNeo4jRepository{
		Driver: driver(neo4jUri, neo4j.BasicAuth(neo4jUsername, neo4jPassword, "")),
	}
	// 画面的最初展现出全景
	queryAllHandler := &stackoverflow.QueryAllHandler{
		Path:                    "/query_all",
		StackOverflowRepository: &stackOverflowRepository,
	}
	queryHandler := &stackoverflow.QueryHandler{
		Path:                    "/query",
		StackOverflowRepository: &stackOverflowRepository,
	}

	server.HandleFunc(queryAllHandler.Path, queryAllHandler.QueryAll)
	server.HandleFunc(queryHandler.Path, queryHandler.Query)

	if err := http.ListenAndServe(":3000", server); err != nil {
		panic(err)
	}
}

func driver(target string, token neo4j.AuthToken) neo4j.Driver {
	result, err := neo4j.NewDriver(target, token)
	if err != nil {
		panic(err)
	}
	return result
}
