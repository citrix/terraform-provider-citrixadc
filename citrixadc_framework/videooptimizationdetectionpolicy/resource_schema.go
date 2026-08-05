package videooptimizationdetectionpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VideooptimizationdetectionpolicyResourceModel describes the resource data model.
type VideooptimizationdetectionpolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"`
	Newname     types.String `tfsdk:"newname"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *VideooptimizationdetectionpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the videooptimizationdetectionpolicy resource.",
			},
			// SDK v2: Required (updateable)
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the videooptimization detection action to perform if the request matches this videooptimization detection policy. Built-in actions should be used. These are:\n* DETECT_CLEARTEXT_PD - Cleartext PD is detected and increment related counters.\n* DETECT_CLEARTEXT_ABR - Cleartext ABR is detected and increment related counters.\n* DETECT_ENCRYPTED_ABR - Encrypted ABR is detected and increment related counters.\n* TRIGGER_ENC_ABR_DETECTION - This is potentially encrypted ABR. Internal traffic heuristics algorithms will further process traffic to confirm detection.\n* TRIGGER_CT_ABR_BODY_DETECTION -  This is potentially cleartext ABR. Internal traffic heuristics algorithms will further process traffic to confirm detection.\n* RESET - Reset the client connection by closing it.\n* DROP - Drop the connection without sending a response.",
			},
			// SDK v2: Optional + Computed (updateable)
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this videooptimization detection policy.",
			},
			// SDK v2: Optional + Computed (updateable)
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the videooptimization detection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.Can be modified, removed or renamed.",
			},
			// newname is the rename trigger (NITRO ?action=rename). Not present in SDK v2;
			// kept as a backward-compatible superset. Optional-only: changing it must NOT
			// force replacement (it drives an in-place rename via Update), and it is a pure
			// user input never echoed back by GET (so NOT Computed).
			"newname": schema.StringAttribute{
				Optional:    true,
				Description: "New name for the videooptimization detection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			// SDK v2: Required (updateable)
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression that determines which request or response match the video optimization detection policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			// SDK v2: Optional + Computed (updateable)
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.",
			},
		},
	}
}

// videooptimizationdetectionpolicyGetThePayloadFromthePlan builds the NITRO payload for
// add/update. newname is deliberately excluded (rename-only argument).
func videooptimizationdetectionpolicyGetThePayloadFromthePlan(ctx context.Context, data *VideooptimizationdetectionpolicyResourceModel) videooptimization.Videooptimizationdetectionpolicy {
	tflog.Debug(ctx, "In videooptimizationdetectionpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	videooptimizationdetectionpolicy := videooptimization.Videooptimizationdetectionpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		videooptimizationdetectionpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		videooptimizationdetectionpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		videooptimizationdetectionpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		videooptimizationdetectionpolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of the
	// add/update payload, so it is deliberately excluded.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		videooptimizationdetectionpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		videooptimizationdetectionpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return videooptimizationdetectionpolicy
}

// videooptimizationdetectionpolicySetAttrFromGet reads the GET response into the resource
// model. It preserves user-configured values where NITRO omits them from the GET response
// (omit-on-default trap) and never clobbers the key attribute (rename safety).
func videooptimizationdetectionpolicySetAttrFromGet(ctx context.Context, data *VideooptimizationdetectionpolicyResourceModel, getResponseData map[string]interface{}) *VideooptimizationdetectionpolicyResourceModel {
	tflog.Debug(ctx, "In videooptimizationdetectionpolicySetAttrFromGet Function")

	// Convert API response to model.
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else if data.Action.IsUnknown() {
		data.Action = types.StringNull()
	}
	// comment/logaction/undefaction are Optional+Computed. NITRO omits them from GET when
	// empty; only null them out when still unknown (plan-time), otherwise preserve the
	// configured/prior value to avoid clobbering it (omit-on-default trap).
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
	// name is the user-facing key. After a rename (via newname), the live object name
	// (tracked by data.Id) diverges from the configured name, and GET returns the live
	// (new) name. Overwriting name from GET would clobber the user's configured value and
	// trigger a spurious RequiresReplace diff. Only adopt the GET value when we don't
	// already have one (e.g. on import, where state carries only the ID); else preserve.
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

// videooptimizationdetectionpolicySetAttrFromGetForDatasource faithfully copies every
// field from the GET response. The datasource has no prior plan/state to preserve, so it
// populates the model directly from the API response and sets the ID itself.
func videooptimizationdetectionpolicySetAttrFromGetForDatasource(ctx context.Context, data *VideooptimizationdetectionpolicyResourceModel, getResponseData map[string]interface{}) *VideooptimizationdetectionpolicyResourceModel {
	tflog.Debug(ctx, "In videooptimizationdetectionpolicySetAttrFromGetForDatasource Function")

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
