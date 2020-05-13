package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"github.com/dinopuguh/kawulo-webservice/services"
	"github.com/labstack/echo"
)

type RestaurantDetail struct {
	Cluster      models.Cluster            `json:"cluster"`
	Temporal     []models.Temporal         `json:"temporal"`
	Sentiments   []models.Sentiment        `json:"sentiments"`
	ServiceStats Stats                     `json:"service_stats"`
	ValueStats   Stats                     `json:"value_stats"`
	FoodStats    Stats                     `json:"food_stats"`
	VaderStats   Stats                     `json:"vader_stats"`
	WordnetStats Stats                     `json:"wordnet_stats"`
	Heatmap      []models.SentimentHeatmap `json:"heatmap"`
}

type Stats struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	Neutral  int `json:"neutral"`
}

func GetRestaurantDetail(ctx echo.Context) error {
	mongoCtx := database.Ctx
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	restId := ctx.Param("restId")
	month, _ := strconv.Atoi(ctx.Param("month"))
	year, _ := strconv.Atoi(ctx.Param("year"))

	cluster, err := services.FindClusterByRestaurant(db, restId, int32(month), int32(year))
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "Failed to get restaurant detail")
	}

	temporal, err := services.FindTemporalByRestaurant(db, restId)
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "Failed to get restaurant temporal data")
	}

	sentiments, err := services.FindSentimentByRestaurant(db, restId, 20)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "Failed to get restaurant sentiment data")
	}

	heatmap, err := services.CountSentimentsHeatmap(db, restId)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "Failed to get restaurant review heatmap data")
	}

	serviceStats := &Stats{
		Positive: 0,
		Negative: 0,
		Neutral:  0,
	}
	valueStats := &Stats{
		Positive: 0,
		Negative: 0,
		Neutral:  0,
	}
	foodStats := &Stats{
		Positive: 0,
		Negative: 0,
		Neutral:  0,
	}
	vaderStats := &Stats{
		Positive: 0,
		Negative: 0,
		Neutral:  0,
	}
	wordnetStats := &Stats{
		Positive: 0,
		Negative: 0,
		Neutral:  0,
	}

	for _, temp := range temporal {
		if temp.Service == 3 {
			serviceStats.Neutral += 1
		} else if temp.Service > 3 {
			serviceStats.Positive += 1
		} else {
			serviceStats.Negative += 1
		}
		if temp.Value == 3 {
			valueStats.Neutral += 1
		} else if temp.Food > 3 {
			valueStats.Positive += 1
		} else {
			valueStats.Negative += 1
		}
		if temp.Food == 3 {
			foodStats.Neutral += 1
		} else if temp.Food > 3 {
			foodStats.Positive += 1
		} else {
			foodStats.Negative += 1
		}
		if temp.Vader == 0 {
			vaderStats.Neutral += 1
		} else if temp.Vader > 0 {
			vaderStats.Positive += 1
		} else {
			vaderStats.Negative += 1
		}
		if temp.Wordnet == 0 {
			wordnetStats.Neutral += 1
		} else if temp.Wordnet > 0 {
			wordnetStats.Positive += 1
		} else {
			wordnetStats.Negative += 1
		}
	}

	data := RestaurantDetail{
		Cluster:      cluster,
		Temporal:     temporal,
		Sentiments:   sentiments,
		ServiceStats: *serviceStats,
		ValueStats:   *valueStats,
		FoodStats:    *foodStats,
		VaderStats:   *vaderStats,
		WordnetStats: *wordnetStats,
		Heatmap:      heatmap,
	}

	err = db.Client().Disconnect(mongoCtx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to close mongo connection")
	}

	return ctx.JSON(http.StatusOK, data)
}
