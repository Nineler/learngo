package main

import (
	"connecter/infra"
	"fmt"
)

func getRetriever() retriever {
	return infra.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
