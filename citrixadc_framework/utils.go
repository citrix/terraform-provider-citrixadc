package citrixadc_framework

// intPtr returns a pointer to the provided int value
// This is useful for optional fields in structs that require *int
func intPtr(i int) *int {
	return &i
}

// boolPtr returns a pointer to the provided bool value
// This is useful for optional fields in structs that require *bool
func boolPtr(b bool) *bool {
	return &b
}
