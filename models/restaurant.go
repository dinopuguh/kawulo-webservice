package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID                 primitive.ObjectID `bson:"_id" json:"_id" query:"_id"`
	LocationId         string             `bson:"location_id" json:"location_id" query:"location_id"`
	Name               string             `bson:"name" json:"name" query:"name"`
	Latitude           string             `bson:"latitude" json:"latitude" query:"latitude"`
	Longitude          string             `bson:"longitude" json:"longitude" query:"longitude"`
	NumReviews         string             `bson:"num_reviews" json:"num_reviews" query:"num_reviews"`
	Photo              Photo              `bson:"photo" json:"photo" query:"photo"`
	Rating             string             `bson:"rating" json:"rating" query:"rating"`
	PriceLevel         string             `bson:"price_level" json:"price_level" query:"price_level"`
	Price              string             `bson:"price" json:"price" query:"price"`
	Address            string             `bson:"address" json:"address" query:"address"`
	Phone              string             `bson:"phone" json:"phone" query:"phone"`
	Website            string             `bson:"website" json:"website" query:"website"`
	RawRanking         string             `bson:"raw_ranking" json:"raw_ranking" query:"raw_ranking"`
	RankingGeo         string             `bson:"ranking_geo" json:"ranking_geo" query:"ranking_geo"`
	RankingPosition    string             `bson:"ranking_position" json:"ranking_position" query:"ranking_position"`
	RankingDenominator string             `bson:"ranking_denominator" json:"ranking_denominator" query:"ranking_denominator"`
	RankingCategory    string             `bson:"ranking_category" json:"ranking_category" query:"ranking_category"`
	Ranking            string             `bson:"ranking" json:"ranking" query:"ranking"`
	SubCategory        []SubCategory      `bson:"subcategory" json:"subcategory" query:"subcategory"`
	LocationID         string             `bson:"locationID" json:"locationID" query:"locationID"`
	LocationObjectID   primitive.ObjectID `bson:"location_ObjectId" json:"location_ObjectId" query:"location_ObjectId"`
	CreatedAt          time.Time          `bson:"created_at" json:"created_at" query:"created_at"`
}

type Photo struct {
	Images Images `bson:"images" json:"images" query:"images"`
}

type Images struct {
	Thumbnail Image `bson:"thumbnail" json:"thumbnail" query:"thumbnail"`
	Original  Image `bson:"original" json:"original" query:"original"`
	Medium    Image `bson:"medium" json:"medium" query:"medium"`
	Large     Image `bson:"large" json:"large" query:"large"`
}

type Image struct {
	Width  string `bson:"width" json:"width" query:"width"`
	Url    string `bson:"url" json:"url" query:"url"`
	Height string `bson:"height" json:"height" query:"height"`
}

type SubCategory struct {
	Key  string `bson:"key" json:"key" query:"key"`
	Name string `bson:"name" json:"name" query:"name"`
}
