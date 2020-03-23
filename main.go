package main

import (
	"github.com/dinopuguh/kawulo-webservice/controllers"
	"github.com/labstack/echo"
)

type M map[string]interface{}

func main() {
	r := echo.New()

	r.GET("/location/search/:query", controllers.SearchLocation)
	r.GET("/cluster/:loc_id/:month/:year", controllers.FindCluster)

	r.Start(":9000")
}
