package main

import (
	// "fmt"
	// "net/http"

	// "github.com/gorilla/mux"
	"github.com/triano4/uploaddrive/code"
	// "fmt"
    // "io/ioutil"
    // "log"
    // "os"
    // "path/filepath"
)

func main() {
	// r := mux.NewRouter()

	// http.Handle("/", r)
	// r.HandleFunc("/file", code.Client)
	// code.Client()
 	
	code.UploadFile()
	// fmt.Println("Serve localhost:9000")
	// http.ListenAndServe(":9000", r)


}


