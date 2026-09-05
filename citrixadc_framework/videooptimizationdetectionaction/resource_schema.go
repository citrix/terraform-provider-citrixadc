package videooptimizationdetectionaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset — mirroring the
// SDK v2 unset-on-remove contract. Without it an Optional+Computed attribute is
// "sticky": the prior value is carried forward and removal is a silent no-op.
// It intentionally does nothing when the config still carries a value, on create
// (no prior state), or when the prior value is already empty (avoids churn).
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// VideooptimizationdetectionactionResourceModel describes the resource data model.
type VideooptimizationdetectionactionResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Type    types.String `tfsdk:"type"`
}

func (r *VideooptimizationdetectionactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the videooptimizationdetectionaction resource.",
			},
			"comment": schema.StringAttribute{
				// Optional + Computed. NITRO omits comment from GET when empty; its
				// server default is the empty string. Instead of a schema Default of ""
				// (which breaks import round-trip: import reads back NULL != ""), the
				// unsetOnRemoveStringModifier marks the value unknown when removed from
				// config while a prior non-empty value exists, so removal still drives
				// Update -> NITRO ?action=unset while import round-trips cleanly.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Comment. Any type of information about this video optimization detection action.",
			},
			"name": schema.StringAttribute{
				// SDK v2 parity: name was Required + ForceNew -> Required + RequiresReplace.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the video optimization detection action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it must
				// NOT force replacement - it drives an in-place rename via Update. Not
				// Computed: it is a pure user input, never echoed back by GET.
				Optional:    true,
				Description: "New name for the videooptimization detection action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"type": schema.StringAttribute{
				// SDK v2 parity: type was Required (Computed:false), not ForceNew and
				// updateable -> Required, no RequiresReplace.
				Required:    true,
				Description: "Type of video optimization action. Available settings function as follows:\n* clear_text_pd - Cleartext PD type is detected.\n* clear_text_abr - Cleartext ABR is detected.\n* encrypted_abr - Encrypted ABR is detected.\n* trigger_enc_abr - Possible encrypted ABR is detected.\n* trigger_body_detection - Possible cleartext ABR is detected. Triggers body content detection.",
			},
		},
	}
}

func videooptimizationdetectionactionGetThePayloadFromthePlan(ctx context.Context, data *VideooptimizationdetectionactionResourceModel) videooptimization.Videooptimizationdetectionaction {
	tflog.Debug(ctx, "In videooptimizationdetectionactionGetThePayloadFromthePlan Function")

	// Create API request body from the model (add payload)
	videooptimizationdetectionaction := videooptimization.Videooptimizationdetectionaction{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		videooptimizationdetectionaction.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		videooptimizationdetectionaction.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of the
	// add payload, so it is deliberately excluded from the create POST body.
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		videooptimizationdetectionaction.Type = data.Type.ValueString()
	}

	return videooptimizationdetectionaction
}

func videooptimizationdetectionactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *VideooptimizationdetectionactionResourceModel) videooptimization.Videooptimizationdetectionaction {
	tflog.Debug(ctx, "In videooptimizationdetectionactionGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields.
	videooptimizationdetectionaction := videooptimization.Videooptimizationdetectionaction{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		videooptimizationdetectionaction.Comment = data.Comment.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		videooptimizationdetectionaction.Type = data.Type.ValueString()
	}
	// name is set by the caller to the current live name (== data.Id). newname is
	// rename-only and excluded from the update payload.

	return videooptimizationdetectionaction
}

func videooptimizationdetectionactionSetAttrFromGet(ctx context.Context, data *VideooptimizationdetectionactionResourceModel, getResponseData map[string]interface{}) *VideooptimizationdetectionactionResourceModel {
	tflog.Debug(ctx, "In videooptimizationdetectionactionSetAttrFromGet Function")

	// Convert API response to model.
	// comment: NITRO omits it from GET when empty. Only overwrite when present.
	// Otherwise resolve unknown -> null (create with no config) but never clobber a
	// known/configured value (omit-on-default trap).
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	// name is the user-facing key. After a rename (via newname) the live object name
	// (tracked by data.Id) diverges from the configured name, and GET returns the live
	// (new) name. Overwriting name from GET would clobber the user's configured value
	// and trigger a spurious RequiresReplace diff. So only adopt the GET value when we
	// don't already have one (e.g. on import, where state carries only the ID).
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	}

	return data
}

// videooptimizationdetectionactionSetAttrFromGetForDatasource faithfully copies every
// field from the GET response. The datasource has no prior plan/state to preserve, so
// it must populate the model directly from the API response and set the ID itself.
func videooptimizationdetectionactionSetAttrFromGetForDatasource(ctx context.Context, data *VideooptimizationdetectionactionResourceModel, getResponseData map[string]interface{}) *VideooptimizationdetectionactionResourceModel {
	tflog.Debug(ctx, "In videooptimizationdetectionactionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
