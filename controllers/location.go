package controllers

import (
	"log"
	"net/http"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/services"
	"github.com/labstack/echo"
)

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
