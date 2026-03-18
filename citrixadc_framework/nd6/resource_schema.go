package nd6

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Nd6ResourceModel describes the resource data model.
type Nd6ResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Ifnum    types.String `tfsdk:"ifnum"`
	Mac      types.String `tfsdk:"mac"`
	Neighbor types.String `tfsdk:"neighbor"`
	Nodeid   types.Int64  `tfsdk:"nodeid"`
	Td       types.Int64  `tfsdk:"td"`
	Vlan     types.Int64  `tfsdk:"vlan"`
	Vtep     types.String `tfsdk:"vtep"`
	Vxlan    types.Int64  `tfsdk:"vxlan"`
}

func (r *Nd6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nd6 resource.",
			},
			"ifnum": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interface through which the adjacent network device is available, specified in slot/port notation (for example, 1/3). Use spaces to separate multiple entries.",
			},
			"mac": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "MAC address of the adjacent network device.",
			},
			"neighbor": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Link-local IPv6 address of the adjacent network device to add to the ND6 table.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
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
				Description: "Integer value that uniquely identifies the VLAN on which the adjacent network device exists.",
			},
			"vtep": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address of the VXLAN tunnel endpoint (VTEP) through which the IPv6 address of this ND6 entry is reachable.",
			},
			"vxlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the VXLAN on which the IPv6 address of this ND6 entry is reachable.",
			},
		},
	}
}

func nd6GetThePayloadFromtheConfig(ctx context.Context, data *Nd6ResourceModel) network.Nd6 {
	tflog.Debug(ctx, "In nd6GetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nd6 := network.Nd6{}
	if !data.Ifnum.IsNull() {
		nd6.Ifnum = data.Ifnum.ValueString()
	}
	if !data.Mac.IsNull() {
		nd6.Mac = data.Mac.ValueString()
	}
	if !data.Neighbor.IsNull() {
		nd6.Neighbor = data.Neighbor.ValueString()
	}
	// Only send nodeid and td if they are explicitly set (not null and not unknown)
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		nd6.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		nd6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vlan.IsNull() {
		nd6.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vtep.IsNull() {
		nd6.Vtep = data.Vtep.ValueString()
	}
	if !data.Vxlan.IsNull() {
		nd6.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}

	return nd6
}

func nd6SetAttrFromGet(ctx context.Context, data *Nd6ResourceModel, getResponseData map[string]interface{}) *Nd6ResourceModel {
	tflog.Debug(ctx, "In nd6SetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	} else {
		data.Ifnum = types.StringNull()
	}
	if val, ok := getResponseData["mac"]; ok && val != nil {
		data.Mac = types.StringValue(val.(string))
	} else {
		data.Mac = types.StringNull()
	}
	if val, ok := getResponseData["neighbor"]; ok && val != nil {
		data.Neighbor = types.StringValue(val.(string))
	} else {
		data.Neighbor = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
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

	// Set ID for the resource - just use neighbor as the primary key
	data.Id = data.Neighbor

	return data
}
