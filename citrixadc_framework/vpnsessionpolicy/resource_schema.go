package vpnsessionpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnsessionpolicyResourceModel describes the resource data model.
type VpnsessionpolicyResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Rule   types.String `tfsdk:"rule"`
}

func (r *VpnsessionpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnsessionpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action to be applied by the new session policy if the rule criteria are met.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new session policy that is applied after the user logs on to Citrix Gateway.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, or name of a named expression, specifying the traffic that matches the policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func vpnsessionpolicyGetThePayloadFromtheConfig(ctx context.Context, data *VpnsessionpolicyResourceModel) vpn.Vpnsessionpolicy {
	tflog.Debug(ctx, "In vpnsessionpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnsessionpolicy := vpn.Vpnsessionpolicy{}
	if !data.Action.IsNull() {
		vpnsessionpolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() {
		vpnsessionpolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() {
		vpnsessionpolicy.Rule = data.Rule.ValueString()
	}

	return vpnsessionpolicy
}

func vpnsessionpolicySetAttrFromGet(ctx context.Context, data *VpnsessionpolicyResourceModel, getResponseData map[string]interface{}) *VpnsessionpolicyResourceModel {
	tflog.Debug(ctx, "In vpnsessionpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
