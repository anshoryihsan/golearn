package cookie

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(res http.ResponseWriter, req *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = req.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(res, cookie)
	fmt.Fprint(res, "success create cookie")
}

func GetCookie(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(res, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(res, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=ihsan", nil)
	recoder := httptest.NewRecorder()

	SetCookie(recoder, request)

	cookies := recoder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookies %s:%s", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/=ihsan", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "ihsan"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
