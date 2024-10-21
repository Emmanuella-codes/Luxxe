package product

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-shared/misc"
)

type mgRepository struct {
	log *log.Logger
}

func newMgRepository(log *log.Logger) ProductRepository {
	return &mgRepository{
		log: log,
	}
}

func (r *mgRepository) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	product.CreatedAt = time.Now()
	result, err := entities.ProductCollection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		fmt.Println("failed to insert")
		return nil, err
	}

	product.ID = insertedID
	return product, nil
}

func (r *mgRepository) QueryAllProducts(ctx context.Context, page int) (*[]entities.Product, int64, error) {

	skip, limit := misc.Pagination(misc.PaginationStruct{Page: page})

	filter := &bson.M{}

	productCount, err := entities.ProductCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetSkip(int64(skip)).SetLimit(int64(limit))
	cursor, err := entities.ProductCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []entities.Product = []entities.Product{}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}
	return &products, productCount, nil
}

func (r *mgRepository) UpdateProductByID(ctx context.Context, 
	productID string, 
	name string, 
	description string,
	price float64,
	category entities.ProductCategories,
	productImage string,
	productInfo string,
	quantity int,
) (*entities.Product, error) {
	update := bson.M{
		"name": 			 	name,
		"description": 	description,
		"price": 			 	price,
		"category": 	 	category,
		"quantity": 	 	quantity,
		"productImage": productImage,
		"productInfo": 	productInfo,
		"updatedAt": 	 	time.Now(),
	}

	return entities.ProductModel.FindByIdAndUpdate(
		ctx,
		productID,
		&bson.M{
			"$set": &update,
		},
	)
}

func (r *mgRepository) QueryByID(ctx context.Context, ID string) (*entities.Product, error) {
	return entities.ProductModel.FindById(ctx, ID)
}

func (r *mgRepository) QueryProductsByCategory(ctx context.Context, category entities.ProductCategories, page int) (*[]entities.Product, int64, error) {
	skip, limit := misc.Pagination(misc.PaginationStruct{Page: page})

	filter := &bson.M{"category": category}

	productCount, err := entities.ProductCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetSkip(int64(skip)).SetLimit(int64(limit))
	cursor, err := entities.ProductCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []entities.Product = []entities.Product{}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}
	return &products, productCount, nil
} 

func (r *mgRepository) DeleteProduct(ctx context.Context, productID string) {
	entities.ProductModel.FindByIdAndDelete(ctx, productID)
}
