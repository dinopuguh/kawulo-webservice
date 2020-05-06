package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Prediction struct {
	ID      primitive.ObjectID `bson:"_id" bson:"_id" query:"_id"`
	Month   int32              `bson:"month" bson:"month" query:"month"`
	Year    int32              `bson:"year" bson:"year" query:"year"`
	Service float64            `bson:"service" bson:"service" query:"service"`
	Value   float64            `bson:"value" bson:"value" query:"value"`
	Food    float64            `bson:"food" bson:"food" query:"food"`
	Vader   float64            `bson:"vader" bson:"vader" query:"vader"`
	Wordnet float64            `bson:"wordnet" bson:"wordnet" query:"wordnet"`
}
