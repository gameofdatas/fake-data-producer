package fakedata

import (
	"encoding/json"
	"fmt"
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
func CreateStructFromJSON(jsonData map[string]interface{}) (interface{}, error) {
	fields := make([]reflect.StructField, 0)

	for key, value := range jsonData {
		fieldType, err := determineFieldType(value)
		if err != nil {
			return nil, err
		}

		field := reflect.StructField{
			Name: key,
			Type: fieldType,
		}

		fields = append(fields, field)
	}

	structType := reflect.StructOf(fields)
	if structType == nil {
		return nil, fmt.Errorf("failed to create struct type")
	}

	return reflect.New(structType).Interface(), nil
}

func determineFieldType(value interface{}) (reflect.Type, error) {
	switch t := value.(type) {
	case string:
		switch t {
		case "string":
			return reflect.TypeOf(""), nil
		case "float":
			return reflect.TypeOf(float64(0)), nil
		case "bool":
			return reflect.TypeOf(false), nil
		case "int":
			return reflect.TypeOf(int64(0)), nil
		}
	case map[string]interface{}:
		fieldMap := value.(map[string]interface{})
		if fieldType, ok := fieldMap["type"].(string); ok {
			switch fieldType {
			case "string":
				return reflect.TypeOf(""), nil
			case "int":
				return reflect.TypeOf(int64(0)), nil
			case "float":
				return reflect.TypeOf(float64(0)), nil
			case "timestamp":
				return reflect.TypeOf(""), nil
			case "epoch":
				return reflect.TypeOf(int64(0)), nil
			}
		}
	}

	return reflect.TypeOf(nil), fmt.Errorf("unsupported field type")
}
