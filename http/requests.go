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

type RequestDistinct struct {
	DatabaseName   string `json:"database_name" binding:"required"`
	CollectionName string `json:"collection_name" binding:"required"`
	FieldName      string `json:"field_name" binding:"required"`
	Filter         string `json:"filter" binding:"required"`
}

type RequestFind struct {
	DatabaseName   string `json:"database_name" binding:"required"`
	CollectionName string `json:"collection_name" binding:"required"`
	Filter         string `json:"filter" binding:"required"`
}

type RequestInsertMany struct {
	DatabaseName   string `json:"database_name" binding:"required"`
	CollectionName string `json:"collection_name" binding:"required"`
	Documents      string `json:"documents" binding:"required"`
}

type RequestDeleteMany struct {
	DatabaseName   string `json:"database_name" binding:"required"`
	CollectionName string `json:"collection_name" binding:"required"`
	Filter         string `json:"filter" binding:"required"`
}

type RequestUpdateMany struct {
	DatabaseName   string `json:"database_name" binding:"required"`
	CollectionName string `json:"collection_name" binding:"required"`
	Filter         string `json:"filter" binding:"required"`
	Update         string `json:"update" binding:"required"`
}
type RequestBulkWriteUpdateOne struct {
	DatabaseName   string `json:"database_name" binding:"required"`
	CollectionName string `json:"collection_name" binding:"required"`
	Operations     string `json:"operations" binding:"required"`
}
