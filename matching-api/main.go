package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func main(){
	fmt.Println("Hello World") 
	r := mux.NewRouter()
    r.HandleFunc("/locations", searchLocation).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:9002", nil)
}