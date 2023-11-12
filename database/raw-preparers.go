package database

import (
	"encoding/json"
)

func PrepareObjectRawValue(data string) interface{} {
	var valueInterface interface{}

	if err := json.Unmarshal([]byte(data), &valueInterface); err != nil {
		panic(err)
	}

	return valueInterface
}

func PrepareArrayRawValue(data string) []interface{} {
	var valueInterface []interface{}

	if err := json.Unmarshal([]byte(data), &valueInterface); err != nil {
		panic(err)
	}

	return valueInterface
}
