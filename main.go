package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		tee := io.MultiWriter(res, log.Writer())
		fmt.Fprintln(tee, "Hello Acorn!")
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
