package nstcpbufparam

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

// NstcpbufparamResourceModel describes the resource data model.
type NstcpbufparamResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Memlimit types.Int64  `tfsdk:"memlimit"`
	Size     types.Int64  `tfsdk:"size"`
}

func (r *NstcpbufparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstcpbufparam resource.",
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(64),
				Description: "Maximum memory, in megabytes, that can be used for buffering.",
			},
			"size": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(64),
				Description: "TCP buffering size per connection, in kilobytes.",
			},
		},
	}
}

func nstcpbufparamGetThePayloadFromtheConfig(ctx context.Context, data *NstcpbufparamResourceModel) ns.Nstcpbufparam {
	tflog.Debug(ctx, "In nstcpbufparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model.
	// Build the payload from configured values only (mirrors SDK v2 d.GetRawConfig()).
	// Guard against Unknown as well as Null so a Computed-but-unconfigured value
	// (which is Unknown in the plan) is never pushed as 0.
	nstcpbufparam := ns.Nstcpbufparam{}
	if !data.Memlimit.IsNull() && !data.Memlimit.IsUnknown() {
		nstcpbufparam.Memlimit = utils.IntPtr(int(data.Memlimit.ValueInt64()))
	}
	if !data.Size.IsNull() && !data.Size.IsUnknown() {
		nstcpbufparam.Size = utils.IntPtr(int(data.Size.ValueInt64()))
	}

	return nstcpbufparam
}

func nstcpbufparamSetAttrFromGet(ctx context.Context, data *NstcpbufparamResourceModel, getResponseData map[string]interface{}) *NstcpbufparamResourceModel {
	tflog.Debug(ctx, "In nstcpbufparamSetAttrFromGet Function")

	// Convert API response to model.
	// Only null the attribute in the else-branch when the current value is Unknown;
	// never clobber a known configured value that NITRO omits on default
	// (omit-on-default trap).
	if val, ok := getResponseData["memlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Memlimit = types.Int64Value(intVal)
		}
	} else if data.Memlimit.IsUnknown() {
		data.Memlimit = types.Int64Null()
	}
	if val, ok := getResponseData["size"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Size = types.Int64Value(intVal)
		}
	} else if data.Size.IsUnknown() {
		data.Size = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nstcpbufparam-config")

	return data
}
