package vlan

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

// VlanResourceModel describes the resource data model.
type VlanResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Aliasname          types.String `tfsdk:"aliasname"`
	Dynamicrouting     types.String `tfsdk:"dynamicrouting"`
	Vlanid             types.Int64  `tfsdk:"vlanid"`
	Ipv6dynamicrouting types.String `tfsdk:"ipv6dynamicrouting"`
	Mtu                types.Int64  `tfsdk:"mtu"`
	Sharing            types.String `tfsdk:"sharing"`
}

func (r *VlanResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vlan resource.",
			},
			"aliasname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A name for the VLAN. Must begin with a letter, a number, or the underscore symbol, and can consist of from 1 to 31 letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a name that helps identify the VLAN. However, you cannot perform any VLAN operation by specifying this name instead of the VLAN ID.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable dynamic routing on this VLAN.",
			},
			"vlanid": schema.Int64Attribute{
				Required:    true,
				Description: "A positive integer that uniquely identifies a VLAN.",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable all IPv6 dynamic routing protocols on this VLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.",
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the maximum transmission unit (MTU), in bytes. The MTU is the largest packet size, excluding 14 bytes of ethernet header and 4 bytes of crc, that can be transmitted and received over this VLAN.",
			},
			"sharing": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If sharing is enabled, then this vlan can be shared across multiple partitions by binding it to all those partitions. If sharing is disabled, then this vlan can be bound to only one of the partitions.",
			},
		},
	}
}

func vlanGetThePayloadFromtheConfig(ctx context.Context, data *VlanResourceModel) network.Vlan {
	tflog.Debug(ctx, "In vlanGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vlan := network.Vlan{}
	if !data.Aliasname.IsNull() {
		vlan.Aliasname = data.Aliasname.ValueString()
	}
	if !data.Dynamicrouting.IsNull() {
		vlan.Dynamicrouting = data.Dynamicrouting.ValueString()
	}
	if !data.Vlanid.IsNull() {
		vlan.Id = utils.IntPtr(int(data.Vlanid.ValueInt64()))
	}
	if !data.Ipv6dynamicrouting.IsNull() {
		vlan.Ipv6dynamicrouting = data.Ipv6dynamicrouting.ValueString()
	}
	if !data.Mtu.IsNull() {
		vlan.Mtu = utils.IntPtr(int(data.Mtu.ValueInt64()))
	}
	if !data.Sharing.IsNull() {
		vlan.Sharing = data.Sharing.ValueString()
	}

	return vlan
}

func vlanSetAttrFromGet(ctx context.Context, data *VlanResourceModel, getResponseData map[string]interface{}) *VlanResourceModel {
	tflog.Debug(ctx, "In vlanSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["aliasname"]; ok && val != nil {
		data.Aliasname = types.StringValue(val.(string))
	} else {
		data.Aliasname = types.StringNull()
	}
	if val, ok := getResponseData["dynamicrouting"]; ok && val != nil {
		data.Dynamicrouting = types.StringValue(val.(string))
	} else {
		data.Dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlanid = types.Int64Value(intVal)
		}
	} else {
		data.Vlanid = types.Int64Null()
	}
	if val, ok := getResponseData["ipv6dynamicrouting"]; ok && val != nil {
		data.Ipv6dynamicrouting = types.StringValue(val.(string))
	} else {
		data.Ipv6dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["mtu"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mtu = types.Int64Value(intVal)
		}
	} else {
		data.Mtu = types.Int64Null()
	}
	if val, ok := getResponseData["sharing"]; ok && val != nil {
		data.Sharing = types.StringValue(val.(string))
	} else {
		data.Sharing = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Vlanid.ValueInt64()))
	return data
}
