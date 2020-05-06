package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Temporal struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id" query:"_id"`
	Month   int32              `bson:"month" json:"month" query:"month"`
	Year    int32              `bson:"year" json:"year" query:"year"`
	Service float64            `bson:"service" json:"service" query:"service"`
	Value   float64            `bson:"value" json:"value" query:"value"`
	Food    float64            `bson:"food" json:"food" query:"food"`
	Vader   float64            `bson:"vader" json:"vader" query:"vader"`
	Wordnet float64            `bson:"wordnet" json:"wordnet" query:"wordnet"`
}
