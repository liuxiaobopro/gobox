package any

import (
	"encoding/json"
	"testing"
)

func TestAnyToStruct(t *testing.T) {
	type outputType struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Other struct {
			Class int `json:"class"`
		} `json:"other"`
	}
	var output outputType
	input := map[string]interface{}{
		"name": "liuxiaobopro",
		"age":  18,
		"other": map[string]interface{}{
			"class": 1,
		},
	}
	if err := AnyToStruct(input, &output); err != nil {
		t.Errorf("AnyToStruct() error = %v", err)
	}

	t.Logf("output: %v", output)
	jsonData, _ := json.Marshal(output)
	t.Logf("jsonData: %v", string(jsonData))
}
