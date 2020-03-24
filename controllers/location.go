package controllers

import (
	"log"
	"net/http"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/services"
	"github.com/labstack/echo"
)

func GetLocationById(ctx echo.Context) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	loc_id := ctx.Param("loc_id")

	data, err := services.FindLocationById(db, loc_id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get location")
	}

	return ctx.JSON(http.StatusOK, data)
}

func SearchLocation(ctx echo.Context) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	query := ctx.Param("query")

	data, err := services.FindLocationByQuery(db, query)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get locations")
	}

	return ctx.JSON(http.StatusOK, data)
}
