package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {

	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("File not found")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0, errors.New("Invalid value")
	}

	return value, nil
}

func WriteFloatToFile(value *float64, fileName string) {
	os.WriteFile(
		fileName,
		[]byte(fmt.Sprint(*value)),
		0644,
	)
}
