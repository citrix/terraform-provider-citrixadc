package utils

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UnsetOnRemoveOrKeepDefaultString is the Framework analogue of an SDK v2
// DiffSuppressFunc paired with unset-on-remove support. It replaces a schema
// Default for an attribute that (a) does not exist on older firmware (so a Default
// would be force-sent and rejected with errorcode 278 on a build that lacks it) and
// (b) is always echoed on GET by firmware that DOES support it (so a naive
// "mark unknown on removal" modifier would churn on every plan, because the echoed
// value keeps prior state non-null).
//
// When the attribute is absent from config:
//   - Create (no prior state): leave it computed from the appliance GET.
//   - Steady state, prior value is null OR equals the NITRO default: keep prior state
//     (suppress the spurious "-> (known after apply)" diff, like UseStateForUnknown).
//   - Steady state, prior value is a previously-set non-default value: mark the plan
//     unknown so Update detects the removal and issues ?action=unset. The value
//     resolves from the post-unset GET, so the plan stays consistent.
//
// Removing a value that already equals the default is a no-op either way (the
// appliance is already at the default), so not unsetting it is harmless.
//
// Verified on NS 14.1 (no churn, unset works) and NS 13.1 (create without the
// attribute no longer force-sends it -> no ec278). See firmware_131_compat_review.md.
type UnsetOnRemoveOrKeepDefaultString struct{ DefaultValue string }

func (m UnsetOnRemoveOrKeepDefaultString) Description(_ context.Context) string {
	return "Suppresses the spurious diff for an unconfigured, always-echoed attribute while still unsetting a previously-set non-default value when it is removed from config."
}

func (m UnsetOnRemoveOrKeepDefaultString) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m UnsetOnRemoveOrKeepDefaultString) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.ConfigValue.IsNull() {
		return // user is setting it -> use the configured value
	}
	if req.State.Raw.IsNull() {
		return // create -> let it be computed from the appliance
	}
	if req.StateValue.IsNull() || req.StateValue.ValueString() == m.DefaultValue {
		resp.PlanValue = req.StateValue // never user-set / already default -> no spurious diff
		return
	}
	resp.PlanValue = types.StringUnknown() // non-default value removed -> unset on apply
}

// UnsetOnRemoveOrKeepDefaultInt64 is the types.Int64 counterpart of
// UnsetOnRemoveOrKeepDefaultString. See that type for the full rationale.
type UnsetOnRemoveOrKeepDefaultInt64 struct{ DefaultValue int64 }

func (m UnsetOnRemoveOrKeepDefaultInt64) Description(_ context.Context) string {
	return "Suppresses the spurious diff for an unconfigured, always-echoed attribute while still unsetting a previously-set non-default value when it is removed from config."
}

func (m UnsetOnRemoveOrKeepDefaultInt64) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m UnsetOnRemoveOrKeepDefaultInt64) PlanModifyInt64(_ context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if !req.ConfigValue.IsNull() {
		return
	}
	if req.State.Raw.IsNull() {
		return
	}
	if req.StateValue.IsNull() || req.StateValue.ValueInt64() == m.DefaultValue {
		resp.PlanValue = req.StateValue
		return
	}
	resp.PlanValue = types.Int64Unknown()
}

// UnsetOnRemoveOrKeepDefaultBool is the types.Bool counterpart of
// UnsetOnRemoveOrKeepDefaultString. See that type for the full rationale.
type UnsetOnRemoveOrKeepDefaultBool struct{ DefaultValue bool }

func (m UnsetOnRemoveOrKeepDefaultBool) Description(_ context.Context) string {
	return "Suppresses the spurious diff for an unconfigured, always-echoed attribute while still unsetting a previously-set non-default value when it is removed from config."
}

func (m UnsetOnRemoveOrKeepDefaultBool) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m UnsetOnRemoveOrKeepDefaultBool) PlanModifyBool(_ context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}
	if req.State.Raw.IsNull() {
		return
	}
	if req.StateValue.IsNull() || req.StateValue.ValueBool() == m.DefaultValue {
		resp.PlanValue = req.StateValue
		return
	}
	resp.PlanValue = types.BoolUnknown()
}
