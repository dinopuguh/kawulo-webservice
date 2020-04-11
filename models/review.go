package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID            primitive.ObjectID `bson:"_id"`
	Id            string             `bson:"id"`
	Lang          string             `bson:"lang"`
	LocationId    string             `bson:"location_id"`
	PublishedDate string             `bson:"published_date"`
	Text          string             `bson:"text"`
	Rating        string             `bson:"rating"`
	Subratings    []Subrating        `bson:"subratings"`
}

type Subrating struct {
	Name  string `bson:"name"`
	Value string `bson:"value"`
}
