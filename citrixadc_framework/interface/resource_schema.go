package Interface

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// InterfaceResourceModel describes the resource data model.
type InterfaceResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Autoneg          types.String `tfsdk:"autoneg"`
	Bandwidthhigh    types.Int64  `tfsdk:"bandwidthhigh"`
	Bandwidthnormal  types.Int64  `tfsdk:"bandwidthnormal"`
	Duplex           types.String `tfsdk:"duplex"`
	Flowctl          types.String `tfsdk:"flowctl"`
	Haheartbeat      types.String `tfsdk:"haheartbeat"`
	Hamonitor        types.String `tfsdk:"hamonitor"`
	Interfaceid      types.String `tfsdk:"interface_id"`
	Ifalias          types.String `tfsdk:"ifalias"`
	Lacpkey          types.Int64  `tfsdk:"lacpkey"`
	Lacpmode         types.String `tfsdk:"lacpmode"`
	Lacppriority     types.Int64  `tfsdk:"lacppriority"`
	Lacptimeout      types.String `tfsdk:"lacptimeout"`
	Lagtype          types.String `tfsdk:"lagtype"`
	Linkredundancy   types.String `tfsdk:"linkredundancy"`
	Lldpmode         types.String `tfsdk:"lldpmode"`
	Lrsetpriority    types.Int64  `tfsdk:"lrsetpriority"`
	Mtu              types.Int64  `tfsdk:"mtu"`
	Ringsize         types.Int64  `tfsdk:"ringsize"`
	Ringtype         types.String `tfsdk:"ringtype"`
	Speed            types.String `tfsdk:"speed"`
	Tagall           types.String `tfsdk:"tagall"`
	Throughput       types.Int64  `tfsdk:"throughput"`
	Trunk            types.String `tfsdk:"trunk"`
	Trunkallowedvlan types.List   `tfsdk:"trunkallowedvlan"`
	Trunkmode        types.String `tfsdk:"trunkmode"`
}

