package lacp

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LacpResourceModel describes the resource data model.
type LacpResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Ownernode   types.Int64  `tfsdk:"ownernode"`
	Syspriority types.Int64  `tfsdk:"syspriority"`
}

func (r *LacpResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lacp resource.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(255),
				Description: "The owner node in a cluster for which we want to set the lacp priority. Owner node can vary from 0 to 31. Ownernode value of 254 is used for Cluster.",
			},
			"syspriority": schema.Int64Attribute{
				Required:    true,
				Default:     int64default.StaticInt64(32768),
				Description: "Priority number that determines which peer device of an LACP LA channel can have control over the LA channel. This parameter is globally applied to all LACP channels on the Citrix ADC. The lower the number, the higher the priority.",
			},
		},
	}
}

func lacpGetThePayloadFromtheConfig(ctx context.Context, data *LacpResourceModel) network.Lacp {
	tflog.Debug(ctx, "In lacpGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lacp := network.Lacp{}
	if !data.Ownernode.IsNull() {
		lacp.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Syspriority.IsNull() {
		lacp.Syspriority = utils.IntPtr(int(data.Syspriority.ValueInt64()))
	}

	return lacp
}

func lacpSetAttrFromGet(ctx context.Context, data *LacpResourceModel, getResponseData map[string]interface{}) *LacpResourceModel {
	tflog.Debug(ctx, "In lacpSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else {
		data.Ownernode = types.Int64Null()
	}
	if val, ok := getResponseData["syspriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Syspriority = types.Int64Value(intVal)
		}
	} else {
		data.Syspriority = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))

	return data
}
