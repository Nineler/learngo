package main

import (
	"GOProject/errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appfileHandle func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appfileHandle) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic:%v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf(
				"Error handing request:%s",
				err.Error(),
			)
			if userErr, ok := err.(usererror); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code),
				code)
		}
	}
}

type usererror interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.Handler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
