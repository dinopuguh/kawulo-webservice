package main

import (
	"github.com/dinopuguh/kawulo-webservice/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	r := echo.New()

	r.Use(middleware.CORS())

	r.GET("/location/search/:query", controllers.SearchLocation)
	r.GET("/location/:locId", controllers.GetLocationById)
	r.GET("/cluster/:locId/:month/:year", controllers.GetRestaurantClusters)
	r.GET("/prediction/:restId", controllers.GetRestaurantPredictions)

	r.Start(":9000")
}
