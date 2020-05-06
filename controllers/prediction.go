package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/services"
	"github.com/labstack/echo"
)

func GetRestaurantPredictions(ctx echo.Context) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	restId := ctx.Param("restId")
	nTrain, err := strconv.ParseFloat(ctx.Param("nTrain"), 64)
	alpha, err := strconv.ParseFloat(ctx.Param("alpha"), 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Request value is false")
	}

	data, err := services.FindRestaurantPredictions(db, restId, nTrain, alpha)

	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "Failed to get predictions")
	}

	return ctx.JSON(http.StatusOK, data)
}

func GetRestaurantBaselinePredictions(ctx echo.Context) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	restId := ctx.Param("restId")
	baseline, err := strconv.Atoi(ctx.Param("baseline"))
	alpha, err := strconv.ParseFloat(ctx.Param("alpha"), 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Request value is false")
	}

	data, err := services.FindRestaurantBaselinePredictions(db, restId, baseline, alpha)

	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "Failed to get predictions")
	}

	return ctx.JSON(http.StatusOK, data)
}
