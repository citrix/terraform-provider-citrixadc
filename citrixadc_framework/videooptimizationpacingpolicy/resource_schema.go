package videooptimizationpacingpolicy

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

// VideooptimizationpacingpolicyResourceModel describes the resource data model.
type VideooptimizationpacingpolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"`
	Newname     types.String `tfsdk:"newname"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *VideooptimizationpacingpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the videooptimizationpacingpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the videooptimization pacing action to perform if the request matches this videooptimization pacing policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this videooptimization pacing policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the videooptimization pacing policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.Can be modified, removed or renamed.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the videooptimization pacing policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression that determines which request or response match the video optimization pacing policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.",
			},
		},
	}
}

func videooptimizationpacingpolicyGetThePayloadFromtheConfig(ctx context.Context, data *VideooptimizationpacingpolicyResourceModel) videooptimization.Videooptimizationpacingpolicy {
	tflog.Debug(ctx, "In videooptimizationpacingpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	videooptimizationpacingpolicy := videooptimization.Videooptimizationpacingpolicy{}
	if !data.Action.IsNull() {
		videooptimizationpacingpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() {
		videooptimizationpacingpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() {
		videooptimizationpacingpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() {
		videooptimizationpacingpolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		videooptimizationpacingpolicy.Newname = data.Newname.ValueString()
	}
	if !data.Rule.IsNull() {
		videooptimizationpacingpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() {
		videooptimizationpacingpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return videooptimizationpacingpolicy
}

func videooptimizationpacingpolicySetAttrFromGet(ctx context.Context, data *VideooptimizationpacingpolicyResourceModel, getResponseData map[string]interface{}) *VideooptimizationpacingpolicyResourceModel {
	tflog.Debug(ctx, "In videooptimizationpacingpolicySetAttrFromGet Function")

	// Convert API response to model
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
