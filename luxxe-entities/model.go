package entities 

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type Model[T any] struct {
	name       ModelNames
	collection *mongo.Collection
}

func InitModel[T any](name ModelNames, collection *mongo.Collection) *Model[T] {
	model := &Model[T]{
		name:       name,
		collection: collection,
	}
	return model
}

func (model Model[T]) Create(ctx context.Context, doc *bson.M) (*T, error) {
	res, err := model.collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	var result T
	err = model.collection.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (model Model[T]) FindOne(ctx context.Context, filter *bson.M) (*T, error) {
	var result T
	err := model.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (model Model[T]) FindById(ctx context.Context, id string) (*T, error) {
	var result T
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error occurred creating objecID from %v\n", id)
		// log.Fatal(err)
		return nil, err
	}
	err = model.collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (model Model[T]) Find(ctx context.Context, filter *bson.M) (*[]T, error) {
	var result []T

	cursor, err := model.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (model Model[T]) Count(ctx context.Context, filter *bson.M) (int64, error) {
	return model.collection.CountDocuments(ctx, filter)
}

func (model Model[T]) FindByIdAndUpdateStr(ctx context.Context, id string, jsonString string) (*T, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error occurred creating objecID from %v\n", id)
		// log.Fatal(err)
		return nil, err
	}
	var actualUpdate map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &actualUpdate)
	if err != nil {
		fmt.Printf("json error occurred %v, \n", err)
		log.Fatal(err)
	}
	_, err = model.collection.UpdateByID(ctx, _id, actualUpdate)
	if err != nil {
		fmt.Printf("update error occurred %v, \n", err)
		log.Fatal(err)
	}
	var result T
	err = model.collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (model Model[T]) FindByIdAndUpdate(ctx context.Context, id string, update interface{}) (*T, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error occurred creating objecID from %v\n", id)
		// log.Fatal(err)
		return nil, err
	}
	_, err = model.collection.UpdateByID(ctx, _id, update)
	if err != nil {
		fmt.Printf("update error occurred %v, \n", err)
		log.Fatal(err)
	}
	var result T
	err = model.collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (model Model[T]) FindOneAndUpdateStr(ctx context.Context, filter *bson.M, jsonString string) (*T, error) {
	var actualUpdate map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &actualUpdate)
	if err != nil {
		fmt.Printf("json error occurred %v, \n", err)
		log.Fatal(err)
	}
	res, err := model.collection.UpdateOne(ctx, filter, actualUpdate)
	if err != nil {
		fmt.Printf("update error occurred %v, \n", err)
		log.Fatal(err)
	}
	var result T
	err = model.collection.FindOne(ctx, bson.M{"_id": res.UpsertedID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (model Model[T]) FindOneAndUpdate(ctx context.Context, filter *bson.M, update interface{}) (*T, error) {
	res, err := model.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Printf("find one and update error occurred %v, \n", err)
		// log.Fatal(err)
		return nil, err
	}

	var result T
	err = model.collection.FindOne(ctx, bson.M{"_id": res.UpsertedID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (model Model[T]) FindByIdAndDelete(ctx context.Context, id string) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error occurred creating objecID from %v\n", id)
		// log.Fatal(err)
		// return nil, err
	} else {
		model.collection.FindOneAndDelete(ctx, map[string]interface{}{"_id": _id})
	}
}

func (model *Model[T]) FindOneAndDelete(ctx context.Context, filter interface{}) {
	model.collection.FindOneAndDelete(ctx, filter)
}

func (model *Model[T]) UpdateMany(ctx context.Context, filter *bson.M, update interface{}) {
	res, err := model.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		fmt.Printf("update many error occurred %v, \n", err)
		// log.Fatal(err)
	}

	fmt.Println(res)
}
