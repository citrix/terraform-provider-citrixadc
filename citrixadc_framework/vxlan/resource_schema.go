package vxlan

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
			// SDK v2: vxlanid TypeInt Required ForceNew -> Required + RequiresReplace.
			"vxlanid": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.",
			},
			// SDK v2: Optional+Computed, no Default (value read from ADC).
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable dynamic routing on this VXLAN.",
			},
			"innervlantagging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether Citrix ADC should generate VXLAN packets with inner VLAN tag.",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable all IPv6 dynamic routing protocols on this VXLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies UDP destination port for VXLAN packets.",
			},
			// SDK v2: protocol Optional+Computed, NOT ForceNew (Update handles it).
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "VXLAN-GPE next protocol. RESERVED, IPv4, IPv6, ETHERNET, NSH",
			},
			// SDK v2: type Optional+Computed, NOT ForceNew (Update handles it).
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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

// vxlanGetThePayloadFromthePlan builds the full create payload from the plan.
func vxlanGetThePayloadFromthePlan(ctx context.Context, data *VxlanResourceModel) network.Vxlan {
	tflog.Debug(ctx, "In vxlanGetThePayloadFromthePlan Function")

	vxlan := network.Vxlan{}
	if !data.Vxlanid.IsNull() && !data.Vxlanid.IsUnknown() {
		vxlan.Id = utils.IntPtr(int(data.Vxlanid.ValueInt64()))
	}
	if !data.Dynamicrouting.IsNull() && !data.Dynamicrouting.IsUnknown() {
		vxlan.Dynamicrouting = data.Dynamicrouting.ValueString()
	}
	if !data.Innervlantagging.IsNull() && !data.Innervlantagging.IsUnknown() {
		vxlan.Innervlantagging = data.Innervlantagging.ValueString()
	}
	if !data.Ipv6dynamicrouting.IsNull() && !data.Ipv6dynamicrouting.IsUnknown() {
		vxlan.Ipv6dynamicrouting = data.Ipv6dynamicrouting.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		vxlan.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		vxlan.Protocol = data.Protocol.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		vxlan.Type = data.Type.ValueString()
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		vxlan.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return vxlan
}

// vxlanGetTheUpdatablePayloadFromThePlan mirrors SDK v2 update semantics: only
// changed fields are placed on the payload. The key (id) is always included.
func vxlanGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *VxlanResourceModel, state *VxlanResourceModel) (network.Vxlan, bool) {
	tflog.Debug(ctx, "In vxlanGetTheUpdatablePayloadFromThePlan Function")

	vxlan := network.Vxlan{}
	if !data.Vxlanid.IsNull() && !data.Vxlanid.IsUnknown() {
		vxlan.Id = utils.IntPtr(int(data.Vxlanid.ValueInt64()))
	}

	hasChange := false
	if !data.Dynamicrouting.Equal(state.Dynamicrouting) {
		vxlan.Dynamicrouting = data.Dynamicrouting.ValueString()
		hasChange = true
	}
	if !data.Innervlantagging.Equal(state.Innervlantagging) {
		vxlan.Innervlantagging = data.Innervlantagging.ValueString()
		hasChange = true
	}
	if !data.Ipv6dynamicrouting.Equal(state.Ipv6dynamicrouting) {
		vxlan.Ipv6dynamicrouting = data.Ipv6dynamicrouting.ValueString()
		hasChange = true
	}
	if !data.Port.Equal(state.Port) {
		vxlan.Port = utils.IntPtr(int(data.Port.ValueInt64()))
		hasChange = true
	}
	if !data.Protocol.Equal(state.Protocol) {
		vxlan.Protocol = data.Protocol.ValueString()
		hasChange = true
	}
	if !data.Type.Equal(state.Type) {
		vxlan.Type = data.Type.ValueString()
		hasChange = true
	}
	if !data.Vlan.Equal(state.Vlan) {
		vxlan.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
		hasChange = true
	}

	return vxlan, hasChange
}

func vxlanSetAttrFromGet(ctx context.Context, data *VxlanResourceModel, getResponseData map[string]interface{}) *VxlanResourceModel {
	tflog.Debug(ctx, "In vxlanSetAttrFromGet Function")

	// Convert API response to model. else-branches only null when the current
	// value is Unknown, so a configured value NITRO omits from GET is preserved
	// (omit-on-default guard).
	if val, ok := getResponseData["dynamicrouting"]; ok && val != nil {
		data.Dynamicrouting = types.StringValue(val.(string))
	} else if data.Dynamicrouting.IsUnknown() {
		data.Dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlanid = types.Int64Value(intVal)
		}
	} else if data.Vxlanid.IsUnknown() {
		data.Vxlanid = types.Int64Null()
	}
	if val, ok := getResponseData["innervlantagging"]; ok && val != nil {
		data.Innervlantagging = types.StringValue(val.(string))
	} else if data.Innervlantagging.IsUnknown() {
		data.Innervlantagging = types.StringNull()
	}
	if val, ok := getResponseData["ipv6dynamicrouting"]; ok && val != nil {
		data.Ipv6dynamicrouting = types.StringValue(val.(string))
	} else if data.Ipv6dynamicrouting.IsUnknown() {
		data.Ipv6dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else if data.Port.IsUnknown() {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		data.Protocol = types.StringValue(val.(string))
	} else if data.Protocol.IsUnknown() {
		data.Protocol = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else if data.Vlan.IsUnknown() {
		data.Vlan = types.Int64Null()
	}

	// Set ID for the resource (single unique attribute -> plain value), matching
	// the SDK v2 d.SetId(vxlanIdStr) scheme.
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Vxlanid.ValueInt64()))

	return data
}
