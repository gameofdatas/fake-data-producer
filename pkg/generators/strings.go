package generators

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/brianvoe/gofakeit"
)

type StringGenerator struct {
	timestampGen DataGenerator
}

func (g StringGenerator) Generate(fieldValue reflect.Value, fieldType reflect.StructField, jsonFieldType interface{}) error {
	if fieldRange, ok := jsonFieldType.(map[string]interface{}); ok {
		if fieldRange["type"].(string) == "timestamp" {
			err := g.timestampGen.Generate(fieldValue, fieldType, jsonFieldType)
			if err != nil {
				return err
			}
		}
		if isID, hasIsID := fieldRange["isId"].(bool); hasIsID && isID {
			fieldValue.SetString(gofakeit.UUID())
		}
		if options, hasOptions := fieldRange["options"].([]interface{}); hasOptions {
			// Pick a random value from the options
			if len(options) > 0 {
				index := rand.Intn(len(options))
				fieldValue.SetString(fmt.Sprintf("%v", options[index]))
			}
		}

	} else {
		fieldGenerator, exists := FieldGenerators[fieldType.Name]
		if exists {
			fieldValue.SetString(fmt.Sprintf("%v", fieldGenerator()))
		} else {
			fieldValue.SetString(gofakeit.Name())
		}
	}
	return nil
}

func NewStringGenerator() DataGenerator {
	return &StringGenerator{
		timestampGen: NewTimestampGenerator(),
	}
}
