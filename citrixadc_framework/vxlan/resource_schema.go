package vxlan

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VxlanResourceModel describes the resource data model.
type VxlanResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Dynamicrouting     types.String `tfsdk:"dynamicrouting"`
	Vxlanid            types.Int64  `tfsdk:"vxlanid"`
	Innervlantagging   types.String `tfsdk:"innervlantagging"`
	Ipv6dynamicrouting types.String `tfsdk:"ipv6dynamicrouting"`
	Port               types.Int64  `tfsdk:"port"`
	Protocol           types.String `tfsdk:"protocol"`
	Type               types.String `tfsdk:"type"`
	Vlan               types.Int64  `tfsdk:"vlan"`
}

func (r *VxlanResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vxlan resource.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable dynamic routing on this VXLAN.",
			},
			"vxlanid": schema.Int64Attribute{
				Required:    true,
				Description: "A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.",
			},
			"innervlantagging": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Specifies whether Citrix ADC should generate VXLAN packets with inner VLAN tag.",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable all IPv6 dynamic routing protocols on this VXLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4789),
				Description: "Specifies UDP destination port for VXLAN packets.",
			},
			"protocol": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ETHERNET"),
				Description: "VXLAN-GPE next protocol. RESERVED, IPv4, IPv6, ETHERNET, NSH",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("VXLAN"),
				Description: "VXLAN encapsulation type. VXLAN, VXLANGPE",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of VLANs whose traffic is allowed over this VXLAN. If you do not specify any VLAN IDs, the Citrix ADC allows traffic of all VLANs that are not part of any other VXLANs.",
			},
		},
	}
}

func vxlanGetThePayloadFromtheConfig(ctx context.Context, data *VxlanResourceModel) network.Vxlan {
	tflog.Debug(ctx, "In vxlanGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vxlan := network.Vxlan{}
	if !data.Dynamicrouting.IsNull() {
		vxlan.Dynamicrouting = data.Dynamicrouting.ValueString()
	}
	if !data.Vxlanid.IsNull() {
		vxlan.Id = utils.IntPtr(int(data.Vxlanid.ValueInt64()))
	}
	if !data.Innervlantagging.IsNull() {
		vxlan.Innervlantagging = data.Innervlantagging.ValueString()
	}
	if !data.Ipv6dynamicrouting.IsNull() {
		vxlan.Ipv6dynamicrouting = data.Ipv6dynamicrouting.ValueString()
	}
	if !data.Port.IsNull() {
		vxlan.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Protocol.IsNull() {
		vxlan.Protocol = data.Protocol.ValueString()
	}
	if !data.Type.IsNull() {
		vxlan.Type = data.Type.ValueString()
	}
	if !data.Vlan.IsNull() {
		vxlan.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return vxlan
}

func vxlanSetAttrFromGet(ctx context.Context, data *VxlanResourceModel, getResponseData map[string]interface{}) *VxlanResourceModel {
	tflog.Debug(ctx, "In vxlanSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["dynamicrouting"]; ok && val != nil {
		data.Dynamicrouting = types.StringValue(val.(string))
	} else {
		data.Dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlanid = types.Int64Value(intVal)
		}
	} else {
		data.Vxlanid = types.Int64Null()
	}
	if val, ok := getResponseData["innervlantagging"]; ok && val != nil {
		data.Innervlantagging = types.StringValue(val.(string))
	} else {
		data.Innervlantagging = types.StringNull()
	}
	if val, ok := getResponseData["ipv6dynamicrouting"]; ok && val != nil {
		data.Ipv6dynamicrouting = types.StringValue(val.(string))
	} else {
		data.Ipv6dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		data.Protocol = types.StringValue(val.(string))
	} else {
		data.Protocol = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Vxlanid.ValueInt64()))

	return data
}
