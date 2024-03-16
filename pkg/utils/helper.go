package utils

import "github.com/google/uuid"

// IsIncludedInSlice
func IsIncludedInSlice[T comparable](value T, values []T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

// IsUUID
func IsUUID(s string) bool {
	if _, err := uuid.Parse(s); err != nil {
		return false
	}
	return true
}
