package bridgegroup

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BridgegroupResourceModel describes the resource data model.
type BridgegroupResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Dynamicrouting     types.String `tfsdk:"dynamicrouting"`
	Bridgegroupid      types.Int64  `tfsdk:"bridgegroup_id"`
	Ipv6dynamicrouting types.String `tfsdk:"ipv6dynamicrouting"`
}

func (r *BridgegroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the bridgegroup resource.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable dynamic routing for this bridgegroup.",
			},
			"bridgegroup_id": schema.Int64Attribute{
				Required:    true,
				Description: "An integer that uniquely identifies the bridge group.",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable all IPv6 dynamic routing protocols on all VLANs bound to this bridgegroup. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.",
			},
		},
	}
}

func bridgegroupGetThePayloadFromtheConfig(ctx context.Context, data *BridgegroupResourceModel) network.Bridgegroup {
	tflog.Debug(ctx, "In bridgegroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	bridgegroup := network.Bridgegroup{}
	if !data.Dynamicrouting.IsNull() {
		bridgegroup.Dynamicrouting = data.Dynamicrouting.ValueString()
	}
	if !data.Bridgegroupid.IsNull() {
		bridgegroup.Id = utils.IntPtr(int(data.Bridgegroupid.ValueInt64()))
	}
	if !data.Ipv6dynamicrouting.IsNull() {
		bridgegroup.Ipv6dynamicrouting = data.Ipv6dynamicrouting.ValueString()
	}

	return bridgegroup
}

func bridgegroupSetAttrFromGet(ctx context.Context, data *BridgegroupResourceModel, getResponseData map[string]interface{}) *BridgegroupResourceModel {
	tflog.Debug(ctx, "In bridgegroupSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["dynamicrouting"]; ok && val != nil {
		data.Dynamicrouting = types.StringValue(val.(string))
	} else {
		data.Dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bridgegroupid = types.Int64Value(intVal)
		}
	} else {
		data.Bridgegroupid = types.Int64Null()
	}
	if val, ok := getResponseData["ipv6dynamicrouting"]; ok && val != nil {
		data.Ipv6dynamicrouting = types.StringValue(val.(string))
	} else {
		data.Ipv6dynamicrouting = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Bridgegroupid.ValueInt64()))

	return data
}
