package entities

import (
	"time"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductCategories string

const (
	HairCategory 									ProductCategories = "hair"
	SkincareCategory 							ProductCategories = "skincare"
	FragranceCategory 						ProductCategories = "fragrance"
	WellnessCategory 							ProductCategories = "wellness"
	MakeupCategory 								ProductCategories = "makeup"
	FashionAccessoriesCategory 		ProductCategories = "accessories"
	PersonalCareCategory 					ProductCategories = "personalcare"
	BathAndBodyCategory 					ProductCategories = "bath&body"
	JewelryCategory 							ProductCategories = "jewelry"
	HomeSpaCategory 							ProductCategories = "homespa"
)

type Product struct {
	ID            primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
	Name          string              `json:"name" bson:"name"`
	Description   string              `json:"description" bson:"description"`
	Price         float64             `json:"price" bson:"price"`
	Category      ProductCategories   `json:"category" bson:"category"`
	Quantity      int                 `json:"quantity" bson:"quantity"`
	ProductImage	string							`json:"productImage" bson:"productImage"`
	ProductInfo		string							`json:"productInfo" bson:"productInfo"`
	CreatedAt     time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt" bson:"updatedAt"`
}

var ProductModel *Model[Product]
var ProductCollection *mongo.Collection

func initProduct() {
  ProductCollection = config.GetCollection(string(ModelNamesProduct))
  ProductModel = InitModel[Product](ModelNamesProduct, ProductCollection)
}
