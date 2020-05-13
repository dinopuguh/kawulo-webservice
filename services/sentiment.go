package services

import (
	"math"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindSentimentByRestaurant(db *mongo.Database, restId string, limit int64) ([]models.Sentiment, error) {
	ctx := database.Ctx

	findOptions := options.Find()
	if limit != 0 {
		findOptions.SetLimit(limit)
	}

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

func FindReviewByRestaurant(db *mongo.Database, restId string) ([]models.Review, error) {
	ctx := database.Ctx

	csr, err := db.Collection("review").Find(ctx, bson.M{"restaurant_locationID": restId})
	if err != nil {
		return nil, err
	}

	result := make([]models.Review, 0)
	for csr.Next(ctx) {
		var row models.Review
		err := csr.Decode(&row)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}

func CountSentimentsHeatmap(db *mongo.Database, restId string) ([]models.SentimentHeatmap, error) {
	ctx := database.Ctx

	matchStage := bson.D{
		primitive.E{
			Key: "$match",
			Value: bson.D{
				primitive.E{
					Key:   "restaurant_id",
					Value: restId,
				},
			},
		},
	}
	groupStage := bson.D{
		primitive.E{
			Key: "$group",
			Value: bson.D{
				primitive.E{
					Key: "_id",
					Value: bson.D{
						primitive.E{
							Key:   "restaurant_id",
							Value: "$restaurant_id",
						},
						primitive.E{
							Key:   "month",
							Value: "$month",
						},
						primitive.E{
							Key:   "year",
							Value: "$year",
						},
					},
				},
				primitive.E{
					Key: "month",
					Value: bson.D{
						primitive.E{
							Key:   "$first",
							Value: "$month",
						},
					},
				},
				primitive.E{
					Key: "year",
					Value: bson.D{
						primitive.E{
							Key:   "$first",
							Value: "$year",
						},
					},
				},
				primitive.E{
					Key: "count",
					Value: bson.D{
						primitive.E{
							Key:   "$sum",
							Value: 1,
						},
					},
				},
			},
		},
	}
	sortStage := bson.D{
		primitive.E{
			Key: "$sort",
			Value: bson.D{
				primitive.E{
					Key:   "_id.year",
					Value: 1,
				},
				primitive.E{
					Key:   "_id.month",
					Value: 1,
				},
			},
		},
	}

	var result []models.SentimentHeatmap

	csr, err := db.Collection("sentiment").Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, sortStage})
	if err != nil {
		return result, err
	}
	defer csr.Close(ctx)

	if err = csr.All(ctx, &result); err != nil {
		return result, err
	}

	return result, nil
}
