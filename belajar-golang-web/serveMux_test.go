package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

//serveMux digunakan untuk router
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "hello world")
	})
	mux.HandleFunc("/hi", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "hi")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
