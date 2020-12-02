package main

import (
	"github.com/triano4/uploaddrive/api"
)

func main() {
	// r := mux.NewRouter()

	// http.Handle("/", r)
	// r.HandleFunc("/file", code.Client)

	// fmt.Println("Serve localhost:8081")
	// http.ListenAndServe(":8081", r)

	api.Run()

}
