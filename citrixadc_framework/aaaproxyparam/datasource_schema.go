package aaaproxyparam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func AaaproxyparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"proxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address and Port of the proxy server to be used for HTTP access for this request. Configure in ipaddress:port format (a.b.c.d:e) or as a URL (http://a.b.c.d or http://a.b.c.d:8080).",
			},
			"proxyauthorization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This indicates whether Proxy-Authorization header will be sent or not. Possible values = disabled, basic",
			},
			"proxyusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username that will be sent as part of Basic Proxy-Authorization header.",
			},
			"proxypassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Password that will be sent as part of Basic Proxy-Authorization header.",
			},
		},
	}
}

type AaaproxyparamDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Proxy              types.String `tfsdk:"proxy"`
	Proxyauthorization types.String `tfsdk:"proxyauthorization"`
	Proxyusername      types.String `tfsdk:"proxyusername"`
	Proxypassword      types.String `tfsdk:"proxypassword"`
}

func aaaproxyparamDataSourceSetAttrFromGet(ctx context.Context, data *AaaproxyparamDataSourceModel, getResponseData map[string]interface{}) *AaaproxyparamDataSourceModel {
	tflog.Debug(ctx, "In aaaproxyparamDataSourceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["proxy"]; ok && val != nil {
		data.Proxy = types.StringValue(val.(string))
	} else {
		data.Proxy = types.StringNull()
	}
	if val, ok := getResponseData["proxyauthorization"]; ok && val != nil {
		data.Proxyauthorization = types.StringValue(val.(string))
	} else {
		data.Proxyauthorization = types.StringNull()
	}
	if val, ok := getResponseData["proxyusername"]; ok && val != nil {
		data.Proxyusername = types.StringValue(val.(string))
	} else {
		data.Proxyusername = types.StringNull()
	}
	// proxypassword is not returned by NITRO API in usable form (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 1: No unique attributes - static ID (singleton)
	data.Id = types.StringValue("aaaproxyparam-config")

	return data
}
