package generators

import (
	"reflect"
)

type DataGenerator interface {
	Generate(fieldValue reflect.Value, fieldType reflect.StructField, jsonFieldType interface{}) error
}

func NewDataGenerator(dataType string) DataGenerator {
	switch dataType {
	case "string":
		return NewStringGenerator()
	case "float64":
		return NewFloatGenerator()
	case "int64":
		return NewIntegerGenerator()
	}
	return nil
}
