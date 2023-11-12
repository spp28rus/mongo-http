package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mongo_http/database"
	httppackage "mongo_http/http"
	"net/http"
)

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

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
