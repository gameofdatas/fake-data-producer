package models

import (
	"fmt"
	"math"
	"reflect"

	"github.com/brianvoe/gofakeit/v6"
)

func CreateGenericData(dataStruct interface{}, confData map[string]interface{}) (interface{}, error) {
	generateFakeData(dataStruct, confData)
	return dataStruct, nil
}

func generateFakeData(dataStruct interface{}, jsonData map[string]interface{}) {
	gofakeit.Seed(0)

	val := reflect.ValueOf(dataStruct)
	if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		val = val.Elem()

		for i := 0; i < val.NumField(); i++ {
			fieldValue := val.Field(i)
			fieldType := val.Type().Field(i)
			jsonFieldType := jsonData[fieldType.Name]
			switch fieldValue.Kind() {
			case reflect.String:
				setStringValue(fieldValue, fieldType)
			case reflect.Float64:
				setFloatValue(fieldValue, fieldType, jsonFieldType)
			case reflect.Bool:
				fieldValue.SetBool(gofakeit.Bool())
			case reflect.Int64:
				setInt64Value(fieldValue, fieldType, jsonFieldType)
			case reflect.Slice:
				for j := 0; j < fieldValue.Len(); j++ {
					elementValue := fieldValue.Index(j)
					generateFakeData(elementValue.Addr().Interface(), jsonData[fieldType.Name].([]interface{})[0].(map[string]interface{}))
				}
			case reflect.Struct:
				generateFakeData(fieldValue.Addr().Interface(), jsonData[fieldType.Name].(map[string]interface{}))
			}
		}
	}
}

func setStringValue(fieldValue reflect.Value, fieldType reflect.StructField) {
	if fieldGenerator, ok := fieldGenerators[fieldType.Name]; ok {
		fieldValue.SetString(fmt.Sprintf("%v", fieldGenerator()))
	} else {
		fieldValue.SetString(gofakeit.Name())
	}
}

func setFloatValue(fieldValue reflect.Value, fieldType reflect.StructField, jsonFieldType interface{}) {
	if fieldRange, ok := jsonFieldType.(map[string]interface{}); ok {
		min := fieldRange["min"].(float64)
		max := fieldRange["max"].(float64)
		precision := int(fieldRange["precision"].(float64))
		value := gofakeit.Float64Range(min, max)
		reducedValue := reduceFloatPrecision(value, precision)
		fieldValue.SetFloat(reducedValue)
	} else {
		setFloatFieldValue(fieldValue, fieldType)
	}
}

func setFloatFieldValue(fieldValue reflect.Value, fieldType reflect.StructField) {
	if fieldGenerator, ok := fieldGenerators[fieldType.Name]; ok {
		fieldValue.SetFloat(fieldGenerator().(float64))
	} else {
		fieldValue.SetFloat(gofakeit.Float64())
	}
}

func setInt64Value(fieldValue reflect.Value, fieldType reflect.StructField, jsonFieldType interface{}) {
	if fieldRange, ok := jsonFieldType.(map[string]interface{}); ok {
		min := int(fieldRange["min"].(float64))
		max := int(fieldRange["max"].(float64))
		fieldValue.SetInt(int64(gofakeit.IntRange(min, max)))
	} else {
		setInt64FieldValue(fieldValue, fieldType)
	}
}

func setInt64FieldValue(fieldValue reflect.Value, fieldType reflect.StructField) {
	if fieldGenerator, ok := fieldGenerators[fieldType.Name]; ok {
		fieldValue.SetInt(fieldGenerator().(int64))
	} else {
		fieldValue.SetInt(gofakeit.Int64())
	}
}

func reduceFloatPrecision(value float64, precision int) float64 {
	scale := math.Pow10(precision)
	return math.Round(value*scale) / scale
}
