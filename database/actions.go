package database

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: prepare _id as object for a filtering

func Count(ctx *gin.Context, collection *mongo.Collection, filter string) int64 {
	filterInstance := PrepareObjectRawValue(filter)

	count, err := collection.CountDocuments(ctx, filterInstance)

	if err != nil {
		panic(err)
	}

	return count
}

func InsertOne(ctx *gin.Context, collection *mongo.Collection, data string) interface{} {
	dataInstance := PrepareObjectRawValue(data)

	one, err := collection.InsertOne(ctx, dataInstance)

	if err != nil {
		panic(err)
	}

	return one.InsertedID
}

func Aggregate(ctx *gin.Context, collection *mongo.Collection, pipeline string) []interface{} {
	pipelineInstance := PrepareArrayRawValue(pipeline)

	cursor, err := collection.Aggregate(ctx, pipelineInstance)

	if err != nil {
		panic(err)
	}

	var elements []interface{}

	err = cursor.All(ctx, &elements)

	if err != nil {
		panic(err)
	}

	return elements
}

func Distinct(ctx *gin.Context, collection *mongo.Collection, fieldName string, filter string) []interface{} {
	filterInstance := PrepareObjectRawValue(filter)

	elements, err := collection.Distinct(ctx, fieldName, filterInstance)

	if err != nil {
		panic(err)
	}

	return elements
}

func Find(ctx *gin.Context, collection *mongo.Collection, filter string) []interface{} {
	filterInstance := PrepareObjectRawValue(filter)

	cursor, err := collection.Find(ctx, filterInstance)

	if err != nil {
		panic(err)
	}

	var elements []interface{}

	err = cursor.All(ctx, &elements)

	if err != nil {
		panic(err)
	}

	return elements
}

func InsertMany(ctx *gin.Context, collection *mongo.Collection, documents string) []interface{} {
	documentsInstance := PrepareArrayRawValue(documents)

	insertManyResult, err := collection.InsertMany(ctx, documentsInstance)

	if err != nil {
		panic(err)
	}

	return insertManyResult.InsertedIDs
}

func DeleteMany(ctx *gin.Context, collection *mongo.Collection, filter string) int64 {
	filterInstance := PrepareObjectRawValue(filter)

	deleteResult, err := collection.DeleteMany(ctx, filterInstance)

	if err != nil {
		panic(err)
	}

	return deleteResult.DeletedCount
}

func UpdateMany(ctx *gin.Context, collection *mongo.Collection, filter string, update string) *mongo.UpdateResult {
	filterInstance := PrepareObjectRawValue(filter)
	updateInstance := PrepareObjectRawValue(update)

	updateResult, err := collection.UpdateMany(ctx, filterInstance, updateInstance)

	if err != nil {
		panic(err)
	}

	return updateResult
}
