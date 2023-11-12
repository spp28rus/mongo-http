package http

type RequestCount struct {
	DatabaseName   string `json:"database_name" binding:"required"`
	CollectionName string `json:"collection_name" binding:"required"`
	Filter         string `json:"filter" binding:"required"`
}

type RequestInsertOne struct {
	DatabaseName   string `json:"database_name" binding:"required"`
	CollectionName string `json:"collection_name" binding:"required"`
	Data           string `json:"data" binding:"required"`
}

type RequestAggregate struct {
	DatabaseName   string `json:"database_name" binding:"required"`
	CollectionName string `json:"collection_name" binding:"required"`
	Pipeline       string `json:"pipeline" binding:"required"`
}
