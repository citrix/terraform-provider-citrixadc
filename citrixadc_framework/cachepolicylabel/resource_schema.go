package cachepolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CachepolicylabelResourceModel describes the resource data model.
type CachepolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Evaluates types.String `tfsdk:"evaluates"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
}

func (r *CachepolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cachepolicylabel resource.",
			},
			"evaluates": schema.StringAttribute{
				// SDK v2 had `evaluates` as Optional + ForceNew (NOT Required). Keep it
				// Optional for backward compatibility (a config that omits it must still
				// be accepted at plan-time). It is not updateable, so it RequiresReplace.
				// Not Computed: the value is user-supplied and always echoed back by GET,
				// so an Optional-only attribute round-trips cleanly without a perpetual diff.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "When to evaluate policies bound to this label: request-time or response-time.",
			},
			"labelname": schema.StringAttribute{
				// SDK v2 primary key: Required + ForceNew -> Required + RequiresReplace.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the label is created.",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it must
				// NOT force replacement - it drives an in-place rename via Update. It is a
				// pure user input, never echoed back by GET, so it is neither Computed nor
				// RequiresReplace.
				Optional:    true,
				Description: "New name for the cache-policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
		},
	}
}

func cachepolicylabelGetThePayloadFromthePlan(ctx context.Context, data *CachepolicylabelResourceModel) cache.Cachepolicylabel {
	tflog.Debug(ctx, "In cachepolicylabelGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cachepolicylabel := cache.Cachepolicylabel{}
	if !data.Evaluates.IsNull() && !data.Evaluates.IsUnknown() {
		cachepolicylabel.Evaluates = data.Evaluates.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		cachepolicylabel.Labelname = data.Labelname.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of the
	// add payload, so it is deliberately excluded from the create POST body.

	return cachepolicylabel
}

func cachepolicylabelSetAttrFromGet(ctx context.Context, data *CachepolicylabelResourceModel, getResponseData map[string]interface{}) *CachepolicylabelResourceModel {
	tflog.Debug(ctx, "In cachepolicylabelSetAttrFromGet Function")

	// Convert API response to model.
	// evaluates is always returned by GET; only overwrite when present so a config
	// value is never clobbered by an absent field.
	if val, ok := getResponseData["evaluates"]; ok && val != nil {
		data.Evaluates = types.StringValue(val.(string))
	}
	// labelname is the user-facing key. Once a rename has happened (via newname), the
	// live object name (tracked by data.Id) diverges from the configured labelname,
	// and GET returns the live (new) name. Overwriting labelname from GET would clobber
	// the user's configured value and trigger a spurious RequiresReplace diff. So only
	// adopt the GET value when we don't already have one (e.g. on import, where state
	// carries only the ID); otherwise preserve.
	if data.Labelname.IsNull() || data.Labelname.IsUnknown() || data.Labelname.ValueString() == "" {
		if val, ok := getResponseData["labelname"]; ok && val != nil {
			data.Labelname = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.

	return data
}

// cachepolicylabelSetAttrFromGetForDatasource faithfully copies every field from the
// GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func cachepolicylabelSetAttrFromGetForDatasource(ctx context.Context, data *CachepolicylabelResourceModel, getResponseData map[string]interface{}) *CachepolicylabelResourceModel {
	tflog.Debug(ctx, "In cachepolicylabelSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["evaluates"]; ok && val != nil {
		data.Evaluates = types.StringValue(val.(string))
	} else {
		data.Evaluates = types.StringNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	return data
}
