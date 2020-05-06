package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Prediction struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id" query:"_id"`
	Baseline int                `bson:"baseline" json:"baseline" query:"baseline"`
	NTrain   float64            `bson:"n_train" json:"n_train" query:"n_train"`
	Alpha    float64            `bson:"alpha" json:"alpha" query:"alpha"`
	Month    int32              `bson:"month" json:"month" query:"month"`
	Year     int32              `bson:"year" json:"year" query:"year"`
	Service  float64            `bson:"service" json:"service" query:"service"`
	Value    float64            `bson:"value" json:"value" query:"value"`
	Food     float64            `bson:"food" json:"food" query:"food"`
	Vader    float64            `bson:"vader" json:"vader" query:"vader"`
	Wordnet  float64            `bson:"wordnet" json:"wordnet" query:"wordnet"`
}
