package contentinspectionpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ContentinspectionpolicyResourceModel describes the resource data model.
type ContentinspectionpolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"`
	Newname     types.String `tfsdk:"newname"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *ContentinspectionpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the contentinspectionpolicy resource.",
			},
			// SDK v2: Required (updateable).
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the contentInspection action to perform if the request matches this contentInspection policy.\n    There are also some built-in actions which can be used. These are:\n    * NOINSPECTION - Send the request from the client to the server or response from the server to the client without sending it to Inspection device for Content Inspection.\n    * RESET - Resets the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n    * DROP - Drop the request without sending a response to the user.",
			},
			// SDK v2: Optional+Computed (updateable).
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this contentInspection policy.",
			},
			// SDK v2: Optional+Computed (updateable).
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace. The key
			// itself is immutable; in-place rename is driven by `newname` (see Update).
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the contentInspection policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the contentInspection policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my contentInspection policy\" or 'my contentInspection policy').",
			},
			// newname is the rename trigger (NITRO ?action=rename). It must NOT force
			// replacement - it drives an in-place rename via Update. It is Optional
			// only (a pure user input never echoed back by GET), so it is neither
			// Computed nor RequiresReplace. Not present in SDK v2 config, so this is
			// additive and backward compatible.
			"newname": schema.StringAttribute{
				Optional:    true,
				Description: "New name for the contentInspection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my contentInspection policy\" or 'my contentInspection policy').",
			},
			// SDK v2: Required (updateable).
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression that the policy uses to determine whether to execute the specified action.",
			},
			// SDK v2: Optional+Computed (updateable).
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.",
			},
		},
	}
}

// contentinspectionpolicyGetThePayloadFromthePlan builds the create (add) payload.
// `newname` is a rename-only argument and is deliberately excluded from the add body.
func contentinspectionpolicyGetThePayloadFromthePlan(ctx context.Context, data *ContentinspectionpolicyResourceModel) contentinspection.Contentinspectionpolicy {
	tflog.Debug(ctx, "In contentinspectionpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	contentinspectionpolicy := contentinspection.Contentinspectionpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		contentinspectionpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		contentinspectionpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		contentinspectionpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		contentinspectionpolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename); excluded from add POST.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		contentinspectionpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		contentinspectionpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return contentinspectionpolicy
}

// contentinspectionpolicyGetTheUpdatablePayloadFromThePlan builds the update (PUT)
// payload restricted to NITRO-updatable fields. The caller sets the Name field to
// the live object name (data.Id) to remain correct after a rename. `newname` is
// rename-only and excluded here.
func contentinspectionpolicyGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *ContentinspectionpolicyResourceModel) contentinspection.Contentinspectionpolicy {
	tflog.Debug(ctx, "In contentinspectionpolicyGetTheUpdatablePayloadFromThePlan Function")

	contentinspectionpolicy := contentinspection.Contentinspectionpolicy{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		contentinspectionpolicy.Name = data.Name.ValueString()
	}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		contentinspectionpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		contentinspectionpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		contentinspectionpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		contentinspectionpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		contentinspectionpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return contentinspectionpolicy
}

// contentinspectionpolicySetAttrFromGet is the RESOURCE state setter. It preserves
// user-facing configured values (and the rename-aware `name`/`id`) rather than
// blindly overwriting them from the GET response.
func contentinspectionpolicySetAttrFromGet(ctx context.Context, data *ContentinspectionpolicyResourceModel, getResponseData map[string]interface{}) *ContentinspectionpolicyResourceModel {
	tflog.Debug(ctx, "In contentinspectionpolicySetAttrFromGet Function")

	// Convert API response to model.
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else if data.Action.IsUnknown() {
		data.Action = types.StringNull()
	}
	// comment/logaction/undefaction are Optional+Computed. NITRO omits them from the
	// GET response when empty. Only overwrite when present; otherwise resolve an
	// UNKNOWN (Computed, unconfigured) value to null while preserving a configured
	// value - prevents "inconsistent result after apply".
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else if data.Logaction.IsUnknown() {
		data.Logaction = types.StringNull()
	}
	// name is the user-facing key. After a rename (via newname) the live object name
	// (tracked by data.Id) diverges from the configured name, and GET returns the
	// live (new) name. Overwriting name from GET would clobber the user's configured
	// value and trigger a spurious RequiresReplace diff. Only adopt the GET value
	// when we don't already have one (e.g. on import, where state carries only the
	// ID); otherwise preserve.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else if data.Rule.IsUnknown() {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else if data.Undefaction.IsUnknown() {
		data.Undefaction = types.StringNull()
	}

	return data
}

// contentinspectionpolicySetAttrFromGetForDatasource is the DATASOURCE state setter.
// The datasource has no prior plan/state to preserve, so it faithfully copies every
// field from the GET response and sets the ID itself.
func contentinspectionpolicySetAttrFromGetForDatasource(ctx context.Context, data *ContentinspectionpolicyResourceModel, getResponseData map[string]interface{}) *ContentinspectionpolicyResourceModel {
	tflog.Debug(ctx, "In contentinspectionpolicySetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else {
		data.Logaction = types.StringNull()
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
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
