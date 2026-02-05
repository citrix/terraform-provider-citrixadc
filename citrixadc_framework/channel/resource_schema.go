package channel

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

// ChannelResourceModel describes the resource data model.
type ChannelResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Bandwidthhigh   types.Int64  `tfsdk:"bandwidthhigh"`
	Bandwidthnormal types.Int64  `tfsdk:"bandwidthnormal"`
	Conndistr       types.String `tfsdk:"conndistr"`
	Flowctl         types.String `tfsdk:"flowctl"`
	Haheartbeat     types.String `tfsdk:"haheartbeat"`
	Hamonitor       types.String `tfsdk:"hamonitor"`
	Channelid       types.String `tfsdk:"channelid"`
	Ifalias         types.String `tfsdk:"ifalias"`
	Ifnum           types.List   `tfsdk:"ifnum"`
	Lamac           types.String `tfsdk:"lamac"`
	Linkredundancy  types.String `tfsdk:"linkredundancy"`
	Lrminthroughput types.Int64  `tfsdk:"lrminthroughput"`
	Macdistr        types.String `tfsdk:"macdistr"`
	Mode            types.String `tfsdk:"mode"`
	Mtu             types.Int64  `tfsdk:"mtu"`
	Speed           types.String `tfsdk:"speed"`
	State           types.String `tfsdk:"state"`
	Tagall          types.String `tfsdk:"tagall"`
	Throughput      types.Int64  `tfsdk:"throughput"`
	Trunk           types.String `tfsdk:"trunk"`
}

