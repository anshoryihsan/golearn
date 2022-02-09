package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func formPost(res http.ResponseWriter, req *http.Request) {
	//parsing terelebih dahulu
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := req.PostForm.Get("first_name")
	lastName := req.PostForm.Get("last_name")
	//setelah itu ambil
	fmt.Fprintf(res, "helo %s %s", firstName, lastName)

}

func TestFormPost(t *testing.T) {
	requsetBody := strings.NewReader("first_name=ihsan&last_name=anshory")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requsetBody)
	request.Header.Add("Content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	formPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
