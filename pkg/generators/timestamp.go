package generators

import (
	"errors"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

type TimestampsGenerator struct{}

func (g TimestampsGenerator) Generate(fieldValue reflect.Value, fieldType reflect.StructField, jsonFieldType interface{}) error {
	config, ok := jsonFieldType.(map[string]interface{})
	if !ok {
		return errors.New("invalid configuration for timestamps generator")
	}

	format, formatOK := config["format"].(string)
	zone, zoneOK := config["zone"].(string)
	if !formatOK {
		return errors.New("missing or invalid 'format' configuration for timestamps generator")
	}
	if !zoneOK {
		zone = "Local"
	}

	fakeTime, err := generateFakeTimestamp(format, zone)
	if err != nil {
		return err
	}
	fieldValue.SetString(fakeTime)
	return nil
}

func generateFakeTimestamp(format, zone string) (string, error) {
	// Get the current time to use as the base for generating the fake timestamp
	var now time.Time
	if zone != "" {
		loc, err := time.LoadLocation(zone)
		if err == nil {
			now = time.Now().In(loc)
		} else {
			now = time.Now()
		}
	} else {
		now = time.Now()
	}

	// Generate a random number of seconds to add or subtract from the current time
	// to create a fake timestamp within a certain range from the current time.
	rand.Seed(time.Now().UnixNano())
	randomSeconds := rand.Intn(3600) // Generating a random number of seconds between 0 and 3600 (1 hour)

	// Add or subtract the random number of seconds to/from the current time
	fakeTime := now.Add(time.Second * time.Duration(randomSeconds))
	return fakeTime.Format(parseGolangTimeFormat(format)), nil
}

func parseGolangTimeFormat(format string) string {
	replacer := strings.NewReplacer(
		"yyyy", "2006",
		"MM", "01",
		"dd", "02",
		"HH", "15",
		"mm", "04",
		"ss", "05",
	)
	return replacer.Replace(format)
}

func NewTimestampGenerator() DataGenerator {
	return &TimestampsGenerator{}
}