func (r *InterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the interface resource.",
			},
			"autoneg": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NSA_DVC_AUTONEG_ON"),
				Description: "Auto-negotiation state of the interface. With the ENABLED setting, the Citrix ADC auto-negotiates the speed and duplex settings with the peer network device on the link. The Citrix ADC appliance auto-negotiates the settings of only those parameters (speed or duplex mode) for which the value is set as AUTO.",
			},
			"bandwidthhigh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "High threshold value for the bandwidth usage of the interface, in Mbps. The Citrix ADC generates an SNMP trap message when the bandwidth usage of the interface is greater than or equal to the specified high threshold value.",
			},
			"bandwidthnormal": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Normal threshold value for the bandwidth usage of the interface, in Mbps. When the bandwidth usage of the interface becomes less than or equal to the specified normal threshold after exceeding the high threshold, the Citrix ADC generates an SNMP trap message to indicate that the bandwidth usage has returned to normal.",
			},
			"duplex": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("AUTO"),
				Description: "The duplex mode for the interface. Notes:* If you set the duplex mode to AUTO, the Citrix ADC attempts to auto-negotiate the duplex mode of the interface when it is UP. You must enable auto negotiation on the interface. If you set a duplex mode other than AUTO, you must specify the same duplex mode for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors.",
			},
			"flowctl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "802.3x flow control setting for the interface.  The 802.3x specification does not define flow control for 10 Mbps and 100 Mbps speeds, but if a Gigabit Ethernet interface operates at those speeds, the flow control settings can be applied. The flow control setting that is finally applied to an interface depends on auto-negotiation. With the ON option, the peer negotiates the flow control, but the appliance then forces two-way flow control for the interface.",
			},
			"haheartbeat": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "In a High Availability (HA) or Cluster configuration, configure the interface for sending heartbeats. In an HA or Cluster configuration, an interface that has HA Heartbeat disabled should not send the heartbeats.",
			},
			"hamonitor": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "In a High Availability (HA) configuration, monitor the interface for failure events. In an HA configuration, an interface that has HA MON enabled and is not bound to any Failover Interface Set (FIS), is a critical interface. Failure or disabling of any critical interface triggers HA failover.",
			},
			"interface_id": schema.StringAttribute{
				Required:    true,
				Description: "Interface number, in C/U format, where C can take one of the following values:\n* 0 - Indicates a management interface.\n* 1 - Indicates a 1 Gbps port.\n* 10 - Indicates a 10 Gbps port.\n* LA - Indicates a link aggregation port.\n* LO - Indicates a loop back port.\nU is a unique integer for representing an interface in a particular port group.",
			},
			"ifalias": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString(" "),
				Description: "Alias name for the interface. Used only to enhance readability. To perform any operations, you have to specify the interface ID.",
			},
			"lacpkey": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer identifying the LACP LA channel to which the interface is to be bound.\nFor an LA channel of the Citrix ADC, this digit specifies the variable x of an LA channel in LA/x notation, where x can range from 1 to 8. For example, if you specify 3 as the LACP key for an LA channel, the interface is bound to the LA channel LA/3.\nFor an LA channel of a cluster configuration, this digit specifies the variable y of a cluster LA channel in CLA/(y-4) notation, where y can range from 5 to 8. For example, if you specify 6 as the LACP key for a cluster LA channel, the interface is bound to the cluster LA channel CLA/2.",
			},
			"lacpmode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Bind the interface to a LA channel created by the Link Aggregation control protocol (LACP).\nAvailable settings function as follows:\n* Active - The LA channel port of the Citrix ADC generates LACPDU messages on a regular basis, regardless of any need expressed by its peer device to receive them.\n* Passive - The LA channel port of the Citrix ADC does not transmit LACPDU messages unless the peer device port is in the active mode. That is, the port does not speak unless spoken to.\n* Disabled - Unbinds the interface from the LA channel. If this is the only interface in the LA channel, the LA channel is removed.",
			},
			"lacppriority": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(32768),
				Description: "LACP port priority, expressed as an integer. The lower the number, the higher the priority. The Citrix ADC limits the number of interfaces in an LA channel to sixteen.",
			},
			"lacptimeout": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("LONG"),
				Description: "Interval at which the Citrix ADC sends LACPDU messages to the peer device on the LA channel.\nAvailable settings function as follows:\nLONG - 30 seconds.\nSHORT - 1 second.",
			},
			"lagtype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NODE"),
				Description: "Type of entity (Citrix ADC or cluster configuration) for which to create the channel.",
			},
			"linkredundancy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Link Redundancy for Cluster LAG.",
			},
			"lldpmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Link Layer Discovery Protocol (LLDP) mode for an interface. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels.",
			},
			"lrsetpriority": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1024),
				Description: "LRSET port priority, expressed as an integer ranging from 1 to 1024. The highest priority is 1. The Citrix ADC limits the number of interfaces in an LRSET to 8. Within a LRSET the highest LR Priority Interface is considered as the first candidate for the Active interface, if the interface is UP.",
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1500),
				Description: "The Maximum Transmission Unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes CRC, that can be transmitted and received by an interface. The default value of MTU is 1500 on all the interface of Citrix ADC, some Cloud Platforms will restrict Citrix ADC to use the lesser default value. Any MTU value more than 1500 is called Jumbo MTU and will make the interface as jumbo enabled. The Maximum Jumbo MTU in Citrix ADC is 9216, however, some Virtualized / Cloud Platforms will have lesser Maximum Jumbo MTU Value (9000). In the case of Cluster, the Backplane interface requires an MTU value of 78 bytes more than the Max MTU configured on any other Data-Plane Interface. When the Data plane interfaces are all at default 1500 MTU, Cluster Back Plane will be automatically set to 1578 (1500 + 78) MTU. If a Backplane interface is reset to Data Plane Interface, then the 1578 MTU will be automatically reset to the default MTU of 1500(or whatever lesser default value). If any data plane interface of a Cluster is configured with a Jumbo MTU ( > 1500), then all backplane interfaces require to be configured with a minimum MTU of 'Highest Data Plane MTU in the Cluster + 78'. That makes the maximum Jumbo MTU for any Data-Plane Interface in a Cluster System to be '9138 (9216 - 78)., where 9216 is the maximum Jumbo MTU. On certain Virtualized / Cloud Platforms, the maximum  possible MTU is restricted to a lesser value, Similar calculation can be applied, Maximum Data Plane MTU in Cluster = (Maximum possible MTU - 78).",
			},
			"ringsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2048),
				Description: "The receive ringsize of the interface. A higher number provides more number of buffers in handling incoming traffic.",
			},
			"ringtype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Elastic"),
				Description: "The receive ringtype of the interface (Fixed or Elastic). A fixed ring type pre-allocates configured number of buffers irrespective of traffic rate. In contrast, an elastic ring, expands and shrinks based on incoming traffic rate.",
			},
			"speed": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("AUTO"),
				Description: "Ethernet speed of the interface, in Mbps.\nNotes:\n* If you set the speed as AUTO, the Citrix ADC attempts to auto-negotiate or auto-sense the link speed of the interface when it is UP. You must enable auto negotiation on the interface.\n* If you set a speed other than AUTO, you must specify the same speed for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors.\nSome interfaces do not support certain speeds. If you specify an unsupported speed, an error message appears.",
			},
			"tagall": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add a four-byte 802.1q tag to every packet sent on this interface.  The ON setting applies the tag for this interface's native VLAN. OFF applies the tag for all VLANs other than the native VLAN.",
			},
			"throughput": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Low threshold value for the throughput of the interface, in Mbps. In an HA configuration, failover is triggered if the interface has HA MON enabled and the throughput is below the specified the threshold.",
			},
			"trunk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is deprecated by tagall.",
			},
			"trunkallowedvlan": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "VLAN ID or range of VLAN IDs will be allowed on this trunk interface. In the command line interface, separate the range with a hyphen. For example: 40-90.",
			},
			"trunkmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Accept and send 802.1q VLAN tagged packets, based on Allowed Vlan List of this interface.",
			},
		},
	}
}

