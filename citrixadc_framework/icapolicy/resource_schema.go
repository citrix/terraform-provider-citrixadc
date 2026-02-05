package icapolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IcapolicyResourceModel describes the resource data model.
type IcapolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Action    types.String `tfsdk:"action"`
	Comment   types.String `tfsdk:"comment"`
	Logaction types.String `tfsdk:"logaction"`
	Name      types.String `tfsdk:"name"`
	Newname   types.String `tfsdk:"newname"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *IcapolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the icapolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the ica action to be associated with this policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this ICA policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica policy\" or 'my ica policy').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the policy. Must begin with an ASCII alphabetic or underscore (_)character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), s\npace, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\n\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica policy\" or 'my ica policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression or other value against which the traffic is evaluated. Must be a Boolean expression.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func icapolicyGetThePayloadFromtheConfig(ctx context.Context, data *IcapolicyResourceModel) ica.Icapolicy {
	tflog.Debug(ctx, "In icapolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	icapolicy := ica.Icapolicy{}
	if !data.Action.IsNull() {
		icapolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() {
		icapolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() {
		icapolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() {
		icapolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		icapolicy.Newname = data.Newname.ValueString()
	}
	if !data.Rule.IsNull() {
		icapolicy.Rule = data.Rule.ValueString()
	}

	return icapolicy
}

func icapolicySetAttrFromGet(ctx context.Context, data *IcapolicyResourceModel, getResponseData map[string]interface{}) *IcapolicyResourceModel {
	tflog.Debug(ctx, "In icapolicySetAttrFromGet Function")

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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
