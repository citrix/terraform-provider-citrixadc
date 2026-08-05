package nsweblogparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsweblogparamResourceModel describes the resource data model.
type NsweblogparamResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Buffersizemb  types.Int64  `tfsdk:"buffersizemb"`
	Customreqhdrs types.List   `tfsdk:"customreqhdrs"`
	Customrsphdrs types.List   `tfsdk:"customrsphdrs"`
}

func (r *NsweblogparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsweblogparam resource.",
			},
			"buffersizemb": schema.Int64Attribute{
				// SDK v2 parity: Optional + Computed, no Default (value read from ADC).
				Optional:    true,
				Computed:    true,
				Description: "Buffer size, in MB, allocated for log transaction data on the system. The maximum value is limited to the memory available on the system.",
			},
			"customreqhdrs": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name(s) of HTTP request headers whose values should be exported by the Web Logging feature.",
			},
			"customrsphdrs": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name(s) of HTTP response headers whose values should be exported by the Web Logging feature.",
			},
		},
	}
}

func nsweblogparamGetThePayloadFromtheConfig(ctx context.Context, data *NsweblogparamResourceModel) ns.Nsweblogparam {
	tflog.Debug(ctx, "In nsweblogparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsweblogparam := ns.Nsweblogparam{}
	if !data.Buffersizemb.IsNull() && !data.Buffersizemb.IsUnknown() {
		nsweblogparam.Buffersizemb = utils.IntPtr(int(data.Buffersizemb.ValueInt64()))
	}
	if !data.Customreqhdrs.IsNull() && !data.Customreqhdrs.IsUnknown() {
		var customreqhdrsList []string
		data.Customreqhdrs.ElementsAs(ctx, &customreqhdrsList, false)
		nsweblogparam.Customreqhdrs = customreqhdrsList
	}
	if !data.Customrsphdrs.IsNull() && !data.Customrsphdrs.IsUnknown() {
		var customrsphdrsList []string
		data.Customrsphdrs.ElementsAs(ctx, &customrsphdrsList, false)
		nsweblogparam.Customrsphdrs = customrsphdrsList
	}

	return nsweblogparam
}

func nsweblogparamSetAttrFromGet(ctx context.Context, data *NsweblogparamResourceModel, getResponseData map[string]interface{}) *NsweblogparamResourceModel {
	tflog.Debug(ctx, "In nsweblogparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["buffersizemb"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Buffersizemb = types.Int64Value(intVal)
		}
	} else if data.Buffersizemb.IsUnknown() {
		data.Buffersizemb = types.Int64Null()
	}
	if val, ok := getResponseData["customreqhdrs"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(v))
			data.Customreqhdrs = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Customreqhdrs = listValue
		default:
			if data.Customreqhdrs.IsUnknown() {
				data.Customreqhdrs = types.ListNull(types.StringType)
			}
		}
	} else if data.Customreqhdrs.IsUnknown() {
		data.Customreqhdrs = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["customrsphdrs"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(v))
			data.Customrsphdrs = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Customrsphdrs = listValue
		default:
			if data.Customrsphdrs.IsUnknown() {
				data.Customrsphdrs = types.ListNull(types.StringType)
			}
		}
	} else if data.Customrsphdrs.IsUnknown() {
		data.Customrsphdrs = types.ListNull(types.StringType)
	}

	// Set ID for the resource
	// Case 1: No unique attributes (singleton) - static ID
	data.Id = types.StringValue("nsweblogparam-config")

	return data
}
