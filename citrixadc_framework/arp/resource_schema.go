package arp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
			// SDK v2: Optional+Computed, NOT ForceNew -> no RequiresReplace.
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove all ARP entries from the ARP table of the Citrix ADC.",
			},
			"ifnum": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
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
					// GH #1436
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"ownernode": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// GH #1436
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The owner node for the Arp entry.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// GH #1436
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// GH #1436
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The VLAN ID through which packets are to be sent after matching the ARP entry. This is a numeric value.",
			},
			"vtep": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "IP address of the VXLAN tunnel endpoint (VTEP) through which the IP address of this ARP entry is reachable.",
			},
			"vxlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// GH #1436
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "ID of the VXLAN on which the IP address of this ARP entry is reachable.",
			},
		},
	}
}

// arpGetThePayloadFromthePlan builds the NITRO add payload, mirroring the SDK v2
// createArpFunc exactly: only ipaddress, mac, ifnum, td, vlan, vtep, vxlan and
// ownernode are sent (nodeid and all are not part of the NITRO "add" operation).
func arpGetThePayloadFromthePlan(ctx context.Context, data *ArpResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In arpGetThePayloadFromthePlan Function")

	arp := make(map[string]interface{})
	arp["ipaddress"] = data.Ipaddress.ValueString()
	arp["mac"] = data.Mac.ValueString()
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		arp["ifnum"] = data.Ifnum.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		arp["td"] = int(data.Td.ValueInt64())
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		arp["vlan"] = int(data.Vlan.ValueInt64())
	}
	if !data.Vtep.IsNull() && !data.Vtep.IsUnknown() {
		arp["vtep"] = data.Vtep.ValueString()
	}
	if !data.Vxlan.IsNull() && !data.Vxlan.IsUnknown() {
		arp["vxlan"] = int(data.Vxlan.ValueInt64())
	}
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		arp["ownernode"] = int(data.Ownernode.ValueInt64())
	}

	return arp
}

// arpSetAttrFromGet maps the NITRO GET response onto the resource model.
// Mirrors the SDK v2 readArpFunc: it intentionally does NOT overwrite "mac"
// (the appliance normalises the MAC to lower-case, which would otherwise cause
// a spurious diff against the user-supplied value) and it sets the ID to the
// plain ipaddress value (matching SDK v2 d.SetId(arpName)).
func arpSetAttrFromGet(ctx context.Context, data *ArpResourceModel, getResponseData map[string]interface{}) *ArpResourceModel {
	tflog.Debug(ctx, "In arpSetAttrFromGet Function")

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
	// mac is intentionally NOT set from the API response (SDK v2 parity) -
	// preserve the value already present in plan/state.
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

	// Set ID for the resource - single key attribute (ipaddress), plain value.
	data.Id = types.StringValue(data.Ipaddress.ValueString())

	return data
}

// arpSetAttrFromGetForDatasource maps the NITRO GET response onto the model for
// the data source. Unlike the resource variant it DOES copy "mac" from the API
// (data sources surface the appliance's normalised value) and sets the ID.
func arpSetAttrFromGetForDatasource(ctx context.Context, data *ArpResourceModel, getResponseData map[string]interface{}) *ArpResourceModel {
	tflog.Debug(ctx, "In arpSetAttrFromGetForDatasource Function")

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

	// Set ID for the resource - single key attribute (ipaddress), plain value.
	data.Id = types.StringValue(data.Ipaddress.ValueString())

	return data
}
