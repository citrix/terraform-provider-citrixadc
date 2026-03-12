package arp

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ArpResourceModel describes the resource data model.
type ArpResourceModel struct {
	Id        types.String `tfsdk:"id"`
	All       types.Bool   `tfsdk:"all"`
	Ifnum     types.String `tfsdk:"ifnum"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Mac       types.String `tfsdk:"mac"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Ownernode types.Int64  `tfsdk:"ownernode"`
	Td        types.Int64  `tfsdk:"td"`
	Vlan      types.Int64  `tfsdk:"vlan"`
	Vtep      types.String `tfsdk:"vtep"`
	Vxlan     types.Int64  `tfsdk:"vxlan"`
}

func (r *ArpResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the arp resource.",
			},
			"all": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Remove all ARP entries from the ARP table of the Citrix ADC.",
			},
			"ifnum": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interface through which the network device is accessible. Specify the interface in (slot/port) notation. For example, 1/3.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address of the network device that you want to add to the ARP table.",
			},
			"mac": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "MAC address of the network device.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"ownernode": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The owner node for the Arp entry.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The VLAN ID through which packets are to be sent after matching the ARP entry. This is a numeric value.",
			},
			"vtep": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address of the VXLAN tunnel endpoint (VTEP) through which the IP address of this ARP entry is reachable.",
			},
			"vxlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the VXLAN on which the IP address of this ARP entry is reachable.",
			},
		},
	}
}

func arpGetThePayloadFromtheConfig(ctx context.Context, data *ArpResourceModel) network.Arp {
	tflog.Debug(ctx, "In arpGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	arp := network.Arp{}
	if !data.All.IsNull() {
		arp.All = data.All.ValueBool()
	}
	if !data.Ifnum.IsNull() {
		arp.Ifnum = data.Ifnum.ValueString()
	}
	if !data.Ipaddress.IsNull() {
		arp.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Mac.IsNull() {
		arp.Mac = data.Mac.ValueString()
	}
	if !data.Nodeid.IsNull() {
		arp.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Ownernode.IsNull() {
		arp.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Td.IsNull() {
		arp.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vlan.IsNull() {
		arp.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vtep.IsNull() {
		arp.Vtep = data.Vtep.ValueString()
	}
	if !data.Vxlan.IsNull() {
		arp.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}

	return arp
}

func arpSetAttrFromGet(ctx context.Context, data *ArpResourceModel, getResponseData map[string]interface{}) *ArpResourceModel {
	tflog.Debug(ctx, "In arpSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["all"]; ok && val != nil {
		data.All = types.BoolValue(val.(bool))
	} else {
		data.All = types.BoolNull()
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	} else {
		data.Ifnum = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["mac"]; ok && val != nil {
		data.Mac = types.StringValue(val.(string))
	} else {
		data.Mac = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else {
		data.Ownernode = types.Int64Null()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vtep"]; ok && val != nil {
		data.Vtep = types.StringValue(val.(string))
	} else {
		data.Vtep = types.StringNull()
	}
	if val, ok := getResponseData["vxlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlan = types.Int64Value(intVal)
		}
	} else {
		data.Vxlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%d,%d", data.Ipaddress.ValueString(), data.Ownernode.ValueInt64(), data.Td.ValueInt64()))

	return data
}
