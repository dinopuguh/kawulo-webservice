package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dinopuguh/kawulo-webservice/database"
	"github.com/dinopuguh/kawulo-webservice/models"
	"github.com/dinopuguh/kawulo-webservice/response"
	"github.com/dinopuguh/kawulo-webservice/services"
	"github.com/gemcook/pagination-go"
	"github.com/labstack/echo"
)

type clusterRepository struct {
	locId string
	month int
	year  int
}

func newClusterRepository() *clusterRepository {
	return &clusterRepository{
		locId: "297715",
		month: 4,
		year:  2020,
	}
}

func (fr *clusterRepository) GetCluster(orders []*pagination.Order) []models.Cluster {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	locId := fr.locId
	month := fr.month
	year := fr.year

	result, err := services.FindClusterByLocation(db, locId, int32(month), int32(year))

	return result
}

type clusterCondition struct {
	LocationId *string
	Month      *int
	Year       *int
}

func newClusterCondition(location_id string, month, year int) *clusterCondition {
	return &clusterCondition{
		LocationId: &location_id,
		Month:      &month,
		Year:       &year,
	}
}

func parseClusterCondition(ctx echo.Context) *clusterCondition {
	locId := ctx.Param("locId")
	month, _ := strconv.Atoi(ctx.Param("month"))
	year, _ := strconv.Atoi(ctx.Param("year"))

	return newClusterCondition(locId, month, year)
}

type clusterFetcher struct {
	repo *clusterRepository
}

func newClusterFetcher() *clusterFetcher {
	return &clusterFetcher{
		repo: &clusterRepository{},
	}
}

func (cf *clusterFetcher) applyCondition(cond *clusterCondition) {
	if cond.LocationId != nil {
		cf.repo.locId = *cond.LocationId
	}
	if cond.Month != nil {
		cf.repo.month = *cond.Month
	}
	if cond.Year != nil {
		cf.repo.year = *cond.Year
	}
}

func (cf *clusterFetcher) Count(cond interface{}) (int, error) {
	if cond != nil {
		cf.applyCondition(cond.(*clusterCondition))
	}
	orders := make([]*pagination.Order, 0, 0)
	clusters := cf.repo.GetCluster(orders)
	return len(clusters), nil
}

func (cf *clusterFetcher) FetchPage(cond interface{}, input *pagination.PageFetchInput, result *pagination.PageFetchResult) error {
	if cond != nil {
		cf.applyCondition(cond.(*clusterCondition))
	}
	clusters := cf.repo.GetCluster(input.Orders)
	var toIndex int
	toIndex = input.Offset + input.Limit
	if toIndex > len(clusters) {
		toIndex = len(clusters)
	}
	for _, cluster := range clusters[input.Offset:toIndex] {
		*result = append(*result, cluster)
	}
	return nil
}

func GetRestaurantClusters(ctx echo.Context) error {
	p := pagination.ParseQuery(ctx.Request().URL.RequestURI())
	cond := parseClusterCondition(ctx)
	fetcher := newClusterFetcher()

	totalCount, totalPages, res, err := pagination.Fetch(fetcher, &pagination.Setting{
		Limit:  p.Limit,
		Page:   p.Page,
		Cond:   cond,
		Orders: p.Sort,
	})

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get clusters")
	}

	result := &response.PaginationResponse{
		Pages:      res.Pages,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}

	return ctx.JSON(http.StatusOK, result)
}
