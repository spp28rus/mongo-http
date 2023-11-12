package database

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

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
