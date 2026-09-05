package utils

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// AnyToString coerces an arbitrary NITRO GET response value to a string,
// falling back to fmt.Sprintf for non-string values so callers never panic on a
// type assertion.
func AnyToString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}

// MapGetString / MapGetInt64 / MapGetBool / MapGetStringList project a value out
// of a NITRO GET response map into a Terraform Plugin Framework type, returning
// the typed Null when the key is absent/nil (or the element type does not match).
//
// These are intended for DATA SOURCE getters (and any pure-Computed read-back),
// where attributes start Null and are simply filled from the GET. They do NOT
// implement the resource "resolve unknown -> null, but preserve a configured
// value the GET omits" idiom, so do not use them for a resource's
// Optional+Computed echo attributes.
func MapGetString(data map[string]interface{}, key string) types.String {
	if v, ok := data[key]; ok && v != nil {
		return types.StringValue(AnyToString(v))
	}
	return types.StringNull()
}

func MapGetInt64(data map[string]interface{}, key string) types.Int64 {
	if v, ok := data[key]; ok && v != nil {
		if iv, err := ConvertToInt64(v); err == nil {
			return types.Int64Value(iv)
		}
	}
	return types.Int64Null()
}

func MapGetBool(data map[string]interface{}, key string) types.Bool {
	if v, ok := data[key]; ok && v != nil {
		if b, ok := v.(bool); ok {
			return types.BoolValue(b)
		}
	}
	return types.BoolNull()
}

func MapGetStringList(data map[string]interface{}, key string) types.List {
	if v, ok := data[key]; ok && v != nil {
		if arr, ok := v.([]interface{}); ok {
			elems := make([]attr.Value, 0, len(arr))
			for _, e := range arr {
				elems = append(elems, types.StringValue(AnyToString(e)))
			}
			if lv, d := types.ListValue(types.StringType, elems); !d.HasError() {
				return lv
			}
		}
	}
	return types.ListNull(types.StringType)
}
