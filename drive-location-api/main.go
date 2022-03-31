package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

)

type Location struct {
	Type string  `json:"type"`
	coordinates [2]float64	
}


func createLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	return
}


func searchLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	return
}



func main(){
	fmt.Println("Hello World")

	r := mux.NewRouter()
    r.HandleFunc("/locations", createLocation).Methods("POST")
    r.HandleFunc("/locations", searchLocation).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:9000", nil)
}