package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "hello world,")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()
	//memanggil method yg mau di test
	HelloHandler(recorder, request)
	//test apakah response sesuai atau tidak
	response := recorder.Result()
	//baca responesnya dengan menggunakan io
	payload, _ := io.ReadAll(response.Body)
	bodyString := string(payload)

	// fmt.Println(bodyString)
	assert.Equal(t, "hello world,", bodyString, "result must be 'hello world,'")
	fmt.Println("tes done")
}
