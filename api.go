package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateProfileEndpoint(w http.ResponseWriter, req *http.Request) {
//	params := mux.Vars(req)
	fmt.Printf("Received request")
	return
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/profiles", CreateProfileEndpoint).Methods("POST")
	fmt.Printf("Starting server")
	log.Fatal(http.ListenAndServe(":80", router))
}
