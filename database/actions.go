package database

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func Aggregate(ctx *gin.Context, collection *mongo.Collection, pipeline string) []bson.M {
	pipelineInstance := PrepareArrayRawValue(pipeline)

	cursor, err := collection.Aggregate(ctx, pipelineInstance)

	if err != nil {
		panic(err)
	}

	var docs []bson.M

	err = cursor.All(ctx, &docs)

	if err != nil {
		panic(err)
	}

	return docs
}

func Distinct(ctx *gin.Context, collection *mongo.Collection, fieldName string, filter string) []interface{} {
	filterInstance := PrepareObjectRawValue(filter)

	elements, err := collection.Distinct(ctx, fieldName, filterInstance)

	if err != nil {
		panic(err)
	}

	return elements
}

func Find(ctx *gin.Context, collection *mongo.Collection, filter string) []bson.M {
	filterInstance := PrepareObjectRawValue(filter)

	cursor, err := collection.Find(ctx, filterInstance)

	if err != nil {
		panic(err)
	}

	var docs []bson.M

	err = cursor.All(ctx, &docs)

	if err != nil {
		panic(err)
	}

	return docs
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

func BulkWriteUpdateOne(ctx *gin.Context, collection *mongo.Collection, operations string) *mongo.BulkWriteResult {
	operationsInstance := PrepareArrayOfArraysRawValue(operations)

	var models []mongo.WriteModel

	for _, element := range operationsInstance {
		models = append(models, mongo.NewUpdateOneModel().SetFilter(element[0]).SetUpdate(element[1]))
	}

	bulkWriteResult, err := collection.BulkWrite(ctx, models)

	if err != nil {
		panic(err)
	}

	return bulkWriteResult
}
