package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"github.com/dinopuguh/kawulo-webservice/services"
	"github.com/labstack/echo"
)

type RestaurantDetail struct {
	Cluster    models.Cluster
	Temporal   []models.Temporal
	Prediction models.Prediction
	Sentiments []models.Sentiment
}

func GetRestaurantDetail(ctx echo.Context) error {
	mongoCtx := database.Ctx
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	restId := ctx.Param("restId")
	month, _ := strconv.Atoi(ctx.Param("month"))
	year, _ := strconv.Atoi(ctx.Param("year"))

	cluster, err := services.FindClusterByRestaurant(db, restId, int32(month), int32(year))
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "Failed to get restaurant detail")
	}

	temporal, err := services.FindTemporalByRestaurant(db, restId)
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "Failed to get restaurant temporal data")
	}

	prediction, err := services.FindRestaurantPrediction(db, restId)
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "Failed to get restaurant prediction")
	}

	sentiments, err := services.FindSentimentByRestaurant(db, restId, 10)
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "Failed to get restaurant reviews")
	}

	data := RestaurantDetail{
		Cluster:    cluster,
		Temporal:   temporal,
		Prediction: prediction,
		Sentiments: sentiments,
	}

	err = db.Client().Disconnect(mongoCtx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to close mongo connection")
	}

	return ctx.JSON(http.StatusOK, data)
}
