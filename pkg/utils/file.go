package utils

import (
	"fmt"
	"os"
)

func WritePrettyJSONFile(filename string, data interface{}) error {
	outData, err := PrettyJSON(data)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	if err = os.WriteFile(filename, outData, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	fmt.Println("Comparison saved to " + filename)
	return nil
}
