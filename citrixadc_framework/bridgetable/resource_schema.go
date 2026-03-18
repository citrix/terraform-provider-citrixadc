package bridgetable

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BridgetableResourceModel describes the resource data model.
type BridgetableResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Bridgeage  types.Int64  `tfsdk:"bridgeage"`
	Devicevlan types.Int64  `tfsdk:"devicevlan"`
	Ifnum      types.String `tfsdk:"ifnum"`
	Mac        types.String `tfsdk:"mac"`
	Nodeid     types.Int64  `tfsdk:"nodeid"`
	Vlan       types.Int64  `tfsdk:"vlan"`
	Vni        types.Int64  `tfsdk:"vni"`
	Vtep       types.String `tfsdk:"vtep"`
	Vxlan      types.Int64  `tfsdk:"vxlan"`
}

func (r *BridgetableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the bridgetable resource.",
			},
			"bridgeage": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(300),
				Description: "Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value.",
			},
			"devicevlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The vlan on which to send multicast packets when the VXLAN tunnel endpoint is a muticast group address.",
			},
			"ifnum": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "INTERFACE  whose entries are to be removed.",
			},
			"mac": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The MAC address of the target.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"vlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "VLAN  whose entries are to be removed.",
			},
			"vni": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The VXLAN VNI Network Identifier (or VXLAN Segment ID) to use to connect to the remote VXLAN tunnel endpoint.  If omitted the value specified as vxlan will be used.",
			},
			"vtep": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address of the destination VXLAN tunnel endpoint where the Ethernet MAC ADDRESS resides.",
			},
			"vxlan": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The VXLAN to which this address is associated.",
			},
		},
	}
}

func bridgetableGetThePayloadFromtheConfig(ctx context.Context, data *BridgetableResourceModel) network.Bridgetable {
	tflog.Debug(ctx, "In bridgetableGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	bridgetable := network.Bridgetable{}
	if !data.Bridgeage.IsNull() {
		bridgetable.Bridgeage = utils.IntPtr(int(data.Bridgeage.ValueInt64()))
	}
	if !data.Devicevlan.IsNull() {
		bridgetable.Devicevlan = utils.IntPtr(int(data.Devicevlan.ValueInt64()))
	}
	if !data.Ifnum.IsNull() {
		bridgetable.Ifnum = data.Ifnum.ValueString()
	}
	if !data.Mac.IsNull() {
		bridgetable.Mac = data.Mac.ValueString()
	}
	if !data.Nodeid.IsNull() {
		bridgetable.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Vlan.IsNull() {
		bridgetable.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vni.IsNull() {
		bridgetable.Vni = utils.IntPtr(int(data.Vni.ValueInt64()))
	}
	if !data.Vtep.IsNull() {
		bridgetable.Vtep = data.Vtep.ValueString()
	}
	if !data.Vxlan.IsNull() {
		bridgetable.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}

	return bridgetable
}

func bridgetableSetAttrFromGet(ctx context.Context, data *BridgetableResourceModel, getResponseData map[string]interface{}) *BridgetableResourceModel {
	tflog.Debug(ctx, "In bridgetableSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bridgeage"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bridgeage = types.Int64Value(intVal)
		}
	} else {
		data.Bridgeage = types.Int64Null()
	}
	if val, ok := getResponseData["devicevlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Devicevlan = types.Int64Value(intVal)
		}
	} else {
		data.Devicevlan = types.Int64Null()
	}
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
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vni"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vni = types.Int64Value(intVal)
		}
	} else {
		data.Vni = types.Int64Null()
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
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("bridgetable-config")

	return data
}