func (r *ChannelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the channel resource.",
			},
			"bandwidthhigh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "High threshold value for the bandwidth usage of the LA channel, in Mbps. The Citrix ADC generates an SNMP trap message when the bandwidth usage of the LA channel is greater than or equal to the specified high threshold value.",
			},
			"bandwidthnormal": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Normal threshold value for the bandwidth usage of the LA channel, in Mbps. When the bandwidth usage of the LA channel returns to less than or equal to the specified normal threshold after exceeding the high threshold, the Citrix ADC generates an SNMP trap message to indicate that the bandwidth usage has returned to normal.",
			},
			"conndistr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The 'connection' distribution mode for the LA channel.",
			},
			"flowctl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the flow control type for this LA channel to manage the flow of frames. Flow control is a function as mentioned in clause 31 of the IEEE 802.3 standard. Flow control allows congested ports to pause traffic from the peer device. Flow control is achieved by sending PAUSE frames.",
			},
			"haheartbeat": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "In a High Availability (HA) configuration, configure the LA channel for sending heartbeats. LA channel that has HA Heartbeat disabled should not send the heartbeats.",
			},
			"hamonitor": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "In a High Availability (HA) configuration, monitor the LA channel for failure events. Failure of any LA channel that has HA MON enabled triggers HA failover.",
			},
			"channelid": schema.StringAttribute{
				Required:    true,
				Description: "ID for the LA channel or cluster LA channel or LR channel to be created. Specify an LA channel in LA/x notation, where x can range from 1 to 8 or cluster LA channel in CLA/x notation or Link redundant channel in LR/x notation, where x can range from 1 to 4. Cannot be changed after the LA channel is created.",
			},
			"ifalias": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString(" "),
				Description: "Alias name for the LA channel. Used only to enhance readability. To perform any operations, you have to specify the LA channel ID.",
			},
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Interfaces to be bound to the LA channel of a Citrix ADC or to the LA channel of a cluster configuration.\nFor an LA channel of a Citrix ADC, specify an interface in C/U notation (for example, 1/3).\nFor an LA channel of a cluster configuration, specify an interface in N/C/U notation (for example, 2/1/3).\nwhere C can take one of the following values:\n* 0 - Indicates a management interface.\n* 1 - Indicates a 1 Gbps port.\n* 10 - Indicates a 10 Gbps port.\nU is a unique integer for representing an interface in a particular port group.\nN is the ID of the node to which an interface belongs in a cluster configuration.\nUse spaces to separate multiple entries.",
			},
			"lamac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies a MAC address for the LA channels configured in Citrix ADC virtual appliances (VPX). This MAC address is persistent after each reboot.\nIf you don't specify this parameter, a MAC address is generated randomly for each LA channel. These MAC addresses change after each reboot.",
			},
			"linkredundancy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Link Redundancy for Cluster LAG.",
			},
			"lrminthroughput": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the minimum throughput threshold (in Mbps) to be met by the active subchannel. Setting this parameter automatically divides an LACP channel into logical subchannels, with one subchannel active and the others in standby mode.  When the maximum supported throughput of the active channel falls below the lrMinThroughput value, link failover occurs and a standby subchannel becomes active.",
			},
			"macdistr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The  'MAC' distribution mode for the LA channel.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The initital mode for the LA channel.",
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1500),
				Description: "The Maximum Transmission Unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes CRC, that can be transmitted and received by an interface. The default value of MTU is 1500 on all the interface of Citrix ADC, some Cloud Platforms will restrict Citrix ADC to use the lesser default value. Any MTU value more than 1500 is called Jumbo MTU and will make the interface as jumbo enabled. The Maximum Jumbo MTU in Citrix ADC is 9216, however, some Virtualized / Cloud Platforms will have lesser Maximum Jumbo MTU Value (9000). In the case of Cluster, the Backplane interface requires an MTU value of 78 bytes more than the Max MTU configured on any other Data-Plane Interface. When the Data plane interfaces are all at default 1500 MTU, Cluster Back Plane will be automatically set to 1578 (1500 + 78) MTU. If a Backplane interface is reset to Data Plane Interface, then the 1578 MTU will be automatically reset to the default MTU of 1500(or whatever lesser default value). If any data plane interface of a Cluster is configured with a Jumbo MTU ( > 1500), then all backplane interfaces require to be configured with a minimum MTU of 'Highest Data Plane MTU in the Cluster + 78'. That makes the maximum Jumbo MTU for any Data-Plane Interface in a Cluster System to be '9138 (9216 - 78)., where 9216 is the maximum Jumbo MTU. On certain Virtualized / Cloud Platforms, the maximum  possible MTU is restricted to a lesser value, Similar calculation can be applied, Maximum Data Plane MTU in Cluster = (Maximum possible MTU - 78).",
			},
			"speed": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("AUTO"),
				Description: "Ethernet speed of the channel, in Mbps. If the speed of any bound interface is greater than or equal to the value set for this parameter, the state of the interface is UP. Otherwise, the state is INACTIVE. Bound Interfaces whose state is INACTIVE do not process any traffic.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable the LA channel.",
			},
			"tagall": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Adds a four-byte 802.1q tag to every packet sent on this channel.  The ON setting applies tags for all VLANs that are bound to this channel. OFF applies the tag for all VLANs other than the native VLAN.",
			},
			"throughput": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Low threshold value for the throughput of the LA channel, in Mbps. In an high availability (HA) configuration, failover is triggered when the LA channel has HA MON enabled and the throughput is below the specified threshold.",
			},
			"trunk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is deprecated by tagall",
			},
		},
	}
}

