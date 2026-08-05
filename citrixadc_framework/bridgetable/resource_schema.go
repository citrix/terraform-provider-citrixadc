package bridgetable

import (
	"context"
	"fmt"

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
				// Optional+Computed, no Default. bridgeage is a table-wide
				// setting whose default (300) is supplied by the ADC. A schema
				// Default without Computed is invalid; SDK v2 kept this
				// Optional+Computed with no Default.
				Optional:    true,
				Computed:    true,
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

// bridgetableGetThePayloadFromthePlan builds the per-entry "add" payload.
// bridgeage is intentionally excluded here: it is a table-wide setting applied
// via a separate UpdateUnnamedResource call (matches SDK v2 behavior and the
// NITRO "add" payload which does not include bridgeage).
func bridgetableGetThePayloadFromthePlan(ctx context.Context, data *BridgetableResourceModel) network.Bridgetable {
	tflog.Debug(ctx, "In bridgetableGetThePayloadFromthePlan Function")

	bridgetable := network.Bridgetable{}
	if !data.Devicevlan.IsNull() && !data.Devicevlan.IsUnknown() {
		bridgetable.Devicevlan = utils.IntPtr(int(data.Devicevlan.ValueInt64()))
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		bridgetable.Ifnum = data.Ifnum.ValueString()
	}
	if !data.Mac.IsNull() && !data.Mac.IsUnknown() {
		bridgetable.Mac = data.Mac.ValueString()
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		bridgetable.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		bridgetable.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vni.IsNull() && !data.Vni.IsUnknown() {
		bridgetable.Vni = utils.IntPtr(int(data.Vni.ValueInt64()))
	}
	if !data.Vtep.IsNull() && !data.Vtep.IsUnknown() {
		bridgetable.Vtep = data.Vtep.ValueString()
	}
	if !data.Vxlan.IsNull() && !data.Vxlan.IsUnknown() {
		bridgetable.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}

	return bridgetable
}

// bridgetableGetTheBridgeagePayload builds the table-wide bridgeage-only payload
// used for the create-time follow-up update and for in-place updates.
func bridgetableGetTheBridgeagePayload(ctx context.Context, data *BridgetableResourceModel) network.Bridgetable {
	tflog.Debug(ctx, "In bridgetableGetTheBridgeagePayload Function")

	bridgetable := network.Bridgetable{}
	if !data.Bridgeage.IsNull() && !data.Bridgeage.IsUnknown() {
		bridgetable.Bridgeage = utils.IntPtr(int(data.Bridgeage.ValueInt64()))
	}
	return bridgetable
}

// bridgetableSetAttrFromGet maps a GET response entry onto the resource model,
// preserving configured values that NITRO omits from GET (prevents
// "inconsistent result after apply" on Optional+Computed attributes).
func bridgetableSetAttrFromGet(ctx context.Context, data *BridgetableResourceModel, getResponseData map[string]interface{}) *BridgetableResourceModel {
	tflog.Debug(ctx, "In bridgetableSetAttrFromGet Function")

	// bridgeage is a table-wide setting; NITRO GET does not reliably reflect the
	// per-entry configured value, so preserve the configured/state value (matches
	// SDK v2, which does not read bridgeage back). Only adopt the GET value when
	// the model value is unknown (unconfigured) so a Computed value stays known.
	if data.Bridgeage.IsUnknown() {
		if val, ok := getResponseData["bridgeage"]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				data.Bridgeage = types.Int64Value(intVal)
			} else {
				data.Bridgeage = types.Int64Null()
			}
		} else {
			data.Bridgeage = types.Int64Null()
		}
	}

	if val, ok := getResponseData["devicevlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Devicevlan = types.Int64Value(intVal)
		}
	} else if data.Devicevlan.IsUnknown() {
		data.Devicevlan = types.Int64Null()
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	} else if data.Ifnum.IsUnknown() {
		data.Ifnum = types.StringNull()
	}
	if val, ok := getResponseData["mac"]; ok && val != nil {
		data.Mac = types.StringValue(val.(string))
	} else if data.Mac.IsUnknown() {
		data.Mac = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else if data.Nodeid.IsUnknown() {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else if data.Vlan.IsUnknown() {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vni"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vni = types.Int64Value(intVal)
		}
	} else if data.Vni.IsUnknown() {
		data.Vni = types.Int64Null()
	}
	if val, ok := getResponseData["vtep"]; ok && val != nil {
		data.Vtep = types.StringValue(val.(string))
	} else if data.Vtep.IsUnknown() {
		data.Vtep = types.StringNull()
	}
	if val, ok := getResponseData["vxlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlan = types.Int64Value(intVal)
		}
	} else if data.Vxlan.IsUnknown() {
		data.Vxlan = types.Int64Null()
	}

	// Backward-compatible composite ID: "mac,vxlan,vtep" (matches SDK v2).
	data.Id = types.StringValue(fmt.Sprintf("%s,%d,%s", data.Mac.ValueString(), data.Vxlan.ValueInt64(), data.Vtep.ValueString()))

	return data
}

// bridgetableSetAttrFromGetForDatasource copies all attributes directly from the
// GET response (datasource reads have no prior state to preserve).
func bridgetableSetAttrFromGetForDatasource(ctx context.Context, data *BridgetableResourceModel, getResponseData map[string]interface{}) *BridgetableResourceModel {
	tflog.Debug(ctx, "In bridgetableSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["bridgeage"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bridgeage = types.Int64Value(intVal)
		} else {
			data.Bridgeage = types.Int64Null()
		}
	} else {
		data.Bridgeage = types.Int64Null()
	}
	if val, ok := getResponseData["devicevlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Devicevlan = types.Int64Value(intVal)
		} else {
			data.Devicevlan = types.Int64Null()
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
		} else {
			data.Nodeid = types.Int64Null()
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		} else {
			data.Vlan = types.Int64Null()
		}
	} else {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vni"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vni = types.Int64Value(intVal)
		} else {
			data.Vni = types.Int64Null()
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
		} else {
			data.Vxlan = types.Int64Null()
		}
	} else {
		data.Vxlan = types.Int64Null()
	}

	data.Id = types.StringValue(fmt.Sprintf("%s,%d,%s", data.Mac.ValueString(), data.Vxlan.ValueInt64(), data.Vtep.ValueString()))

	return data
}
