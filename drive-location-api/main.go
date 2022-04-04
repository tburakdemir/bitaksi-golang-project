package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GeoJSON object. Type is always "Point"
type Location struct {
	Type        string    `json:"-"`
	Coordinates [2]float64 `json:"coordinates"`
}

type LocationDTO struct{
	Type        string    `json:"type",omitempty`
	Coordinates [2]float64 `json:"coordinates"`
}

type LocationModal struct {
	Type string
	Coordinates [2]float64
}



func NewLocation(dto LocationDTO) (*Location, error) {
	if dto.Coordinates[0] < -90 || dto.Coordinates[0] > 90 {
		return nil, fmt.Errorf("latitude must be between -90 and 90")
	}
	if dto.Coordinates[1] < -180 || dto.Coordinates[1] > 180 {
		return nil, fmt.Errorf("longitude must be between -180 and 180")
	}
	if dto.Type != "" && dto.Type != "Point" {
		return nil, fmt.Errorf("type must be 'Point', or omit it")
	}

	return &Location{"Point", [2]float64{dto.Coordinates[0], dto.Coordinates[1]}}, nil
}


func createLocation(w http.ResponseWriter, r *http.Request) {

	var locationDto LocationDTO
	
	decoder := json.NewDecoder(r.Body)
	
	if err := decoder.Decode(&locationDto); err != nil {
        fmt.Println(err)
    }

	fmt.Println("locationDto: ", locationDto)
	location, err := NewLocation(locationDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
		
	fmt.Println("location: ", location)





	//TODO: add to db
	

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