package stackoverflow

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type QueryRequest struct {
	QueryStr string `json:"query_str"`
}

type QueryHandler struct {
	Path                    string
	StackOverflowRepository StackOverflowRepository
}

func (u *QueryHandler) Query(writer http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	queryRequest := QueryRequest{}
	_ = json.Unmarshal(requestBody, &queryRequest)
	queryStr := queryRequest.QueryStr

	result, _ := u.StackOverflowRepository.DoQuery(queryStr)

	if result == nil {
		writer.WriteHeader(401)
		return
	}

	writer.WriteHeader(200)
	bytes, _ := json.Marshal(&result)
	_, _ = writer.Write(bytes)
}
