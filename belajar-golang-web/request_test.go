package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, req.Method)
		fmt.Fprintln(res, req.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
