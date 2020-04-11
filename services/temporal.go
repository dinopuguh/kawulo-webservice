package services

import (
	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindTemporalByRestaurant(db *mongo.Database, restId string) ([]models.Temporal, error) {
	ctx := database.Ctx

	csr, err := db.Collection("temporal").Find(ctx, bson.M{"restaurant_id": restId})
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)

	result := make([]models.Temporal, 0)
	for csr.Next(ctx) {
		var row models.Temporal
		err := csr.Decode(&row)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}
