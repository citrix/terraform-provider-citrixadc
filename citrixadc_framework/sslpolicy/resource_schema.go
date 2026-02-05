package sslpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslpolicyResourceModel describes the resource data model.
type SslpolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Name        types.String `tfsdk:"name"`
	Reqaction   types.String `tfsdk:"reqaction"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *SslpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslpolicy resource.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the built-in or user-defined action to perform on the request. Available built-in actions are NOOP, RESET, DROP, CLIENTAUTH, NOCLIENTAUTH, INTERCEPT AND BYPASS.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new SSL policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.  Cannot be changed after the policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"reqaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the action to be performed on the request. Refer to 'add ssl action' command to add a new action. Builtin actions like NOOP, RESET, DROP, CLIENTAUTH and NOCLIENTAUTH are also allowed.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, against which traffic is evaluated.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the  character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the action to be performed when the result of rule evaluation is undefined. Possible values for control policies: CLIENTAUTH, NOCLIENTAUTH, NOOP, RESET, DROP. Possible values for data policies: NOOP, RESET, DROP and BYPASS",
			},
		},
	}
}

func sslpolicyGetThePayloadFromtheConfig(ctx context.Context, data *SslpolicyResourceModel) ssl.Sslpolicy {
	tflog.Debug(ctx, "In sslpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslpolicy := ssl.Sslpolicy{}
	if !data.Action.IsNull() {
		sslpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() {
		sslpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		sslpolicy.Name = data.Name.ValueString()
	}
	if !data.Reqaction.IsNull() {
		sslpolicy.Reqaction = data.Reqaction.ValueString()
	}
	if !data.Rule.IsNull() {
		sslpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() {
		sslpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return sslpolicy
}

func sslpolicySetAttrFromGet(ctx context.Context, data *SslpolicyResourceModel, getResponseData map[string]interface{}) *SslpolicyResourceModel {
	tflog.Debug(ctx, "In sslpolicySetAttrFromGet Function")

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
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["reqaction"]; ok && val != nil {
		data.Reqaction = types.StringValue(val.(string))
	} else {
		data.Reqaction = types.StringNull()
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
