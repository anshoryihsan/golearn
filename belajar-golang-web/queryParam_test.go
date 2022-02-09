package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//query parameter
func SayHello(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(res, "hello")
	} else {
		fmt.Fprintf(res, "hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=ihsan", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	bodyString, _ := io.ReadAll(response.Body)

	fmt.Println(bodyString)
}

//multiple query param
func MultipleQuery(res http.ResponseWriter, req *http.Request) {
	first_name := req.URL.Query().Get("first_name")
	last_name := req.URL.Query().Get("last_name")

	fmt.Fprintf(res, "hello %s %s", first_name, last_name)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=ihsan&last_name=anshory", nil)
	recorder := httptest.NewRecorder()

	MultipleQuery(recorder, request)

	response := recorder.Result()
	bodyString, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyString))
}
