package stackoverflow

import (
	"encoding/json"
	"net/http"
)

type QueryAllHandler struct {
	Path                    string
	StackOverflowRepository StackOverflowRepository
}

func (u *QueryAllHandler) QueryAll(writer http.ResponseWriter, request *http.Request) {
	result, _ := u.StackOverflowRepository.DoQueryAll()

	if result == nil {
		writer.WriteHeader(401)
		return
	}

	writer.WriteHeader(200)
	bytes, _ := json.Marshal(&result)
	_, _ = writer.Write(bytes)
}
