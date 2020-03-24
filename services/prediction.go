package services

import (
	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindRestaurantPrediction(db *mongo.Database, rest_id string) (models.Prediction, error) {
	ctx := database.Ctx

	var result models.Prediction
	err := db.Collection("prediction").FindOne(ctx, bson.M{"restaurant_id": rest_id}).Decode(&result)
	if err != nil {
		return result, err
	}

	err = db.Client().Disconnect(ctx)
	if err != nil {
		return result, err
	}

	return result, nil
}
