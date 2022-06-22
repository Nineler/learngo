package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type usererror string

func (e usererror) Error() string {
	return e.Message()
}
func (e usererror) Message() string {
	return string(e)
}

func Handler(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, "/list/") != 0 {
		return usererror("path start must is " + "with " + "/list/")
	}
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
