package generators

import (
	"errors"
	"math"
	"reflect"

	"github.com/brianvoe/gofakeit"
)

type FloatGenerator struct{}

func (g FloatGenerator) Generate(fieldValue reflect.Value, fieldType reflect.StructField, jsonFieldType interface{}) error {
	if fieldRange, ok := jsonFieldType.(map[string]interface{}); ok {
		min, minOK := fieldRange["min"].(float64)
		max, maxOK := fieldRange["max"].(float64)
		precision, precisionOK := fieldRange["precision"].(float64)

		if !minOK {
			return errors.New("missing or invalid 'min' configuration for float generator")
		}
		if !maxOK {
			return errors.New("missing or invalid 'max' configuration for float generator")
		}
		if !precisionOK {
			return errors.New("missing or invalid 'precision' configuration for float generator")
		}
		if min > max {
			return errors.New("'min' value cannot be greater than 'max' value in float generator")
		}

		value := gofakeit.Float64Range(min, max)
		reducedValue := reduceFloatPrecision(value, int(precision))
		fieldValue.SetFloat(reducedValue)
	} else {
		if err := setFloatFieldValue(fieldValue, fieldType); err != nil {
			return err
		}
	}
	return nil
}

func setFloatFieldValue(fieldValue reflect.Value, fieldType reflect.StructField) error {
	if fieldGenerator, ok := FieldGenerators[fieldType.Name]; ok {
		fieldValue.SetFloat(fieldGenerator().(float64))
	} else {
		fieldValue.SetFloat(gofakeit.Float64())
	}
	return nil
}

// reduceFloatPrecision reduces the precision of a floating-point value to the specified number of decimal places.
func reduceFloatPrecision(value float64, precision int) float64 {
	scale := math.Pow10(precision)
	return math.Round(value*scale) / scale
}

func NewFloatGenerator() DataGenerator {
	return &FloatGenerator{}
}
