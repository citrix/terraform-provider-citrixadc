package reputationsettings

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/reputation"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ReputationsettingsResourceModel describes the resource data model.
type ReputationsettingsResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Proxypassword          types.String `tfsdk:"proxypassword"`
	ProxypasswordWo        types.String `tfsdk:"proxypassword_wo"`
	ProxypasswordWoVersion types.Int64  `tfsdk:"proxypassword_wo_version"`
	Proxyport              types.Int64  `tfsdk:"proxyport"`
	Proxyserver            types.String `tfsdk:"proxyserver"`
	Proxyusername          types.String `tfsdk:"proxyusername"`
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
				Sensitive:   true,
				Description: "Password with which user logs on.",
			},
			"proxypassword_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Password with which user logs on.",
			},
			"proxypassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a proxypassword_wo update.",
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

func reputationsettingsGetThePayloadFromthePlan(ctx context.Context, data *ReputationsettingsResourceModel) reputation.Reputationsettings {
	tflog.Debug(ctx, "In reputationsettingsGetThePayloadFromthePlan Function")

	// Create API request body from the model
	reputationsettings := reputation.Reputationsettings{}
	if !data.Proxypassword.IsNull() && !data.Proxypassword.IsUnknown() {
		reputationsettings.Proxypassword = data.Proxypassword.ValueString()
	}
	// Skip write-only attribute: proxypassword_wo
	// Skip version tracker attribute: proxypassword_wo_version
	if !data.Proxyport.IsNull() && !data.Proxyport.IsUnknown() {
		reputationsettings.Proxyport = utils.IntPtr(int(data.Proxyport.ValueInt64()))
	}
	if !data.Proxyserver.IsNull() && !data.Proxyserver.IsUnknown() {
		reputationsettings.Proxyserver = data.Proxyserver.ValueString()
	}
	if !data.Proxyusername.IsNull() && !data.Proxyusername.IsUnknown() {
		reputationsettings.Proxyusername = data.Proxyusername.ValueString()
	}

	return reputationsettings
}

func reputationsettingsGetThePayloadFromtheConfig(ctx context.Context, data *ReputationsettingsResourceModel, payload *reputation.Reputationsettings) {
	tflog.Debug(ctx, "In reputationsettingsGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: proxypassword_wo -> proxypassword
	if !data.ProxypasswordWo.IsNull() {
		proxypasswordWo := data.ProxypasswordWo.ValueString()
		if proxypasswordWo != "" {
			payload.Proxypassword = proxypasswordWo
		}
	}
}

func reputationsettingsSetAttrFromGet(ctx context.Context, data *ReputationsettingsResourceModel, getResponseData map[string]interface{}) *ReputationsettingsResourceModel {
	tflog.Debug(ctx, "In reputationsettingsSetAttrFromGet Function")

	// Convert API response to model
	// proxypassword is not returned by NITRO API (secret/ephemeral) - retain from config
	// proxypassword_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// proxypassword_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
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
