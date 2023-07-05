package models

import (
	"encoding/json"
	"reflect"
)

func CreateData(dataStruct interface{}, conf map[string]interface{}) (key []byte, value []byte, err error) {
	data, err := CreateGenericData(dataStruct, conf)
	if err != nil {
		return nil, nil, err
	}
	value, err = json.Marshal(data)
	return nil, value, err
}

func CreateStructFromJSON(jsonData map[string]interface{}) interface{} {
	fields := make([]reflect.StructField, 0)

	for key, value := range jsonData {
		fieldType := determineFieldType(value)

		field := reflect.StructField{
			Name: key,
			Type: fieldType,
		}

		fields = append(fields, field)
	}

	structType := reflect.StructOf(fields)
	return reflect.New(structType).Interface()
}

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
