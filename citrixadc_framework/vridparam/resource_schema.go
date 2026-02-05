package vridparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VridparamResourceModel describes the resource data model.
type VridparamResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Deadinterval  types.Int64  `tfsdk:"deadinterval"`
	Hellointerval types.Int64  `tfsdk:"hellointerval"`
	Sendtomaster  types.String `tfsdk:"sendtomaster"`
}

func (r *VridparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vridparam resource.",
			},
			"deadinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Number of seconds after which a peer node in active-active mode is marked down if vrrp advertisements are not received from the peer node.",
			},
			"hellointerval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1000),
				Description: "Interval, in milliseconds, between vrrp advertisement messages sent to the peer node in active-active mode.",
			},
			"sendtomaster": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Forward packets to the master node, in an active-active mode configuration, if the virtual server is in the backup state and sharing is disabled.",
			},
		},
	}
}

func vridparamGetThePayloadFromtheConfig(ctx context.Context, data *VridparamResourceModel) network.Vridparam {
	tflog.Debug(ctx, "In vridparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vridparam := network.Vridparam{}
	if !data.Deadinterval.IsNull() {
		vridparam.Deadinterval = utils.IntPtr(int(data.Deadinterval.ValueInt64()))
	}
	if !data.Hellointerval.IsNull() {
		vridparam.Hellointerval = utils.IntPtr(int(data.Hellointerval.ValueInt64()))
	}
	if !data.Sendtomaster.IsNull() {
		vridparam.Sendtomaster = data.Sendtomaster.ValueString()
	}

	return vridparam
}

func vridparamSetAttrFromGet(ctx context.Context, data *VridparamResourceModel, getResponseData map[string]interface{}) *VridparamResourceModel {
	tflog.Debug(ctx, "In vridparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["deadinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Deadinterval = types.Int64Value(intVal)
		}
	} else {
		data.Deadinterval = types.Int64Null()
	}
	if val, ok := getResponseData["hellointerval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hellointerval = types.Int64Value(intVal)
		}
	} else {
		data.Hellointerval = types.Int64Null()
	}
	if val, ok := getResponseData["sendtomaster"]; ok && val != nil {
		data.Sendtomaster = types.StringValue(val.(string))
	} else {
		data.Sendtomaster = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("vridparam-config")

	return data
}
