package utils

import (
	"bytes"
	"encoding/json"
)

func PrettyJSON(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(v); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}
