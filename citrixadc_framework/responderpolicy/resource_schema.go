package responderpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ResponderpolicyResourceModel describes the resource data model.
type ResponderpolicyResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Action        types.String `tfsdk:"action"`
	Appflowaction types.String `tfsdk:"appflowaction"`
	Comment       types.String `tfsdk:"comment"`
	Logaction     types.String `tfsdk:"logaction"`
	Name          types.String `tfsdk:"name"`
	Newname       types.String `tfsdk:"newname"`
	Rule          types.String `tfsdk:"rule"`
	Undefaction   types.String `tfsdk:"undefaction"`
}

func (r *ResponderpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the responderpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the responder action to perform if the request matches this responder policy. There are also some built-in actions which can be used. These are:\n* NOOP - Send the request to the protected server instead of responding to it.\n* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n* DROP - Drop the request without sending a response to the user.",
			},
			"appflowaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AppFlow action to invoke for requests that match this policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this responder policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the responder policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the responder policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder policy\" or 'my responder policy').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the responder policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder policy\" or 'my responder policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression that the policy uses to determine whether to respond to the specified request.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.",
			},
		},
	}
}

func responderpolicyGetThePayloadFromtheConfig(ctx context.Context, data *ResponderpolicyResourceModel) responder.Responderpolicy {
	tflog.Debug(ctx, "In responderpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	responderpolicy := responder.Responderpolicy{}
	if !data.Action.IsNull() {
		responderpolicy.Action = data.Action.ValueString()
	}
	if !data.Appflowaction.IsNull() {
		responderpolicy.Appflowaction = data.Appflowaction.ValueString()
	}
	if !data.Comment.IsNull() {
		responderpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() {
		responderpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() {
		responderpolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		responderpolicy.Newname = data.Newname.ValueString()
	}
	if !data.Rule.IsNull() {
		responderpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() {
		responderpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return responderpolicy
}

func responderpolicySetAttrFromGet(ctx context.Context, data *ResponderpolicyResourceModel, getResponseData map[string]interface{}) *ResponderpolicyResourceModel {
	tflog.Debug(ctx, "In responderpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["appflowaction"]; ok && val != nil {
		data.Appflowaction = types.StringValue(val.(string))
	} else {
		data.Appflowaction = types.StringNull()
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
