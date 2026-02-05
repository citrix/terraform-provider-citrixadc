package netbridge

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NetbridgeResourceModel describes the resource data model.
type NetbridgeResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Vxlanvlanmap types.String `tfsdk:"vxlanvlanmap"`
}

func (r *NetbridgeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the netbridge resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the network bridge.",
			},
			"vxlanvlanmap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The vlan to vxlan mapping to be applied to this netbridge.",
			},
		},
	}
}

func netbridgeGetThePayloadFromtheConfig(ctx context.Context, data *NetbridgeResourceModel) network.Netbridge {
	tflog.Debug(ctx, "In netbridgeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	netbridge := network.Netbridge{}
	if !data.Name.IsNull() {
		netbridge.Name = data.Name.ValueString()
	}
	if !data.Vxlanvlanmap.IsNull() {
		netbridge.Vxlanvlanmap = data.Vxlanvlanmap.ValueString()
	}

	return netbridge
}

func netbridgeSetAttrFromGet(ctx context.Context, data *NetbridgeResourceModel, getResponseData map[string]interface{}) *NetbridgeResourceModel {
	tflog.Debug(ctx, "In netbridgeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["vxlanvlanmap"]; ok && val != nil {
		data.Vxlanvlanmap = types.StringValue(val.(string))
	} else {
		data.Vxlanvlanmap = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
