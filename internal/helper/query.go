package helper

import "strconv"

func QueryInt(query map[string]string, key string, fallback int) (int, error) {
	value, exists := query[key]
	if !exists {
		return fallback, nil
	}
	return strconv.Atoi(value)
}
