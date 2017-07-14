package main

import (
	"net/http"

	spafs "github.com/koron/go-spafs"
)

func main() {
	err := http.ListenAndServe(":8080", spafs.FileServer(http.Dir("testdata")))
	if err != nil {
		panic(err)
	}
}
