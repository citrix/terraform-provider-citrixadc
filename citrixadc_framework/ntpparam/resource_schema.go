package ntpparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ntp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NtpparamResourceModel describes the resource data model.
type NtpparamResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Authentication types.String `tfsdk:"authentication"`
	Autokeylogsec  types.Int64  `tfsdk:"autokeylogsec"`
	Revokelogsec   types.Int64  `tfsdk:"revokelogsec"`
	Trustedkey     types.List   `tfsdk:"trustedkey"`
}

func (r *NtpparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ntpparam resource.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Apply NTP authentication, which enables the NTP client (Citrix ADC) to verify that the server is in fact known and trusted.",
			},
			"autokeylogsec": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(12),
				Description: "Autokey protocol requires the keys to be refreshed periodically. This parameter specifies the interval between regenerations of new session keys. In seconds, expressed as a power of 2.",
			},
			"revokelogsec": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(16),
				Description: "Interval between re-randomizations of the autokey seeds to prevent brute-force attacks on the autokey algorithms.",
			},
			"trustedkey": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Key identifiers that are trusted for server authentication with symmetric key cryptography in the keys file.",
			},
		},
	}
}

func ntpparamGetThePayloadFromtheConfig(ctx context.Context, data *NtpparamResourceModel) ntp.Ntpparam {
	tflog.Debug(ctx, "In ntpparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ntpparam := ntp.Ntpparam{}
	if !data.Authentication.IsNull() {
		ntpparam.Authentication = data.Authentication.ValueString()
	}
	if !data.Autokeylogsec.IsNull() {
		ntpparam.Autokeylogsec = utils.IntPtr(int(data.Autokeylogsec.ValueInt64()))
	}
	if !data.Revokelogsec.IsNull() {
		ntpparam.Revokelogsec = utils.IntPtr(int(data.Revokelogsec.ValueInt64()))
	}

	return ntpparam
}

func ntpparamSetAttrFromGet(ctx context.Context, data *NtpparamResourceModel, getResponseData map[string]interface{}) *NtpparamResourceModel {
	tflog.Debug(ctx, "In ntpparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["autokeylogsec"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Autokeylogsec = types.Int64Value(intVal)
		}
	} else {
		data.Autokeylogsec = types.Int64Null()
	}
	if val, ok := getResponseData["revokelogsec"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Revokelogsec = types.Int64Value(intVal)
		}
	} else {
		data.Revokelogsec = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("ntpparam-config")

	return data
}
