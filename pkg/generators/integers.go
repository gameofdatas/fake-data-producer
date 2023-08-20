package generators

import (
	"errors"
	"reflect"

	"github.com/brianvoe/gofakeit/v6"
)

type IntGenerator struct {
	epochGen DataGenerator
}

func (g IntGenerator) Generate(fieldValue reflect.Value, fieldType reflect.StructField, jsonFieldType interface{}) error {
	if fieldRange, ok := jsonFieldType.(map[string]interface{}); ok && fieldRange["type"].(string) == "epoch" {
		err := g.epochGen.Generate(fieldValue, fieldType, jsonFieldType)
		if err != nil {
			return err
		}
	} else {
		err := setInt64Value(fieldValue, fieldType, jsonFieldType)
		if err != nil {
			return err
		}
	}
	return nil
}

// setInt64Value sets a fake int64 value for the provided fieldValue based on the fieldType and jsonFieldType.
func setInt64Value(fieldValue reflect.Value, fieldType reflect.StructField, jsonFieldType interface{}) error {
	if fieldRange, ok := jsonFieldType.(map[string]interface{}); ok {
		min, okMin := fieldRange["min"].(float64)
		max, okMax := fieldRange["max"].(float64)
		if !okMin {
			return errors.New("missing or invalid 'min' configuration for float generator")
		}
		if !okMax {
			return errors.New("missing or invalid 'max' configuration for float generator")
		}
		if min > max {
			return errors.New("'min' value cannot be greater than 'max' value in int generator")
		}
		fieldValue.SetInt(int64(gofakeit.IntRange(int(min), int(max))))
		return nil
	}

	return setInt64FieldValue(fieldValue, fieldType)
}

// setInt64FieldValue sets a fake int64 value for the provided fieldValue based on the fieldType.
func setInt64FieldValue(fieldValue reflect.Value, fieldType reflect.StructField) error {
	if fieldGenerator, ok := FieldGenerators[fieldType.Name]; ok {
		fieldValue.SetInt(fieldGenerator().(int64))
	} else {
		fieldValue.SetInt(gofakeit.Int64())
	}
	return nil
}

func NewIntegerGenerator() DataGenerator {
	return &IntGenerator{
		epochGen: NewEpochGenerator(),
	}
}
