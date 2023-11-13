package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofor-little/env"
	"mongo_http/database"
	httppackage "mongo_http/http"
	"net/http"
)

var appPort string

func init() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	appPort = env.Get("APP_PORT", "0")
}

func main() {
	fmt.Println("Hello!")

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/actions/count", func(ctx *gin.Context) {
		var requestData httppackage.RequestCount

		if err := ctx.ShouldBindJSON(&requestData); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

			return
		}

		collection := database.GetMongoCollectionCollection(requestData.DatabaseName, requestData.CollectionName)

		count := database.Count(ctx, collection, requestData.Filter)

		ctx.JSON(http.StatusOK, gin.H{
			"count": count,
		})
	})

	r.POST("/actions/insertOne", func(ctx *gin.Context) {
		var requestData httppackage.RequestInsertOne

		if err := ctx.ShouldBindJSON(&requestData); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

			return
		}

		collection := database.GetMongoCollectionCollection(requestData.DatabaseName, requestData.CollectionName)

		insertedId := database.InsertOne(ctx, collection, requestData.Data)

		ctx.JSON(http.StatusOK, gin.H{
			"inserted_id": insertedId,
		})
	})

	r.POST("/actions/aggregate", func(ctx *gin.Context) {
		var requestData httppackage.RequestAggregate

		if err := ctx.ShouldBindJSON(&requestData); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

			return
		}

		collection := database.GetMongoCollectionCollection(requestData.DatabaseName, requestData.CollectionName)

		elements := database.Aggregate(ctx, collection, requestData.Pipeline)

		ctx.JSON(http.StatusOK, gin.H{
			"elements": elements,
		})
	})

	r.POST("/actions/distinct", func(ctx *gin.Context) {
		var requestData httppackage.RequestDistinct

		if err := ctx.ShouldBindJSON(&requestData); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

			return
		}

		collection := database.GetMongoCollectionCollection(requestData.DatabaseName, requestData.CollectionName)

		elements := database.Distinct(ctx, collection, requestData.FieldName, requestData.Filter)

		ctx.JSON(http.StatusOK, gin.H{
			"elements": elements,
		})
	})

	r.POST("/actions/find", func(ctx *gin.Context) {
		var requestData httppackage.RequestFind

		if err := ctx.ShouldBindJSON(&requestData); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

			return
		}

		collection := database.GetMongoCollectionCollection(requestData.DatabaseName, requestData.CollectionName)

		elements := database.Find(ctx, collection, requestData.Filter)

		ctx.JSON(http.StatusOK, gin.H{
			"elements": elements,
		})
	})

	r.POST("/actions/insertMany", func(ctx *gin.Context) {
		var requestData httppackage.RequestInsertMany

		if err := ctx.ShouldBindJSON(&requestData); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

			return
		}

		collection := database.GetMongoCollectionCollection(requestData.DatabaseName, requestData.CollectionName)

		insertedIDs := database.InsertMany(ctx, collection, requestData.Documents)

		ctx.JSON(http.StatusOK, gin.H{
			"inserted_ids": insertedIDs,
		})
	})

	r.POST("/actions/deleteMany", func(ctx *gin.Context) {
		var requestData httppackage.RequestDeleteMany

		if err := ctx.ShouldBindJSON(&requestData); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

			return
		}

		collection := database.GetMongoCollectionCollection(requestData.DatabaseName, requestData.CollectionName)

		deletedCount := database.DeleteMany(ctx, collection, requestData.Filter)

		ctx.JSON(http.StatusOK, gin.H{
			"deleted_count": deletedCount,
		})
	})

	r.POST("/actions/updateMany", func(ctx *gin.Context) {
		var requestData httppackage.RequestUpdateMany

		if err := ctx.ShouldBindJSON(&requestData); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

			return
		}

		collection := database.GetMongoCollectionCollection(requestData.DatabaseName, requestData.CollectionName)

		updateResult := database.UpdateMany(ctx, collection, requestData.Filter, requestData.Update)

		ctx.JSON(http.StatusOK, gin.H{
			"matched_count":  updateResult.MatchedCount,
			"upserted_count": updateResult.UpsertedCount,
			"modified_count": updateResult.ModifiedCount,
			"upserted_id":    updateResult.UpsertedID,
		})
	})

	r.Run(":" + appPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
