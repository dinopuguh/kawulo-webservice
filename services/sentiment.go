package services

import (
	"math"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindSentimentByRestaurant(db *mongo.Database, restId string, limit int64) ([]models.Sentiment, error) {
	ctx := database.Ctx

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	csr, err := db.Collection("sentiment").Find(ctx, bson.M{"restaurant_id": restId}, findOptions)
	if err != nil {
		return nil, err
	}

	result := make([]models.Sentiment, 0)
	for csr.Next(ctx) {
		var row models.Sentiment
		err := csr.Decode(&row)
		if err != nil {
			return nil, err
		}

		row.Vader = normalizeSentiment(row.Vader)
		row.Wordnet = normalizeSentiment(row.Wordnet)

		result = append(result, row)
	}

	return result, nil
}

func normalizeSentiment(value float64) float64 {
	return math.Round(5 * (value + 1) / 2)
}
