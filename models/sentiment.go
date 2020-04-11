package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sentiment struct {
	ID             primitive.ObjectID `bson:"_id"`
	PublishedDate  string             `bson:"published_date"`
	Month          int32              `bson:"month"`
	Year           int32              `bson:"year"`
	TranslatedText string             `bson:"translated_text"`
	Service        float64            `bson:"service"`
	Value          float64            `bson:"value"`
	Food           float64            `bson:"food"`
	Vader          float64            `bson:"vader"`
	Wordnet        float64            `bson:"wordnet"`
}
