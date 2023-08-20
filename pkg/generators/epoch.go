package generators

import (
	"reflect"
	"time"
)

type EpochGenerator struct{}

func (g EpochGenerator) Generate(fieldValue reflect.Value, fieldType reflect.StructField, jsonFieldType interface{}) error {
	if fieldRange, ok := jsonFieldType.(map[string]interface{}); ok {
		unit, unitOK := fieldRange["unit"].(string)
		zoneValue, zoneExists := fieldRange["zone"].(string)

		if !unitOK {
			unit = ""
		}
		if !zoneExists {
			zoneValue = "Local"
		}

		// Generate the fake epoch time
		fakeTime, err := generateFakeEpochTime(unit, zoneValue)
		if err != nil {
			return err
		}

		fieldValue.SetInt(fakeTime)
	}
	return nil
}

func generateFakeEpochTime(unit string, zone string) (int64, error) {
	// Get the current time in the specified time zone
	var now time.Time
	if zone != "" {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			return 0, err
		}
		now = time.Now().In(loc)
	} else {
		now = time.Now()
	}

	var fakeTime int64
	switch unit {
	case "micro":
		fakeTime = now.UnixMicro()
	case "nano":
		fakeTime = now.UnixNano()
	case "millis":
		fakeTime = now.UnixMilli()
	default:
		fakeTime = now.Unix()
	}

	return fakeTime, nil
}

func NewEpochGenerator() DataGenerator {
	return &EpochGenerator{}
}
