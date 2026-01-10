package utils

import (
	"strings"
)

func ValidateRequiredFields(data map[string]string) (missing []string) {
	for key, value := range data {
		if strings.TrimSpace(value) == "" {
			missing = append(missing, key)
		}
	}
	return
}
