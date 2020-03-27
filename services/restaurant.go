package services

import (
	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllRestaurants(db *mongo.Database) ([]models.Restaurant, error) {
	ctx := database.Ctx

	csr, err := db.Collection("restaurant").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)

	result := make([]models.Restaurant, 0)
	for csr.Next(ctx) {
		var row models.Restaurant
		err := csr.Decode(&row)
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

func FindRestaurantByLocation(db *mongo.Database, locId string) ([]models.Restaurant, error) {
	ctx := database.Ctx

	csr, err := db.Collection("restaurant").Find(ctx, bson.M{"locationID": locId})
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)

	result := make([]models.Restaurant, 0)
	for csr.Next(ctx) {
		var row models.Restaurant
		err := csr.Decode(&row)
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
