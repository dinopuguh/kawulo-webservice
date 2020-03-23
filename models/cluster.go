package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cluster struct {
	ID           primitive.ObjectID `bson:"_id"`
	LocationId   string             `bson:"location_id"`
	RestaurantId string             `bson:"restaurant_id"`
	Restaurant   Restaurant         `bson:"restaurant"`
	month        int32              `bson:"month"`
	year         int32              `bson:"year"`
	Cluster      int32              `bson:"new_cluster"`
	Service      float64            `bson:"service"`
	Value        float64            `bson:"value"`
	Food         float64            `bson:"food"`
	Vader        float64            `bson:"vader"`
	Wordnet      float64            `bson:"Wordnet"`
	Variance     float64            `bson:"variance"`
	SSE          float64            `bson:"sse"`
}
