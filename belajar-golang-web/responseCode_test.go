package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if name == "" {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "name is empty")
	} else {
		fmt.Fprintf(res, "hello %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=ihsan", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	fmt.Println(string(body))
}
