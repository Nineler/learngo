package infra

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Retriever struct {
}

func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
