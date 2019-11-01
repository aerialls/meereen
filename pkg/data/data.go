package data

import "fmt"

// GetParameter returns a parameter from the data. If the parameter is not here, the
// fallback value is returned
func GetParameter(data map[string]string, parameter string, fallback string) string {
	value, ok := data[parameter]
	if !ok {
		return fallback
	}

	return value
}

// GetRequiredParameter returns a required parameter from the data. If the parameter is not
// here, an error is returned
func GetRequiredParameter(data map[string]string, parameter string) (string, error) {
	value, ok := data[parameter]
	if !ok {
		return "", fmt.Errorf("required parameter %s is missing", parameter)
	}

	return value, nil
}
