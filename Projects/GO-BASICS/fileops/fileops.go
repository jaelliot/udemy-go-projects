package fileops

import (
	"errors"  // for error handling
	"fmt"     // for input/output
	"os"      // for file operations
	"strconv" // for string conversion
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 10000, errors.New("error reading file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
