package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id" query:"_id"`
	Name       string             `bson:"name" json:"name" query:"name"`
	LocationId string             `bson:"location_id" json:"location_id" query:"location_id"`
	Latitude   string             `bson:"latitude" json:"latitude" query:"latitude"`
	Longitude  string             `bson:"longitude" json:"longitude" query:"longitude"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at" query:"created_at"`
}
