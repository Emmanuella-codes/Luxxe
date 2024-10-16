package entities

import (
	"time"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailStats struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Hour        int                `json:"hour" bson:"hour"`
	Day         int                `json:"day" bson:"day"`
	Month       int                `json:"month" bson:"month"`
	Year        int                `json:"year" bson:"year"`
	HourlyCount int                `json:"hourlyCount" bson:"hourlyCount"`
	DailyCount  int                `json:"dailyCount" bson:"dailyCount"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

var EmailStatsModel *Model[EmailStats]
var EmailStatsCollection *mongo.Collection

func initEmailStats() {
	EmailStatsCollection = config.GetCollection(string(ModelNamesEmailStats))
	EmailStatsModel = InitModel[EmailStats](ModelNamesEmailStats, EmailStatsCollection)
}
