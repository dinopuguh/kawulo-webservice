package services

import (
	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindClusterByLocation(db *mongo.Database, locId string, month int32, year int32) ([]models.Cluster, error) {
	ctx := database.Ctx

	filter := bson.M{"location_id": locId, "month": month, "year": year}
	csr, err := db.Collection("cluster").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)

	result := make([]models.Cluster, 0)
	for csr.Next(ctx) {
		var row models.Cluster
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

func FindClusterByRestaurant(db *mongo.Database, restId string, month int32, year int32) (models.Cluster, error) {
	ctx := database.Ctx

	var result models.Cluster

	filter := bson.M{"restaurant_id": restId, "month": month, "year": year}
	err := db.Collection("cluster").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
