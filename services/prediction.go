package services

import (
	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RestaurantPrediction struct {
	Predictions []models.Prediction    `json:"predictions"`
	Error       models.PredictionError `json:"error"`
}

func FindRestaurantPredictions(db *mongo.Database, restId string, nTrain float64, alpha float64) (RestaurantPrediction, error) {
	ctx := database.Ctx

	var result RestaurantPrediction

	filter := bson.M{"restaurant_id": restId, "n_train": nTrain, "alpha": alpha}

	csr, err := db.Collection("prediction").Find(ctx, filter)

	if err != nil {
		return result, err
	}
	defer csr.Close(ctx)

	predictions := make([]models.Prediction, 0)
	for csr.Next(ctx) {
		var row models.Prediction
		err := csr.Decode(&row)
		if err != nil {
			return result, err
		}

		predictions = append(predictions, row)
	}

	var predictionError models.PredictionError
	err = db.Collection("prediction-error").FindOne(ctx, filter).Decode(&predictionError)
	if err != nil {
		return result, err
	}

	err = db.Client().Disconnect(ctx)
	if err != nil {
		return result, err
	}

	result.Predictions = predictions
	result.Error = predictionError

	return result, nil
}

func FindRestaurantBaselinePredictions(db *mongo.Database, restId string, baseline int, alpha float64) (RestaurantPrediction, error) {
	ctx := database.Ctx

	var result RestaurantPrediction

	filter := bson.M{"restaurant_id": restId, "baseline": baseline, "alpha": alpha}

	csr, err := db.Collection("baseline-prediction").Find(ctx, filter)

	if err != nil {
		return result, err
	}
	defer csr.Close(ctx)

	predictions := make([]models.Prediction, 0)
	for csr.Next(ctx) {
		var row models.Prediction
		err := csr.Decode(&row)
		if err != nil {
			return result, err
		}

		predictions = append(predictions, row)
	}

	var predictionError models.PredictionError
	err = db.Collection("baseline-prediction-error").FindOne(ctx, filter).Decode(&predictionError)
	if err != nil {
		return result, err
	}

	err = db.Client().Disconnect(ctx)
	if err != nil {
		return result, err
	}

	result.Predictions = predictions
	result.Error = predictionError

	return result, nil
}