func channelGetThePayloadFromtheConfig(ctx context.Context, data *ChannelResourceModel) network.Channel {
	tflog.Debug(ctx, "In channelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	channel := network.Channel{}
	if !data.Bandwidthhigh.IsNull() {
		channel.Bandwidthhigh = utils.IntPtr(int(data.Bandwidthhigh.ValueInt64()))
	}
	if !data.Bandwidthnormal.IsNull() {
		channel.Bandwidthnormal = utils.IntPtr(int(data.Bandwidthnormal.ValueInt64()))
	}
	if !data.Conndistr.IsNull() {
		channel.Conndistr = data.Conndistr.ValueString()
	}
	if !data.Flowctl.IsNull() {
		channel.Flowctl = data.Flowctl.ValueString()
	}
	if !data.Haheartbeat.IsNull() {
		channel.Haheartbeat = data.Haheartbeat.ValueString()
	}
	if !data.Hamonitor.IsNull() {
		channel.Hamonitor = data.Hamonitor.ValueString()
	}
	if !data.Id.IsNull() {
		channel.Id = data.Id.ValueString()
	}
	if !data.Ifalias.IsNull() {
		channel.Ifalias = data.Ifalias.ValueString()
	}
	if !data.Lamac.IsNull() {
		channel.Lamac = data.Lamac.ValueString()
	}
	if !data.Linkredundancy.IsNull() {
		channel.Linkredundancy = data.Linkredundancy.ValueString()
	}
	if !data.Lrminthroughput.IsNull() {
		channel.Lrminthroughput = utils.IntPtr(int(data.Lrminthroughput.ValueInt64()))
	}
	if !data.Macdistr.IsNull() {
		channel.Macdistr = data.Macdistr.ValueString()
	}
	if !data.Mode.IsNull() {
		channel.Mode = data.Mode.ValueString()
	}
	if !data.Mtu.IsNull() {
		channel.Mtu = utils.IntPtr(int(data.Mtu.ValueInt64()))
	}
	if !data.Speed.IsNull() {
		channel.Speed = data.Speed.ValueString()
	}
	if !data.State.IsNull() {
		channel.State = data.State.ValueString()
	}
	if !data.Tagall.IsNull() {
		channel.Tagall = data.Tagall.ValueString()
	}
	if !data.Throughput.IsNull() {
		channel.Throughput = utils.IntPtr(int(data.Throughput.ValueInt64()))
	}
	if !data.Trunk.IsNull() {
		channel.Trunk = data.Trunk.ValueString()
	}

	return channel
}

func channelSetAttrFromGet(ctx context.Context, data *ChannelResourceModel, getResponseData map[string]interface{}) *ChannelResourceModel {
	tflog.Debug(ctx, "In channelSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["conndistr"]; ok && val != nil {
		data.Conndistr = types.StringValue(val.(string))
	} else {
		data.Conndistr = types.StringNull()
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
		data.Id = types.StringValue(val.(string))
	} else {
		data.Id = types.StringNull()
	}
	if val, ok := getResponseData["ifalias"]; ok && val != nil {
		data.Ifalias = types.StringValue(val.(string))
	} else {
		data.Ifalias = types.StringNull()
	}
	if val, ok := getResponseData["lamac"]; ok && val != nil {
		data.Lamac = types.StringValue(val.(string))
	} else {
		data.Lamac = types.StringNull()
	}
	if val, ok := getResponseData["linkredundancy"]; ok && val != nil {
		data.Linkredundancy = types.StringValue(val.(string))
	} else {
		data.Linkredundancy = types.StringNull()
	}
	if val, ok := getResponseData["lrminthroughput"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Lrminthroughput = types.Int64Value(intVal)
		}
	} else {
		data.Lrminthroughput = types.Int64Null()
	}
	if val, ok := getResponseData["macdistr"]; ok && val != nil {
		data.Macdistr = types.StringValue(val.(string))
	} else {
		data.Macdistr = types.StringNull()
	}
	if val, ok := getResponseData["mode"]; ok && val != nil {
		data.Mode = types.StringValue(val.(string))
	} else {
		data.Mode = types.StringNull()
	}
	if val, ok := getResponseData["mtu"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mtu = types.Int64Value(intVal)
		}
	} else {
		data.Mtu = types.Int64Null()
	}
	if val, ok := getResponseData["speed"]; ok && val != nil {
		data.Speed = types.StringValue(val.(string))
	} else {
		data.Speed = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Channelid.ValueString())

	return data
}
