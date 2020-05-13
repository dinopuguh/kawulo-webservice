package services

import (
	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllReviews(db *mongo.Database) ([]models.Review, error) {
	ctx := database.Ctx

	crs, err := db.Collection("review").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer crs.Close(ctx)

	result := make([]models.Review, 0)
	for crs.Next(ctx) {
		var row models.Review
		err := crs.Decode(&row)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	err = db.Client().Disconnect(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
