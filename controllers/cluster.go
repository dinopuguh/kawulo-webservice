package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/services"
	"github.com/labstack/echo"
)

func GetRestaurantClusters(ctx echo.Context) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	locId := ctx.Param("locId")
	month, _ := strconv.Atoi(ctx.Param("month"))
	year, _ := strconv.Atoi(ctx.Param("year"))

	data, err := services.FindClusterByLocation(db, locId, int32(month), int32(year))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get clusters")
	}

	return ctx.JSON(http.StatusOK, data)
}
