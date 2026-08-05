package ntpparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ntp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
			// SDK v2 parity: Optional+Computed, no Default (value is read back from the ADC).
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Apply NTP authentication, which enables the NTP client (Citrix ADC) to verify that the server is in fact known and trusted.",
			},
			"autokeylogsec": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Autokey protocol requires the keys to be refreshed periodically. This parameter specifies the interval between regenerations of new session keys. In seconds, expressed as a power of 2.",
			},
			"revokelogsec": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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
	if !data.Authentication.IsNull() && !data.Authentication.IsUnknown() {
		ntpparam.Authentication = data.Authentication.ValueString()
	}
	if !data.Autokeylogsec.IsNull() && !data.Autokeylogsec.IsUnknown() {
		ntpparam.Autokeylogsec = utils.IntPtr(int(data.Autokeylogsec.ValueInt64()))
	}
	if !data.Revokelogsec.IsNull() && !data.Revokelogsec.IsUnknown() {
		ntpparam.Revokelogsec = utils.IntPtr(int(data.Revokelogsec.ValueInt64()))
	}
	if !data.Trustedkey.IsNull() && !data.Trustedkey.IsUnknown() {
		var trustedkeyList []int64
		data.Trustedkey.ElementsAs(ctx, &trustedkeyList, false)
		intList := make([]int, len(trustedkeyList))
		for i, v := range trustedkeyList {
			intList[i] = int(v)
		}
		ntpparam.Trustedkey = intList
	}

	return ntpparam
}

func ntpparamSetAttrFromGet(ctx context.Context, data *NtpparamResourceModel, getResponseData map[string]interface{}) *NtpparamResourceModel {
	tflog.Debug(ctx, "In ntpparamSetAttrFromGet Function")

	// Convert API response to model. Else-branches only null the value when it is
	// Unknown so a known/configured value that NITRO omits from GET (omit-on-default
	// trap) is preserved rather than clobbered.
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else if data.Authentication.IsUnknown() {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["autokeylogsec"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Autokeylogsec = types.Int64Value(intVal)
		}
	} else if data.Autokeylogsec.IsUnknown() {
		data.Autokeylogsec = types.Int64Null()
	}
	if val, ok := getResponseData["revokelogsec"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Revokelogsec = types.Int64Value(intVal)
		}
	} else if data.Revokelogsec.IsUnknown() {
		data.Revokelogsec = types.Int64Null()
	}
	if val, ok := getResponseData["trustedkey"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			// NITRO may return the trusted-key ids as strings or as JSON numbers;
			// ConvertToInt64 handles int/float64/string uniformly.
			int64List := make([]int64, 0, len(sliceVal))
			for _, item := range sliceVal {
				if intVal, err := utils.ConvertToInt64(item); err == nil {
					int64List = append(int64List, intVal)
				}
			}
			listValue, _ := types.ListValueFrom(ctx, types.Int64Type, int64List)
			data.Trustedkey = listValue
		} else if data.Trustedkey.IsUnknown() {
			data.Trustedkey = types.ListNull(types.Int64Type)
		}
	} else if data.Trustedkey.IsUnknown() {
		data.Trustedkey = types.ListNull(types.Int64Type)
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID (singleton)
	data.Id = types.StringValue("ntpparam-config")

	return data
}
