package rewritepolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/rewrite"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RewritepolicylabelResourceModel describes the resource data model.
type RewritepolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
	Transform types.String `tfsdk:"transform"`
}

func (r *RewritepolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rewritepolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				// SDK v2 had comment as Optional+Computed+ForceNew, and it stored comment=""
				// when unset. Reproduce that contract as Optional+Computed so that upgraded
				// SDK v2 state (comment="") does NOT read as a change against a config-null
				// value on a RequiresReplace attribute -- which previously forced a spurious
				// destroy+recreate of the label on upgrade, dropping any child *_binding and
				// failing with ec315 "Resource in use".
				//
				// UseStateForUnknown is listed BEFORE RequiresReplace: when comment is unset,
				// the Computed plan value resolves to the prior state value ("") instead of
				// staying unknown, so RequiresReplace compares ""=="" and does not trigger.
				//
				// NITRO omits "comment" from GET when empty; rewritepolicylabelSetAttrFromGet
				// therefore sets it to "" on absence so the Computed value stays KNOWN after a
				// fresh create (avoiding "inconsistent result after apply / still indicated an
				// unknown value"). ForceNew is preserved via RequiresReplace.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this rewrite policy label.",
			},
			"labelname": schema.StringAttribute{
				// SDK v2: Required + ForceNew.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the rewrite policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rewrite policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite policy label\" or 'my rewrite policy label').",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it must
				// NOT force replacement - it drives an in-place rename via Update. Not
				// Computed: it is a pure user input, never echoed back by GET. SDK v2 did
				// not expose rename, so existing configs never set newname; this is purely
				// additive and non-breaking.
				Optional:    true,
				Description: "New name for the rewrite policy label. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy label\" or 'my policy label').",
			},
			"transform": schema.StringAttribute{
				// SDK v2: Optional+Computed+ForceNew. NITRO metadata marks transform as
				// required, but the backward-compatible contract is the SDK v2 schema, so
				// keep it Optional+Computed. transform is always returned by GET, so there
				// is no unknown-after-apply trap. ForceNew preserved via RequiresReplace.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Types of transformations allowed by the policies bound to the label. For Rewrite, the following types are supported:\n* http_req - HTTP requests\n* http_res - HTTP responses\n* othertcp_req - Non-HTTP TCP requests\n* othertcp_res - Non-HTTP TCP responses\n* url - URLs\n* text - Text strings\n* clientless_vpn_req - Citrix ADC clientless VPN requests\n* clientless_vpn_res - Citrix ADC clientless VPN responses\n* sipudp_req - SIP requests\n* sipudp_res - SIP responses\n* diameter_req - DIAMETER requests\n* diameter_res - DIAMETER responses\n* radius_req - RADIUS requests\n* radius_res - RADIUS responses\n* dns_req - DNS requests\n* dns_res - DNS responses\n* mqtt_req - MQTT requests\n* mqtt_res - MQTT responses",
			},
		},
	}
}

func rewritepolicylabelGetThePayloadFromthePlan(ctx context.Context, data *RewritepolicylabelResourceModel) rewrite.Rewritepolicylabel {
	tflog.Debug(ctx, "In rewritepolicylabelGetThePayloadFromthePlan Function")

	// Create API request body from the model
	rewritepolicylabel := rewrite.Rewritepolicylabel{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		rewritepolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		rewritepolicylabel.Labelname = data.Labelname.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of the
	// add payload, so it is deliberately excluded from the create POST body.
	if !data.Transform.IsNull() && !data.Transform.IsUnknown() {
		rewritepolicylabel.Transform = data.Transform.ValueString()
	}

	return rewritepolicylabel
}

func rewritepolicylabelSetAttrFromGet(ctx context.Context, data *RewritepolicylabelResourceModel, getResponseData map[string]interface{}) *RewritepolicylabelResourceModel {
	tflog.Debug(ctx, "In rewritepolicylabelSetAttrFromGet Function")

	// Convert API response to model.
	// comment is Optional+Computed (UseStateForUnknown+RequiresReplace) mirroring the SDK v2
	// Optional+Computed+ForceNew contract, which stored comment="" when unset. NITRO omits
	// "comment" from the GET response when it is empty, so set the model to "" on absence:
	// this keeps the Computed value KNOWN after a fresh create (avoiding "inconsistent result
	// after apply / still indicated an unknown value") and matches the SDK v2 stored "". A
	// non-empty comment is always echoed by GET, so a user-set value is never clobbered.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringValue("")
	}
	// labelname is the user-facing key. Once a rename has happened (via newname), the
	// live object name (tracked by data.Id) diverges from the configured labelname, and
	// GET returns the live (new) name. Overwriting labelname from GET would clobber the
	// user's configured value and trigger a spurious RequiresReplace diff. So only adopt
	// the GET value when we don't already have one (e.g. on import, where state carries
	// only the ID); otherwise preserve.
	if data.Labelname.IsNull() || data.Labelname.IsUnknown() || data.Labelname.ValueString() == "" {
		if val, ok := getResponseData["labelname"]; ok && val != nil {
			data.Labelname = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["transform"]; ok && val != nil {
		data.Transform = types.StringValue(val.(string))
	}

	return data
}

// rewritepolicylabelSetAttrFromGetForDatasource faithfully copies every field from the
// GET response. The datasource has no prior plan/state to preserve, so it must populate
// the model directly from the API response and set the ID itself.
func rewritepolicylabelSetAttrFromGetForDatasource(ctx context.Context, data *RewritepolicylabelResourceModel, getResponseData map[string]interface{}) *RewritepolicylabelResourceModel {
	tflog.Debug(ctx, "In rewritepolicylabelSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
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
	if val, ok := getResponseData["transform"]; ok && val != nil {
		data.Transform = types.StringValue(val.(string))
	} else {
		data.Transform = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	return data
}
