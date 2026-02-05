package rewritepolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/rewrite"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RewritepolicyResourceModel describes the resource data model.
type RewritepolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"`
	Newname     types.String `tfsdk:"newname"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *RewritepolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rewritepolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the rewrite action to perform if the request or response matches this rewrite policy.\nThere are also some built-in actions which can be used. These are:\n* NOREWRITE - Send the request from the client to the server or response from the server to the client without making any changes in the message.\n* RESET - Resets the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n* DROP - Drop the request without sending a response to the user.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this rewrite policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of messagelog action to use when a request matches this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the rewrite policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite policy\" or 'my rewrite policy').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the rewrite policy. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite policy\" or 'my rewrite policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression against which traffic is evaluated.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character. \n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.",
			},
		},
	}
}

func rewritepolicyGetThePayloadFromtheConfig(ctx context.Context, data *RewritepolicyResourceModel) rewrite.Rewritepolicy {
	tflog.Debug(ctx, "In rewritepolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	rewritepolicy := rewrite.Rewritepolicy{}
	if !data.Action.IsNull() {
		rewritepolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() {
		rewritepolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() {
		rewritepolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() {
		rewritepolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		rewritepolicy.Newname = data.Newname.ValueString()
	}
	if !data.Rule.IsNull() {
		rewritepolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() {
		rewritepolicy.Undefaction = data.Undefaction.ValueString()
	}

	return rewritepolicy
}

func rewritepolicySetAttrFromGet(ctx context.Context, data *RewritepolicyResourceModel, getResponseData map[string]interface{}) *RewritepolicyResourceModel {
	tflog.Debug(ctx, "In rewritepolicySetAttrFromGet Function")

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
