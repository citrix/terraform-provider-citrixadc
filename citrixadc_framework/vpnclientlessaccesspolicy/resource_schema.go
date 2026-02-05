package vpnclientlessaccesspolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnclientlessaccesspolicyResourceModel describes the resource data model.
type VpnclientlessaccesspolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Profilename types.String `tfsdk:"profilename"`
	Rule        types.String `tfsdk:"rule"`
}

func (r *VpnclientlessaccesspolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnclientlessaccesspolicy resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the new clientless access policy.",
			},
			"profilename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile to invoke for the clientless access.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, or name of a named expression, specifying the traffic that matches the policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func vpnclientlessaccesspolicyGetThePayloadFromtheConfig(ctx context.Context, data *VpnclientlessaccesspolicyResourceModel) vpn.Vpnclientlessaccesspolicy {
	tflog.Debug(ctx, "In vpnclientlessaccesspolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnclientlessaccesspolicy := vpn.Vpnclientlessaccesspolicy{}
	if !data.Name.IsNull() {
		vpnclientlessaccesspolicy.Name = data.Name.ValueString()
	}
	if !data.Profilename.IsNull() {
		vpnclientlessaccesspolicy.Profilename = data.Profilename.ValueString()
	}
	if !data.Rule.IsNull() {
		vpnclientlessaccesspolicy.Rule = data.Rule.ValueString()
	}

	return vpnclientlessaccesspolicy
}

func vpnclientlessaccesspolicySetAttrFromGet(ctx context.Context, data *VpnclientlessaccesspolicyResourceModel, getResponseData map[string]interface{}) *VpnclientlessaccesspolicyResourceModel {
	tflog.Debug(ctx, "In vpnclientlessaccesspolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["profilename"]; ok && val != nil {
		data.Profilename = types.StringValue(val.(string))
	} else {
		data.Profilename = types.StringNull()
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
