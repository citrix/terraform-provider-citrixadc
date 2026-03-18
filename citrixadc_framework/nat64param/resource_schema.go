package nat64param

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Nat64paramResourceModel describes the resource data model.
type Nat64paramResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Nat64fragheader   types.String `tfsdk:"nat64fragheader"`
	Nat64ignoretos    types.String `tfsdk:"nat64ignoretos"`
	Nat64v6mtu        types.Int64  `tfsdk:"nat64v6mtu"`
	Nat64zerochecksum types.String `tfsdk:"nat64zerochecksum"`
	Td                types.Int64  `tfsdk:"td"`
}

func (r *Nat64paramResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nat64param resource.",
			},
			"nat64fragheader": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "When disabled, translator will not insert IPv6 fragmentation header for non fragmented IPv4 packets",
			},
			"nat64ignoretos": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore TOS.",
			},
			"nat64v6mtu": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1280),
				Description: "MTU setting for the IPv6 side. If the incoming IPv4 packet greater than this, either fragment or send icmp need fragmentation error.",
			},
			"nat64zerochecksum": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Calculate checksum for UDP packets with zero checksum",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func nat64paramGetThePayloadFromtheConfig(ctx context.Context, data *Nat64paramResourceModel) network.Nat64param {
	tflog.Debug(ctx, "In nat64paramGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nat64param := network.Nat64param{}
	if !data.Nat64fragheader.IsNull() {
		nat64param.Nat64fragheader = data.Nat64fragheader.ValueString()
	}
	if !data.Nat64ignoretos.IsNull() {
		nat64param.Nat64ignoretos = data.Nat64ignoretos.ValueString()
	}
	if !data.Nat64v6mtu.IsNull() {
		nat64param.Nat64v6mtu = utils.IntPtr(int(data.Nat64v6mtu.ValueInt64()))
	}
	if !data.Nat64zerochecksum.IsNull() {
		nat64param.Nat64zerochecksum = data.Nat64zerochecksum.ValueString()
	}
	if !data.Td.IsNull() {
		nat64param.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return nat64param
}

func nat64paramSetAttrFromGet(ctx context.Context, data *Nat64paramResourceModel, getResponseData map[string]interface{}) *Nat64paramResourceModel {
	tflog.Debug(ctx, "In nat64paramSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["nat64fragheader"]; ok && val != nil {
		data.Nat64fragheader = types.StringValue(val.(string))
	} else {
		data.Nat64fragheader = types.StringNull()
	}
	if val, ok := getResponseData["nat64ignoretos"]; ok && val != nil {
		data.Nat64ignoretos = types.StringValue(val.(string))
	} else {
		data.Nat64ignoretos = types.StringNull()
	}
	if val, ok := getResponseData["nat64v6mtu"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nat64v6mtu = types.Int64Value(intVal)
		}
	} else {
		data.Nat64v6mtu = types.Int64Null()
	}
	if val, ok := getResponseData["nat64zerochecksum"]; ok && val != nil {
		data.Nat64zerochecksum = types.StringValue(val.(string))
	} else {
		data.Nat64zerochecksum = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Td.ValueInt64()))

	return data
}
