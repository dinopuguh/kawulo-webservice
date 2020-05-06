package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PredictionError struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	NPrediction  int                `bson:"n_prediction" json:"n_prediction"`
	Baseline     int                `bson:"baseline" json:"baseline"`
	NTrain       float64            `bson:"n_train" json:"n_train"`
	Alpha        float64            `bson:"alpha" json:"alpha"`
	RestaurantId string             `bson:"restaurant_id" json:"restaurant_id"`
	ServiceError float64            `bson:"service_error" json:"service_error"`
	ValueError   float64            `bson:"value_error" json:"value_error"`
	FoodError    float64            `bson:"food_error" json:"food_error"`
	VaderError   float64            `bson:"vader_error" json:"vader_error"`
	WordnetError float64            `bson:"wordnet_error" json:"wordnet_error"`
	OverallError float64            `bson:"overall_error" json:"overall_error"`
}
