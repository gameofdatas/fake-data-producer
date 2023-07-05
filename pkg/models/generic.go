package models

import (
	"fmt"
	"math"
	"reflect"

	"github.com/brianvoe/gofakeit/v6"
)

// CreateGenericData generates fake data for the provided dataStruct based on the configuration data (confData).
// It returns the modified dataStruct and any error encountered.
func CreateGenericData(dataStruct interface{}, confData map[string]interface{}) (interface{}, error) {
	generateFakeData(dataStruct, confData)
	return dataStruct, nil
}

// generateFakeData generates fake data for the provided dataStruct based on the provided jsonData configuration.
// It recursively populates the fields of the struct with appropriate fake data.
func generateFakeData(dataStruct interface{}, jsonData map[string]interface{}) {
	// Set the seed for random data generation
	gofakeit.Seed(0)

	val := reflect.ValueOf(dataStruct)
	if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		val = val.Elem()

		// Iterate over the fields of the struct
		for i := 0; i < val.NumField(); i++ {
			fieldValue := val.Field(i)
			fieldType := val.Type().Field(i)
			jsonFieldType := jsonData[fieldType.Name]

			// Determine the kind of the field and set the appropriate fake value
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
				// Handle slice fields by iterating over each element and recursively generating fake data
				for j := 0; j < fieldValue.Len(); j++ {
					elementValue := fieldValue.Index(j)
					generateFakeData(elementValue.Addr().Interface(), jsonData[fieldType.Name].([]interface{})[0].(map[string]interface{}))
				}
			case reflect.Struct:
				// Handle nested struct fields by recursively generating fake data
				generateFakeData(fieldValue.Addr().Interface(), jsonData[fieldType.Name].(map[string]interface{}))
			}
		}
	}
}

// setStringValue sets a fake string value for the provided fieldValue based on the fieldType.
func setStringValue(fieldValue reflect.Value, fieldType reflect.StructField) {
	if fieldGenerator, ok := fieldGenerators[fieldType.Name]; ok {
		fieldValue.SetString(fmt.Sprintf("%v", fieldGenerator()))
	} else {
		fieldValue.SetString(gofakeit.Name())
	}
}

// setFloatValue sets a fake float value for the provided fieldValue based on the fieldType and jsonFieldType.
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

// setFloatFieldValue sets a fake float value for the provided fieldValue based on the fieldType.
func setFloatFieldValue(fieldValue reflect.Value, fieldType reflect.StructField) {
	if fieldGenerator, ok := fieldGenerators[fieldType.Name]; ok {
		fieldValue.SetFloat(fieldGenerator().(float64))
	} else {
		fieldValue.SetFloat(gofakeit.Float64())
	}
}

// setInt64Value sets a fake int64 value for the provided fieldValue based on the fieldType and jsonFieldType.
func setInt64Value(fieldValue reflect.Value, fieldType reflect.StructField, jsonFieldType interface{}) {
	if fieldRange, ok := jsonFieldType.(map[string]interface{}); ok {
		min := int(fieldRange["min"].(float64))
		max := int(fieldRange["max"].(float64))
		fieldValue.SetInt(int64(gofakeit.IntRange(min, max)))
	} else {
		setInt64FieldValue(fieldValue, fieldType)
	}
}

// setInt64FieldValue sets a fake int64 value for the provided fieldValue based on the fieldType.
func setInt64FieldValue(fieldValue reflect.Value, fieldType reflect.StructField) {
	if fieldGenerator, ok := fieldGenerators[fieldType.Name]; ok {
		fieldValue.SetInt(fieldGenerator().(int64))
	} else {
		fieldValue.SetInt(gofakeit.Int64())
	}
}

// reduceFloatPrecision reduces the precision of a floating-point value to the specified number of decimal places.
func reduceFloatPrecision(value float64, precision int) float64 {
	scale := math.Pow10(precision)
	return math.Round(value*scale) / scale
}
