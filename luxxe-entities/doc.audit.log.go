package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
)

type AuditLog struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	RequestIP   string             `json:"requestIP" bson:"requestIP"`
	QueryParams map[string]string  `json:"queryParams" bson:"queryParams"`
	OriginalURL string             `json:"originalURL" bson:"originalURL"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}

var AuditLogModel *Model[AuditLog]
var AuditLogCollection *mongo.Collection

func initAuditLog() {
	AuditLogCollection = config.GetCollection(string(ModelNamesAuditLog))
	AuditLogModel = InitModel[AuditLog](ModelNamesAuditLog, AuditLogCollection)
}
