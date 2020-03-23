package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/services"
	"github.com/labstack/echo"
)

func FindCluster(ctx echo.Context) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	loc_id := ctx.Param("loc_id")
	month, _ := strconv.Atoi(ctx.Param("month"))
	year, _ := strconv.Atoi(ctx.Param("year"))

	data, err := services.FindClusterByLocation(db, loc_id, int32(month), int32(year))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get clusters")
	}

	return ctx.JSON(http.StatusOK, data)
}
