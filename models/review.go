package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id" query:"_id"`
	Id            string             `bson:"id" json:"id" query:"id"`
	Lang          string             `bson:"lang" json:"lang" query:"lang"`
	LocationId    string             `bson:"location_id" json:"location_id" query:"location_id"`
	PublishedDate string             `bson:"published_date" json:"published_date" query:"published_date"`
	Text          string             `bson:"text" json:"text" query:"text"`
	Rating        string             `bson:"rating" json:"rating" query:"rating"`
	Subratings    []Subrating        `bson:"subratings" json:"subratings" query:"subratings"`
}

type Subrating struct {
	Name  string `bson:"name" json:"name" query:"name"`
	Value string `bson:"value" json:"value" query:"value"`
}
