package models

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
   Id       primitive.ObjectID `json:"id,omitempty"`
   Type 	  string    `json:"type,omitempty"`
   Coordinates [2]float64 `json:"coordinates"`
}

type LocationQuery struct {
	Coordinates []string `query:"coordinates"`
	Radius 	string    `query:"radius"`
}

type SearchLocation struct {
	Coordinates [2]float64 
	Radius 	float64  
}

func NewSearchLocation(query LocationQuery) (*SearchLocation, error){
	latitude, err := strconv.ParseFloat(query.Coordinates[0], 64)
	if	err != nil{
		return nil, err
	}
	
	longitude, err := strconv.ParseFloat(query.Coordinates[1], 64)
	if	err != nil{
		return nil, err
	}

	radius, err := strconv.ParseFloat(query.Radius, 64)
	if err != nil{
		return nil, err
	}

	return &SearchLocation{
		Coordinates: [2]float64{latitude, longitude},
		Radius: radius,
	}, nil
}

func NewFilter(query LocationQuery) (primitive.M, error){

	searchLocation, err := NewSearchLocation(query)
	if err != nil{
		return nil, err
	}


	return bson.M{"location": 
			 bson.M{"$near": 
				bson.M{"$geometry": bson.M{"type": "Point", "coordinates": bson.A{ searchLocation.Coordinates[0], searchLocation.Coordinates[1]}}, 
					   "$maxDistance": searchLocation.Radius},
					},
				}, nil
}
