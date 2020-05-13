package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cluster struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id" query:"_id"`
	LocationId   string             `bson:"location_id" json:"location_id" query:"location_id"`
	RestaurantId string             `bson:"restaurant_id" json:"restaurant_id" query:"restaurant_id"`
	Restaurant   Restaurant         `bson:"restaurant" json:"restaurant" query:"restaurant"`
	Month        int32              `bson:"month" json:"month" query:"month"`
	Year         int32              `bson:"year" json:"year" query:"year"`
	Cluster      int32              `bson:"new_cluster" json:"cluster" query:"cluster"`
	Service      float64            `bson:"service" json:"service" query:"service"`
	Value        float64            `bson:"value" json:"value" query:"value"`
	Food         float64            `bson:"food" json:"food" query:"food"`
	Vader        float64            `bson:"vader" json:"vader" query:"vader"`
	Wordnet      float64            `bson:"wordnet" json:"wordnet" query:"wordnet"`
	Variance     float64            `bson:"variance" json:"variance" query:"variance"`
	SSE          float64            `bson:"sse" json:"sse" query:"sse"`
}
