package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	indexHTML := MustAssetString("templates/index.html")
	io.WriteString(w, indexHTML)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3333", nil)
}
