package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCreateGenericData(t *testing.T) {
	testCases := []struct {
		name     string
		config   string
		expected string
	}{
		{
			name:     "Test case 1",
			config:   `{"Age":{"type":"int","min":20,"max":20}}`,
			expected: `{"Age":20}`,
		},
		{
			name:     "Test case 2",
			config:   `{"UnitPrice":{"type":"float","min":2.23,"max":2.23,"precision":1}}`,
			expected: `{"UnitPrice":2.2}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			configData := make(map[string]interface{})
			err := json.Unmarshal([]byte(tc.config), &configData)
			if err != nil {
				t.Fatalf("Failed to unmarshal config: %v", err)
			}

			dataStruct := CreateStructFromJSON(configData)

			result, err := CreateGenericData(dataStruct, configData)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			value, err := json.Marshal(result)
			if err != nil {
				t.Errorf("Could not marshal results: %v", err)
			}

			if !reflect.DeepEqual(string(value), tc.expected) {
				t.Errorf("Expected: %s, but got: %s", tc.expected, value)
			}
		})
	}
}
