package belajargolangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func serveFile(res http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("name") != "" {
		http.ServeFile(res, req, "./resources/ok.html")
	} else {
		http.ServeFile(res, req, "./resources/notfound.html")
	}
}

func TestServerFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(serveFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/notfound.html
var notfound string

func serveFileEmbed(res http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("name") != "" {
		fmt.Fprint(res, resourcesOk)
	} else {
		fmt.Fprint(res, notfound)
	}
}
func TestServerFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(serveFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
