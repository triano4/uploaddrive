package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/triano4/uploaddrive/code"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)
	r.HandleFunc("/file", code.Client)

	fmt.Println("Serve localhost:8081")
	http.ListenAndServe(":8081", r)

}
