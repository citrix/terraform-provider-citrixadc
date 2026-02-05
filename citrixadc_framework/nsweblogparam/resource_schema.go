package nsweblogparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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
				Optional:    true,
				Default:     int64default.StaticInt64(16),
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
	if !data.Buffersizemb.IsNull() {
		nsweblogparam.Buffersizemb = utils.IntPtr(int(data.Buffersizemb.ValueInt64()))
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
	} else {
		data.Buffersizemb = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nsweblogparam-config")

	return data
}