func interfaceGetThePayloadFromtheConfig(ctx context.Context, data *InterfaceResourceModel) network.Interface {
	tflog.Debug(ctx, "In interfaceGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	Interface := network.Interface{}
	if !data.Autoneg.IsNull() {
		Interface.Autoneg = data.Autoneg.ValueString()
	}
	if !data.Bandwidthhigh.IsNull() {
		Interface.Bandwidthhigh = utils.IntPtr(int(data.Bandwidthhigh.ValueInt64()))
	}
	if !data.Bandwidthnormal.IsNull() {
		Interface.Bandwidthnormal = utils.IntPtr(int(data.Bandwidthnormal.ValueInt64()))
	}
	if !data.Duplex.IsNull() {
		Interface.Duplex = data.Duplex.ValueString()
	}
	if !data.Flowctl.IsNull() {
		Interface.Flowctl = data.Flowctl.ValueString()
	}
	if !data.Haheartbeat.IsNull() {
		Interface.Haheartbeat = data.Haheartbeat.ValueString()
	}
	if !data.Hamonitor.IsNull() {
		Interface.Hamonitor = data.Hamonitor.ValueString()
	}
	if !data.Interfaceid.IsNull() {
		Interface.Id = data.Interfaceid.ValueString()
	}
	if !data.Ifalias.IsNull() {
		Interface.Ifalias = data.Ifalias.ValueString()
	}
	if !data.Lacpkey.IsNull() {
		Interface.Lacpkey = utils.IntPtr(int(data.Lacpkey.ValueInt64()))
	}
	if !data.Lacpmode.IsNull() {
		Interface.Lacpmode = data.Lacpmode.ValueString()
	}
	if !data.Lacppriority.IsNull() {
		Interface.Lacppriority = utils.IntPtr(int(data.Lacppriority.ValueInt64()))
	}
	if !data.Lacptimeout.IsNull() {
		Interface.Lacptimeout = data.Lacptimeout.ValueString()
	}
	if !data.Lagtype.IsNull() {
		Interface.Lagtype = data.Lagtype.ValueString()
	}
	if !data.Linkredundancy.IsNull() {
		Interface.Linkredundancy = data.Linkredundancy.ValueString()
	}
	if !data.Lldpmode.IsNull() {
		Interface.Lldpmode = data.Lldpmode.ValueString()
	}
	if !data.Lrsetpriority.IsNull() {
		Interface.Lrsetpriority = utils.IntPtr(int(data.Lrsetpriority.ValueInt64()))
	}
	if !data.Mtu.IsNull() {
		Interface.Mtu = utils.IntPtr(int(data.Mtu.ValueInt64()))
	}
	if !data.Ringsize.IsNull() {
		Interface.Ringsize = utils.IntPtr(int(data.Ringsize.ValueInt64()))
	}
	if !data.Ringtype.IsNull() {
		Interface.Ringtype = data.Ringtype.ValueString()
	}
	if !data.Speed.IsNull() {
		Interface.Speed = data.Speed.ValueString()
	}
	if !data.Tagall.IsNull() {
		Interface.Tagall = data.Tagall.ValueString()
	}
	if !data.Throughput.IsNull() {
		Interface.Throughput = utils.IntPtr(int(data.Throughput.ValueInt64()))
	}
	if !data.Trunk.IsNull() {
		Interface.Trunk = data.Trunk.ValueString()
	}
	if !data.Trunkmode.IsNull() {
		Interface.Trunkmode = data.Trunkmode.ValueString()
	}

	return Interface
}

func interfaceSetAttrFromGet(ctx context.Context, data *InterfaceResourceModel, getResponseData map[string]interface{}) *InterfaceResourceModel {
	tflog.Debug(ctx, "In interfaceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["autoneg"]; ok && val != nil {
		data.Autoneg = types.StringValue(val.(string))
	} else {
		data.Autoneg = types.StringNull()
	}
	if val, ok := getResponseData["bandwidthhigh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bandwidthhigh = types.Int64Value(intVal)
		}
	} else {
		data.Bandwidthhigh = types.Int64Null()
	}
	if val, ok := getResponseData["bandwidthnormal"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bandwidthnormal = types.Int64Value(intVal)
		}
	} else {
		data.Bandwidthnormal = types.Int64Null()
	}
	if val, ok := getResponseData["duplex"]; ok && val != nil {
		data.Duplex = types.StringValue(val.(string))
	} else {
		data.Duplex = types.StringNull()
	}
	if val, ok := getResponseData["flowctl"]; ok && val != nil {
		data.Flowctl = types.StringValue(val.(string))
	} else {
		data.Flowctl = types.StringNull()
	}
	if val, ok := getResponseData["haheartbeat"]; ok && val != nil {
		data.Haheartbeat = types.StringValue(val.(string))
	} else {
		data.Haheartbeat = types.StringNull()
	}
	if val, ok := getResponseData["hamonitor"]; ok && val != nil {
		data.Hamonitor = types.StringValue(val.(string))
	} else {
		data.Hamonitor = types.StringNull()
	}
	if val, ok := getResponseData["id"]; ok && val != nil {
		data.Interfaceid = types.StringValue(val.(string))
	} else {
		data.Interfaceid = types.StringNull()
	}
	if val, ok := getResponseData["ifalias"]; ok && val != nil {
		data.Ifalias = types.StringValue(val.(string))
	} else {
		data.Ifalias = types.StringNull()
	}
	if val, ok := getResponseData["lacpkey"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Lacpkey = types.Int64Value(intVal)
		}
	} else {
		data.Lacpkey = types.Int64Null()
	}
	if val, ok := getResponseData["lacpmode"]; ok && val != nil {
		data.Lacpmode = types.StringValue(val.(string))
	} else {
		data.Lacpmode = types.StringNull()
	}
	if val, ok := getResponseData["lacppriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Lacppriority = types.Int64Value(intVal)
		}
	} else {
		data.Lacppriority = types.Int64Null()
	}
	if val, ok := getResponseData["lacptimeout"]; ok && val != nil {
		data.Lacptimeout = types.StringValue(val.(string))
	} else {
		data.Lacptimeout = types.StringNull()
	}
	if val, ok := getResponseData["lagtype"]; ok && val != nil {
		data.Lagtype = types.StringValue(val.(string))
	} else {
		data.Lagtype = types.StringNull()
	}
	if val, ok := getResponseData["linkredundancy"]; ok && val != nil {
		data.Linkredundancy = types.StringValue(val.(string))
	} else {
		data.Linkredundancy = types.StringNull()
	}
	if val, ok := getResponseData["lldpmode"]; ok && val != nil {
		data.Lldpmode = types.StringValue(val.(string))
	} else {
		data.Lldpmode = types.StringNull()
	}
	if val, ok := getResponseData["lrsetpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Lrsetpriority = types.Int64Value(intVal)
		}
	} else {
		data.Lrsetpriority = types.Int64Null()
	}
	if val, ok := getResponseData["mtu"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mtu = types.Int64Value(intVal)
		}
	} else {
		data.Mtu = types.Int64Null()
	}
	if val, ok := getResponseData["ringsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ringsize = types.Int64Value(intVal)
		}
	} else {
		data.Ringsize = types.Int64Null()
	}
	if val, ok := getResponseData["ringtype"]; ok && val != nil {
		data.Ringtype = types.StringValue(val.(string))
	} else {
		data.Ringtype = types.StringNull()
	}
	if val, ok := getResponseData["speed"]; ok && val != nil {
		data.Speed = types.StringValue(val.(string))
	} else {
		data.Speed = types.StringNull()
	}
	if val, ok := getResponseData["tagall"]; ok && val != nil {
		data.Tagall = types.StringValue(val.(string))
	} else {
		data.Tagall = types.StringNull()
	}
	if val, ok := getResponseData["throughput"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Throughput = types.Int64Value(intVal)
		}
	} else {
		data.Throughput = types.Int64Null()
	}
	if val, ok := getResponseData["trunk"]; ok && val != nil {
		data.Trunk = types.StringValue(val.(string))
	} else {
		data.Trunk = types.StringNull()
	}
	if val, ok := getResponseData["trunkmode"]; ok && val != nil {
		data.Trunkmode = types.StringValue(val.(string))
	} else {
		data.Trunkmode = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("interface-config")

	return data
}
