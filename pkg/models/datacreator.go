package models

import (
	"encoding/json"
	"reflect"
)

// CreateData creates data based on the provided dataStruct and configuration.
// It returns the key (nil), value ([]byte), and any error encountered.
func CreateData(dataStruct interface{}, conf map[string]interface{}) (key []byte, value []byte, err error) {
	// Create generic data based on dataStruct and configuration
	data, err := CreateGenericData(dataStruct, conf)
	if err != nil {
		return nil, nil, err
	}

	// Marshal the data to JSON format
	value, err = json.Marshal(data)
	return nil, value, err
}

// CreateStructFromJSON creates a Go struct from a JSON representation.
// It takes a map[string]interface{} jsonData and returns the created struct.
func CreateStructFromJSON(jsonData map[string]interface{}) interface{} {
	// Initialize an empty slice of struct fields
	fields := make([]reflect.StructField, 0)

	// Iterate over the jsonData map
	for key, value := range jsonData {
		// Determine the field type based on the value
		fieldType := determineFieldType(value)
		// Create a new struct field
		field := reflect.StructField{
			Name: key,
			Type: fieldType,
		}

		// Append the field to the fields slice
		fields = append(fields, field)
	}

	// Create a new struct type based on the collected fields
	structType := reflect.StructOf(fields)

	// Create a new instance of the struct and return it
	return reflect.New(structType).Interface()
}

// determineFieldType determines the reflect.Type based on the provided value.
// It supports string, float, bool, int, and nested maps.
// If the type cannot be determined, it returns reflect.TypeOf(nil).
func determineFieldType(value interface{}) reflect.Type {
	switch t := value.(type) {
	case string:
		switch t {
		case "string":
			return reflect.TypeOf("")
		case "float":
			return reflect.TypeOf(float64(0))
		case "bool":
			return reflect.TypeOf(false)
		case "int":
			return reflect.TypeOf(int64(0))
		}
	case map[string]interface{}:
		fieldMap := value.(map[string]interface{})
		if fieldType, ok := fieldMap["type"].(string); ok {
			if fieldType == "int" {
				return reflect.TypeOf(int64(0))
			} else if fieldType == "float" {
				return reflect.TypeOf(float64(0))
			}
		}
	}

	return reflect.TypeOf(nil)
}
