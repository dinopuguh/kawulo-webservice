package controllers

import (
	"log"
	"net/http"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/services"
	"github.com/labstack/echo"
)

func GetRestaurantPredictions(ctx echo.Context) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	rest_id := ctx.Param("rest_id")

	data, err := services.FindRestaurantPrediction(db, rest_id)

	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "Failed to get predictions")
	}

	return ctx.JSON(http.StatusOK, data)
}
