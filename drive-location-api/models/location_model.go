package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
   Id       primitive.ObjectID `json:"id,omitempty"`
   Type 	  string    `json:"type,omitempty"`
   Coordinates [2]float64 `json:"coordinates"`
}