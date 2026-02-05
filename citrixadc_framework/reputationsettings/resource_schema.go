package reputationsettings

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/reputation"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ReputationsettingsResourceModel describes the resource data model.
type ReputationsettingsResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Proxypassword types.String `tfsdk:"proxypassword"`
	Proxyport     types.Int64  `tfsdk:"proxyport"`
	Proxyserver   types.String `tfsdk:"proxyserver"`
	Proxyusername types.String `tfsdk:"proxyusername"`
}

func (r *ReputationsettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the reputationsettings resource.",
			},
			"proxypassword": schema.StringAttribute{
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

func reputationsettingsGetThePayloadFromtheConfig(ctx context.Context, data *ReputationsettingsResourceModel) reputation.Reputationsettings {
	tflog.Debug(ctx, "In reputationsettingsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	reputationsettings := reputation.Reputationsettings{}
	if !data.Proxypassword.IsNull() {
		reputationsettings.Proxypassword = data.Proxypassword.ValueString()
	}
	if !data.Proxyport.IsNull() {
		reputationsettings.Proxyport = utils.IntPtr(int(data.Proxyport.ValueInt64()))
	}
	if !data.Proxyserver.IsNull() {
		reputationsettings.Proxyserver = data.Proxyserver.ValueString()
	}
	if !data.Proxyusername.IsNull() {
		reputationsettings.Proxyusername = data.Proxyusername.ValueString()
	}

	return reputationsettings
}

func reputationsettingsSetAttrFromGet(ctx context.Context, data *ReputationsettingsResourceModel, getResponseData map[string]interface{}) *ReputationsettingsResourceModel {
	tflog.Debug(ctx, "In reputationsettingsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["proxypassword"]; ok && val != nil {
		data.Proxypassword = types.StringValue(val.(string))
	} else {
		data.Proxypassword = types.StringNull()
	}
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
