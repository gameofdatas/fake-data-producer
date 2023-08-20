package fakedata

import (
	"reflect"

	"fake-data-producer/pkg/generators"

	"github.com/brianvoe/gofakeit/v6"
)

// CreateGenericData generates fake data for the provided dataStruct based on the configuration data (confData).
// It returns the modified dataStruct and any error encountered.
func CreateGenericData(dataStruct interface{}, confData map[string]interface{}) (interface{}, error) {
	err := generateFakeData(dataStruct, confData)
	if err != nil {
		return nil, err
	}
	return dataStruct, nil
}

// generateFakeData generates fake data for the provided dataStruct based on the provided jsonData configuration.
// It recursively populates the fields of the struct with appropriate fake data.
func generateFakeData(dataStruct interface{}, jsonData map[string]interface{}) error {
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

			fakeGenerator := generators.NewDataGenerator(fieldValue.Kind().String())
			// Determine the kind of the field and set the appropriate fake value
			switch fieldValue.Kind() {
			case reflect.String:
				err := fakeGenerator.Generate(fieldValue, fieldType, jsonFieldType)
				if err != nil {
					return err
				}
			case reflect.Float64:
				err := fakeGenerator.Generate(fieldValue, fieldType, jsonFieldType)
				if err != nil {
					return err
				}
			case reflect.Bool:
				fieldValue.SetBool(gofakeit.Bool())
			case reflect.Int64:
				err := fakeGenerator.Generate(fieldValue, fieldType, jsonFieldType)
				if err != nil {
					return err
				}
			case reflect.Slice:
				// Handle slice fields by iterating over each element and recursively generating fake data
				for j := 0; j < fieldValue.Len(); j++ {
					elementValue := fieldValue.Index(j)
					err := generateFakeData(elementValue.Addr().Interface(), jsonData[fieldType.Name].([]interface{})[0].(map[string]interface{}))
					if err != nil {
						return err
					}
				}
			case reflect.Struct:
				err := generateFakeData(fieldValue.Addr().Interface(), jsonData[fieldType.Name].(map[string]interface{}))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
