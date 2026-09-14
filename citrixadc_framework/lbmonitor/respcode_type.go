package lbmonitor

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// respcode is a list of HTTP response codes for an lbmonitor. NetScaler stores it
// canonically: any run of consecutive codes is compressed into a "start-end" range
// (301,302 -> "301-302"; 301,302,303 -> "301-303"), a degenerate "n-n" collapses to
// "n", and non-adjacent codes stay separate. So a configured [200,301,302,401] comes
// back from GET as ["200","301-302","401"] — a semantically identical but differently
// represented value. Without semantic equality the Plugin Framework rejects this with
// "Provider produced inconsistent result after apply" (GH #1462); the legacy SDK v2
// provider only avoided the error because Terraform core relaxes that check for the
// legacy type system, so it applied (with a warning/perpetual diff).
//
// respcodeListType / respcodeListValue teach the framework that two respcode lists are
// equal when they expand to the same SET of integer codes, so any equivalent
// representation ([200,301,302,401], ["200","301-302","401"], reordered, …) applies
// cleanly and is idempotent.

var (
	_ basetypes.ListTypable                    = respcodeListType{}
	_ basetypes.ListValuable                   = respcodeListValue{}
	_ basetypes.ListValuableWithSemanticEquals = respcodeListValue{}
)

// respcodeListType is the custom attr.Type for the respcode list attribute.
type respcodeListType struct {
	basetypes.ListType
}

// newRespcodeListType returns the respcode list type over string elements.
func newRespcodeListType() respcodeListType {
	return respcodeListType{basetypes.ListType{ElemType: types.StringType}}
}

func (t respcodeListType) Equal(o attr.Type) bool {
	other, ok := o.(respcodeListType)
	if !ok {
		return false
	}
	return t.ListType.Equal(other.ListType)
}

func (t respcodeListType) String() string {
	return "lbmonitor.respcodeListType"
}

func (t respcodeListType) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	return respcodeListValue{ListValue: in}, nil
}

func (t respcodeListType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	listValue, ok := attrValue.(basetypes.ListValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T from ListType.ValueFromTerraform", attrValue)
	}
	return respcodeListValue{ListValue: listValue}, nil
}

func (t respcodeListType) ValueType(ctx context.Context) attr.Value {
	return respcodeListValue{ListValue: t.ListType.ValueType(ctx).(basetypes.ListValue)}
}

// respcodeListValue is the custom attr.Value for the respcode list attribute.
type respcodeListValue struct {
	basetypes.ListValue
}

func (v respcodeListValue) Type(ctx context.Context) attr.Type {
	return newRespcodeListType()
}

// Equal is strict (element-by-element) equality; semantic (set) equality lives in
// ListSemanticEquals so standard comparisons (e.g. Update change-detection) stay exact.
func (v respcodeListValue) Equal(o attr.Value) bool {
	other, ok := o.(respcodeListValue)
	if !ok {
		return false
	}
	return v.ListValue.Equal(other.ListValue)
}

// ListSemanticEquals reports whether two respcode lists denote the same set of response
// codes after expanding NetScaler "start-end" ranges — e.g. an applied
// ["200","301-302","401"] equals a planned ["200","301","302","401"]. This prevents the
// "inconsistent result after apply" error and any perpetual diff (GH #1462). The
// framework only invokes this for two known, non-null values.
func (v respcodeListValue) ListSemanticEquals(ctx context.Context, newValuable basetypes.ListValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(respcodeListValue)
	if !ok {
		return false, diags
	}
	if v.IsNull() || v.IsUnknown() || newValue.IsNull() || newValue.IsUnknown() {
		return false, diags
	}

	priorSet, okPrior, d := expandRespcodes(ctx, v.ListValue)
	diags.Append(d...)
	newSet, okNew, d := expandRespcodes(ctx, newValue.ListValue)
	diags.Append(d...)
	if diags.HasError() || !okPrior || !okNew {
		// An element we cannot parse as a code/range -> fall back to strict equality.
		return false, diags
	}

	if len(priorSet) != len(newSet) {
		return false, diags
	}
	for code := range priorSet {
		if !newSet[code] {
			return false, diags
		}
	}
	return true, diags
}

// expandRespcodes turns a respcode list into the set of individual integer codes it
// denotes, expanding "start-end" ranges. ok is false if any element is not a plain
// integer or a numeric range, so the caller can fall back to strict equality.
func expandRespcodes(ctx context.Context, l basetypes.ListValue) (map[int]bool, bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	var elems []string
	diags.Append(l.ElementsAs(ctx, &elems, false)...)
	if diags.HasError() {
		return nil, false, diags
	}
	set := map[int]bool{}
	for _, e := range elems {
		lo, hi, ok := parseRespcodeRange(strings.TrimSpace(e))
		if !ok {
			return nil, false, diags
		}
		for c := lo; c <= hi; c++ {
			set[c] = true
		}
	}
	return set, true, diags
}

// parseRespcodeRange parses "n" -> (n,n) and "a-b" -> (a,b) for non-negative ints (a<=b).
func parseRespcodeRange(s string) (int, int, bool) {
	if n, err := strconv.Atoi(s); err == nil && n >= 0 {
		return n, n, true
	}
	if i := strings.IndexByte(s, '-'); i > 0 {
		lo, err1 := strconv.Atoi(strings.TrimSpace(s[:i]))
		hi, err2 := strconv.Atoi(strings.TrimSpace(s[i+1:]))
		if err1 == nil && err2 == nil && lo >= 0 && lo <= hi {
			return lo, hi, true
		}
	}
	return 0, 0, false
}
