package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingusererror string

func (e testingusererror) Error() string {
	return e.Message()
}

func (e testingusererror) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingusererror("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknow(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("Unkown")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println(writer, "no error")
	return nil
}

var tests = []struct {
	h       appfileHandle
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknow, 500, "Internal Server Error"},
	{noError, 200, ""},
}

func verifyResponse(resp *http.Response, expectCode int, expectMessage string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectCode || body != expectMessage {
		t.Errorf("expect(%d, %s)"+"got(%d, %s)",
			expectCode, expectMessage, resp.StatusCode, body,
		)
	}
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		asdas := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet, "http://imooc.com", nil)
		asdas(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}
