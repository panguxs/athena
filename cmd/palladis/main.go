package main

import (
	"log"
	"net/http"

	"github.com/panguxs/palladis/api"
	"github.com/panguxs/palladis/staticfile"
)

func main() {

	fnhs := make(staticfile.FileNameSpaceHandlerMap)
	fnh := &staticfile.FileNameSpaceHandler{
		Name:    "test",
		Handler: http.FileServer(http.Dir(".")),
	}
	fnhs["localhost:8080"] = fnh

	go api.Register()
	log.Fatal(http.ListenAndServe(":8080", fnhs))
}
