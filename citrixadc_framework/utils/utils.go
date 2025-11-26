package utils

import (
	"fmt"
	"strconv"
)

// intPtr returns a pointer to the provided int value
// This is useful for optional fields in structs that require *int
func IntPtr(i int) *int {
	return &i
}

// boolPtr returns a pointer to the provided bool value
// This is useful for optional fields in structs that require *bool
func BoolPtr(b bool) *bool {
	return &b
}

// Helper function to convert interface{} to int64
func ConvertToInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("cannot convert %T to int64", value)
	}
}
