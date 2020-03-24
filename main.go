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
	r.GET("/location/:loc_id", controllers.GetLocationById)
	r.GET("/cluster/:loc_id/:month/:year", controllers.GetRestaurantClusters)
	r.GET("/prediction/:rest_id", controllers.GetRestaurantPredictions)

	r.Start(":9000")
}
