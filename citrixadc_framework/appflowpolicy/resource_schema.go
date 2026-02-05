package appflowpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppflowpolicyResourceModel describes the resource data model.
type AppflowpolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Name        types.String `tfsdk:"name"`
	Newname     types.String `tfsdk:"newname"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *AppflowpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appflowpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the action to be associated with this policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow policy\" or 'my appflow policy').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the policy. Must begin with an ASCII alphabetic or underscore (_)character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow policy\" or 'my appflow policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression or other value against which the traffic is evaluated. Must be a Boolean expression.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the appflow action to be associated with this policy when an undef event occurs.",
			},
		},
	}
}

func appflowpolicyGetThePayloadFromtheConfig(ctx context.Context, data *AppflowpolicyResourceModel) appflow.Appflowpolicy {
	tflog.Debug(ctx, "In appflowpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appflowpolicy := appflow.Appflowpolicy{}
	if !data.Action.IsNull() {
		appflowpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() {
		appflowpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		appflowpolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		appflowpolicy.Newname = data.Newname.ValueString()
	}
	if !data.Rule.IsNull() {
		appflowpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() {
		appflowpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return appflowpolicy
}

func appflowpolicySetAttrFromGet(ctx context.Context, data *AppflowpolicyResourceModel, getResponseData map[string]interface{}) *AppflowpolicyResourceModel {
	tflog.Debug(ctx, "In appflowpolicySetAttrFromGet Function")

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
