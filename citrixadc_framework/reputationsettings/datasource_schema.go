package reputationsettings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

func ReputationsettingsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"proxypassword": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Password with which user logs on.",
			},
			"proxyport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy server port.",
			},
			"proxyserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy server IP to get Reputation data.",
			},
			"proxyusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Username",
			},
		},
	}
}

type ReputationsettingsDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Proxypassword types.String `tfsdk:"proxypassword"`
	Proxyport     types.Int64  `tfsdk:"proxyport"`
	Proxyserver   types.String `tfsdk:"proxyserver"`
	Proxyusername types.String `tfsdk:"proxyusername"`
}

func reputationsettingsDataSourceSetAttrFromGet(ctx context.Context, data *ReputationsettingsDataSourceModel, getResponseData map[string]interface{}) *ReputationsettingsDataSourceModel {
	tflog.Debug(ctx, "In reputationsettingsDataSourceSetAttrFromGet Function")

	// Convert API response to model
	// proxypassword is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["proxyport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Proxyport = types.Int64Value(intVal)
		}
	} else {
		data.Proxyport = types.Int64Null()
	}
	if val, ok := getResponseData["proxyserver"]; ok && val != nil {
		data.Proxyserver = types.StringValue(val.(string))
	} else {
		data.Proxyserver = types.StringNull()
	}
	if val, ok := getResponseData["proxyusername"]; ok && val != nil {
		data.Proxyusername = types.StringValue(val.(string))
	} else {
		data.Proxyusername = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("reputationsettings-config")

	return data
}
