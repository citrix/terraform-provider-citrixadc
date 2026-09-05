package vpnsecureprivateaccessprofile

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func VpnsecureprivateaccessprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "name of Secure Private Access profile.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Public URL for your Secure Private Access site or load balancing server.",
			},
			"customerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Customer ID of the citrix cloud customer.",
			},
			"chromeenterprisepremiummode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Secure Private Access Chrome Enterprise Premium mode of operation. Possible values = OFF, WITH_PARTNER_CONNECTOR, WITHOUT_PARTNER_CONNECTOR",
			},
			"googlecustomerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Your organization's unique ID on Google's Admin console Profile settings.",
			},
			"googlesecuritygatewayid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The ID of the Google Secure Gateway.",
			},
			"forceclienttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Automatically configures the session for Citrix Secure Access client connectivity. Possible values = ON, OFF",
			},
			"sharedsecret": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Secure Private Access Shared Secret.",
			},
		},
	}
}

type VpnsecureprivateaccessprofileDataSourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Name                        types.String `tfsdk:"name"`
	Url                         types.String `tfsdk:"url"`
	Customerid                  types.String `tfsdk:"customerid"`
	Chromeenterprisepremiummode types.String `tfsdk:"chromeenterprisepremiummode"`
	Googlecustomerid            types.String `tfsdk:"googlecustomerid"`
	Googlesecuritygatewayid     types.String `tfsdk:"googlesecuritygatewayid"`
	Forceclienttype             types.String `tfsdk:"forceclienttype"`
	Sharedsecret                types.String `tfsdk:"sharedsecret"`
}

func vpnsecureprivateaccessprofileDataSourceSetAttrFromGet(ctx context.Context, data *VpnsecureprivateaccessprofileDataSourceModel, getResponseData map[string]interface{}) *VpnsecureprivateaccessprofileDataSourceModel {
	tflog.Debug(ctx, "In vpnsecureprivateaccessprofileDataSourceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["url"]; ok && val != nil {
		data.Url = types.StringValue(val.(string))
	} else {
		data.Url = types.StringNull()
	}
	if val, ok := getResponseData["customerid"]; ok && val != nil {
		data.Customerid = types.StringValue(val.(string))
	} else {
		data.Customerid = types.StringNull()
	}
	if val, ok := getResponseData["chromeenterprisepremiummode"]; ok && val != nil {
		data.Chromeenterprisepremiummode = types.StringValue(val.(string))
	} else {
		data.Chromeenterprisepremiummode = types.StringNull()
	}
	if val, ok := getResponseData["googlecustomerid"]; ok && val != nil {
		data.Googlecustomerid = types.StringValue(val.(string))
	} else {
		data.Googlecustomerid = types.StringNull()
	}
	if val, ok := getResponseData["googlesecuritygatewayid"]; ok && val != nil {
		data.Googlesecuritygatewayid = types.StringValue(val.(string))
	} else {
		data.Googlesecuritygatewayid = types.StringNull()
	}
	if val, ok := getResponseData["forceclienttype"]; ok && val != nil {
		data.Forceclienttype = types.StringValue(val.(string))
	} else {
		data.Forceclienttype = types.StringNull()
	}
	// sharedsecret is not returned by NITRO API in usable form (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
