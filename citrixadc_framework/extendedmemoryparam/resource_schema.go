package extendedmemoryparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ExtendedmemoryparamResourceModel describes the resource data model.
type ExtendedmemoryparamResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Memlimit types.Int64  `tfsdk:"memlimit"`
}

func (r *ExtendedmemoryparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the extendedmemoryparam resource.",
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of NetScaler memory to reserve for the memory used by LSN and Subscriber Session Store feature, in multiples of 2MB.\n\nNote: If you later reduce the value of this parameter, the amount of active memory is not reduced. Changing the configured memory limit can only increase the amount of active memory.",
			},
		},
	}
}

func extendedmemoryparamGetThePayloadFromtheConfig(ctx context.Context, data *ExtendedmemoryparamResourceModel) basic.Extendedmemoryparam {
	tflog.Debug(ctx, "In extendedmemoryparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	extendedmemoryparam := basic.Extendedmemoryparam{}
	if !data.Memlimit.IsNull() {
		extendedmemoryparam.Memlimit = utils.IntPtr(int(data.Memlimit.ValueInt64()))
	}

	return extendedmemoryparam
}

func extendedmemoryparamSetAttrFromGet(ctx context.Context, data *ExtendedmemoryparamResourceModel, getResponseData map[string]interface{}) *ExtendedmemoryparamResourceModel {
	tflog.Debug(ctx, "In extendedmemoryparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["memlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Memlimit = types.Int64Value(intVal)
		}
	} else {
		data.Memlimit = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("extendedmemoryparam-config")

	return data
}
