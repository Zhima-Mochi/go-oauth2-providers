package oauth2providers

import (
	"encoding/json"
	"io"
)

// parseJSON parses JSON from []byte
func parseJSON(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// parseJSONFromReader parses JSON from io.Reader
func parseJSONFromReader(input io.Reader) (map[string]interface{}, error) {
	var result map[string]interface{}
	decoder := json.NewDecoder(input)
	err := decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
