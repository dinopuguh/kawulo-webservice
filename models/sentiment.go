package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sentiment struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id" query:"_id"`
	PublishedDate  string             `bson:"published_date" json:"published_date" query:"published_date"`
	Month          int32              `bson:"month" json:"month" query:"month"`
	Year           int32              `bson:"year" json:"year" query:"year"`
	TranslatedText string             `bson:"translated_text" json:"translated_text" query:"translated_text"`
	Service        float64            `bson:"service" json:"service" query:"service"`
	Value          float64            `bson:"value" json:"value" query:"value"`
	Food           float64            `bson:"food" json:"food" query:"food"`
	Vader          float64            `bson:"vader" json:"vader" query:"vader"`
	Wordnet        float64            `bson:"wordnet" json:"wordnet" query:"wordnet"`
}
