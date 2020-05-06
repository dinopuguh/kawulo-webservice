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
	r.GET("/restaurant/:restId/:month/:year", controllers.GetRestaurantDetail)
	r.GET("/prediction/:restId/:nTrain/:alpha", controllers.GetRestaurantPredictions)
	r.GET("/baseline-prediction/:restId/:baseline/:alpha", controllers.GetRestaurantBaselinePredictions)

	r.Start(":9000")
}
