package l4param

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// L4paramResourceModel describes the resource data model.
type L4paramResourceModel struct {
	Id           types.String `tfsdk:"id"`
	L2connmethod types.String `tfsdk:"l2connmethod"`
	L4switch     types.String `tfsdk:"l4switch"`
}

func (r *L4paramResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the l4param resource.",
			},
			"l2connmethod": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("MacVlanChannel"),
				Description: "Layer 2 connection method based on the combination of  channel number, MAC address and VLAN. It is tuned with l2conn param of lb vserver. If l2conn of lb vserver is ON then method specified here will be used to identify a connection in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>).",
			},
			"l4switch": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "In L4 switch topology, always clients and servers are on the same side. Enable l4switch to allow such connections.",
			},
		},
	}
}

func l4paramGetThePayloadFromtheConfig(ctx context.Context, data *L4paramResourceModel) network.L4param {
	tflog.Debug(ctx, "In l4paramGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	l4param := network.L4param{}
	if !data.L2connmethod.IsNull() {
		l4param.L2connmethod = data.L2connmethod.ValueString()
	}
	if !data.L4switch.IsNull() {
		l4param.L4switch = data.L4switch.ValueString()
	}

	return l4param
}

func l4paramSetAttrFromGet(ctx context.Context, data *L4paramResourceModel, getResponseData map[string]interface{}) *L4paramResourceModel {
	tflog.Debug(ctx, "In l4paramSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["l2connmethod"]; ok && val != nil {
		data.L2connmethod = types.StringValue(val.(string))
	} else {
		data.L2connmethod = types.StringNull()
	}
	if val, ok := getResponseData["l4switch"]; ok && val != nil {
		data.L4switch = types.StringValue(val.(string))
	} else {
		data.L4switch = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("l4param-config")

	return data
}
