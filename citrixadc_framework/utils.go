package citrixadc_framework

// intPtr returns a pointer to the provided int value
// This is useful for optional fields in structs that require *int
func intPtr(i int) *int {
	return &i
}
