package nspbr

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NspbrResourceModel describes the resource data model.
type NspbrResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Interface      types.String `tfsdk:"interface"`
	Action         types.String `tfsdk:"action"`
	Destip         types.Bool   `tfsdk:"destip"`
	Destipop       types.String `tfsdk:"destipop"`
	Destipval      types.String `tfsdk:"destipval"`
	Destport       types.Bool   `tfsdk:"destport"`
	Destportop     types.String `tfsdk:"destportop"`
	Destportval    types.String `tfsdk:"destportval"`
	Detail         types.Bool   `tfsdk:"detail"`
	Iptunnel       types.Bool   `tfsdk:"iptunnel"`
	Iptunnelname   types.String `tfsdk:"iptunnelname"`
	Monitor        types.String `tfsdk:"monitor"`
	Msr            types.String `tfsdk:"msr"`
	Name           types.String `tfsdk:"name"`
	Nexthop        types.Bool   `tfsdk:"nexthop"`
	Nexthopval     types.String `tfsdk:"nexthopval"`
	Ownergroup     types.String `tfsdk:"ownergroup"`
	Priority       types.Int64  `tfsdk:"priority"`
	Protocol       types.String `tfsdk:"protocol"`
	Protocolnumber types.Int64  `tfsdk:"protocolnumber"`
	Srcip          types.Bool   `tfsdk:"srcip"`
	Srcipop        types.String `tfsdk:"srcipop"`
	Srcipval       types.String `tfsdk:"srcipval"`
	Srcmac         types.String `tfsdk:"srcmac"`
	Srcmacmask     types.String `tfsdk:"srcmacmask"`
	Srcport        types.Bool   `tfsdk:"srcport"`
	Srcportop      types.String `tfsdk:"srcportop"`
	Srcportval     types.String `tfsdk:"srcportval"`
	State          types.String `tfsdk:"state"`
	Targettd       types.Int64  `tfsdk:"targettd"`
	Td             types.Int64  `tfsdk:"td"`
	Vlan           types.Int64  `tfsdk:"vlan"`
	Vxlan          types.Int64  `tfsdk:"vxlan"`
	Vxlanvlanmap   types.String `tfsdk:"vxlanvlanmap"`
}

func (r *NspbrResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nspbr resource.",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of an interface. The Citrix ADC compares the PBR only to the outgoing packets on the specified interface. If you do not specify any value, the appliance compares the PBR to the outgoing packets on all interfaces.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action to perform on the outgoing IPv4 packets that match the PBR.\n\nAvailable settings function as follows:\n* ALLOW - The Citrix ADC sends the packet to the designated next-hop router.\n* DENY - The Citrix ADC applies the routing table for normal destination-based routing.",
			},
			"destip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an outgoing IPv4 packet.  In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.",
			},
			"destipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"destipval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.",
			},
			"destport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the destination port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
			},
			"destportop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"destportval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the destination port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
			},
			"detail": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "To get a detailed view.",
			},
			"iptunnel": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Tunnel name.",
			},
			"iptunnelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The iptunnel name where packets need to be forwarded upon.",
			},
			"monitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the monitor.(Can be only of type ping or ARP )",
			},
			"msr": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Monitor the route specified byte Next Hop parameter. This parameter is not applicable if you specify a link load balancing (LLB) virtual server name with the Next Hop parameter.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the PBR. Must begin with an ASCII alphabetic or underscore \\(_\\) character, and must contain only ASCII alphanumeric, underscore, hash \\(\\#\\), period \\(.\\), space, colon \\(:\\), at \\(@\\), equals \\(=\\), and hyphen \\(-\\) characters. Cannot be changed after the PBR is created.",
			},
			"nexthop": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the next hop router or the name of the link load balancing virtual server to which to send matching packets if action is set to ALLOW.\nIf you specify a link load balancing (LLB) virtual server, which can provide a backup if a next hop link fails, first make sure that the next hops bound to the LLB virtual server are actually next hops that are directly connected to the Citrix ADC. Otherwise, the Citrix ADC throws an error when you attempt to create the PBR. The next hop can be null to represent null routes",
			},
			"nexthopval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Next Hop IP address or gateway name.",
			},
			"ownergroup": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("DEFAULT_NG"),
				Description: "The owner node group in a Cluster for this pbr rule. If ownernode is not specified then the pbr rule is treated as Striped pbr rule.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the PBR, which determines the order in which it is evaluated relative to the other PBRs. If you do not specify priorities while creating PBRs, the PBRs are evaluated in the order in which they are created.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, identified by protocol name, to match against the protocol of an outgoing IPv4 packet.",
			},
			"protocolnumber": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, identified by protocol number, to match against the protocol of an outgoing IPv4 packet.",
			},
			"srcip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the source IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.",
			},
			"srcipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"srcipval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the source IP address of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.",
			},
			"srcmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address to match against the source MAC address of an outgoing IPv4 packet.",
			},
			"srcmacmask": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("000000000000"),
				Description: "Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value \"000000111111\".",
			},
			"srcport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the source port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
			},
			"srcportop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"srcportval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the source port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable the PBR. After you apply the PBRs, the Citrix ADC compares outgoing packets to the enabled PBRs.",
			},
			"targettd": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain to which you want to send packet to.",
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
				Optional:    true,
				Computed:    true,
				Description: "ID of the VLAN. The Citrix ADC compares the PBR only to the outgoing packets on the specified VLAN. If you do not specify any interface ID, the appliance compares the PBR to the outgoing packets on all VLANs.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN. The Citrix ADC compares the PBR only to the outgoing packets on the specified VXLAN. If you do not specify any interface ID, the appliance compares the PBR to the outgoing packets on all VXLANs.",
			},
			"vxlanvlanmap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The vlan to vxlan mapping to be applied for incoming packets over this pbr tunnel",
			},
		},
	}
}

func nspbrGetThePayloadFromtheConfig(ctx context.Context, data *NspbrResourceModel) ns.Nspbr {
	tflog.Debug(ctx, "In nspbrGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nspbr := ns.Nspbr{}
	if !data.Interface.IsNull() {
		nspbr.Interface = data.Interface.ValueString()
	}
	if !data.Action.IsNull() {
		nspbr.Action = data.Action.ValueString()
	}
	if !data.Destip.IsNull() {
		nspbr.Destip = data.Destip.ValueBool()
	}
	if !data.Destipop.IsNull() {
		nspbr.Destipop = data.Destipop.ValueString()
	}
	if !data.Destipval.IsNull() {
		nspbr.Destipval = data.Destipval.ValueString()
	}
	if !data.Destport.IsNull() {
		nspbr.Destport = data.Destport.ValueBool()
	}
	if !data.Destportop.IsNull() {
		nspbr.Destportop = data.Destportop.ValueString()
	}
	if !data.Destportval.IsNull() {
		nspbr.Destportval = data.Destportval.ValueString()
	}
	if !data.Detail.IsNull() {
		nspbr.Detail = data.Detail.ValueBool()
	}
	if !data.Iptunnel.IsNull() {
		nspbr.Iptunnel = data.Iptunnel.ValueBool()
	}
	if !data.Iptunnelname.IsNull() {
		nspbr.Iptunnelname = data.Iptunnelname.ValueString()
	}
	if !data.Monitor.IsNull() {
		nspbr.Monitor = data.Monitor.ValueString()
	}
	if !data.Msr.IsNull() {
		nspbr.Msr = data.Msr.ValueString()
	}
	if !data.Name.IsNull() {
		nspbr.Name = data.Name.ValueString()
	}
	if !data.Nexthop.IsNull() {
		nspbr.Nexthop = data.Nexthop.ValueBool()
	}
	if !data.Nexthopval.IsNull() {
		nspbr.Nexthopval = data.Nexthopval.ValueString()
	}
	if !data.Ownergroup.IsNull() {
		nspbr.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Priority.IsNull() {
		nspbr.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Protocol.IsNull() {
		nspbr.Protocol = data.Protocol.ValueString()
	}
	if !data.Protocolnumber.IsNull() {
		nspbr.Protocolnumber = utils.IntPtr(int(data.Protocolnumber.ValueInt64()))
	}
	if !data.Srcip.IsNull() {
		nspbr.Srcip = data.Srcip.ValueBool()
	}
	if !data.Srcipop.IsNull() {
		nspbr.Srcipop = data.Srcipop.ValueString()
	}
	if !data.Srcipval.IsNull() {
		nspbr.Srcipval = data.Srcipval.ValueString()
	}
	if !data.Srcmac.IsNull() {
		nspbr.Srcmac = data.Srcmac.ValueString()
	}
	if !data.Srcmacmask.IsNull() {
		nspbr.Srcmacmask = data.Srcmacmask.ValueString()
	}
	if !data.Srcport.IsNull() {
		nspbr.Srcport = data.Srcport.ValueBool()
	}
	if !data.Srcportop.IsNull() {
		nspbr.Srcportop = data.Srcportop.ValueString()
	}
	if !data.Srcportval.IsNull() {
		nspbr.Srcportval = data.Srcportval.ValueString()
	}
	if !data.State.IsNull() {
		nspbr.State = data.State.ValueString()
	}
	if !data.Targettd.IsNull() {
		nspbr.Targettd = utils.IntPtr(int(data.Targettd.ValueInt64()))
	}
	if !data.Td.IsNull() {
		nspbr.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vlan.IsNull() {
		nspbr.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vxlan.IsNull() {
		nspbr.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}
	if !data.Vxlanvlanmap.IsNull() {
		nspbr.Vxlanvlanmap = data.Vxlanvlanmap.ValueString()
	}

	return nspbr
}

func nspbrSetAttrFromGet(ctx context.Context, data *NspbrResourceModel, getResponseData map[string]interface{}) *NspbrResourceModel {
	tflog.Debug(ctx, "In nspbrSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Interface"]; ok && val != nil {
		data.Interface = types.StringValue(val.(string))
	} else {
		data.Interface = types.StringNull()
	}
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["destip"]; ok && val != nil {
		data.Destip = types.BoolValue(val.(bool))
	} else {
		data.Destip = types.BoolNull()
	}
	if val, ok := getResponseData["destipop"]; ok && val != nil {
		data.Destipop = types.StringValue(val.(string))
	} else {
		data.Destipop = types.StringNull()
	}
	if val, ok := getResponseData["destipval"]; ok && val != nil {
		data.Destipval = types.StringValue(val.(string))
	} else {
		data.Destipval = types.StringNull()
	}
	if val, ok := getResponseData["destport"]; ok && val != nil {
		data.Destport = types.BoolValue(val.(bool))
	} else {
		data.Destport = types.BoolNull()
	}
	if val, ok := getResponseData["destportop"]; ok && val != nil {
		data.Destportop = types.StringValue(val.(string))
	} else {
		data.Destportop = types.StringNull()
	}
	if val, ok := getResponseData["destportval"]; ok && val != nil {
		data.Destportval = types.StringValue(val.(string))
	} else {
		data.Destportval = types.StringNull()
	}
	if val, ok := getResponseData["detail"]; ok && val != nil {
		data.Detail = types.BoolValue(val.(bool))
	} else {
		data.Detail = types.BoolNull()
	}
	if val, ok := getResponseData["iptunnel"]; ok && val != nil {
		data.Iptunnel = types.BoolValue(val.(bool))
	} else {
		data.Iptunnel = types.BoolNull()
	}
	if val, ok := getResponseData["iptunnelname"]; ok && val != nil {
		data.Iptunnelname = types.StringValue(val.(string))
	} else {
		data.Iptunnelname = types.StringNull()
	}
	if val, ok := getResponseData["monitor"]; ok && val != nil {
		data.Monitor = types.StringValue(val.(string))
	} else {
		data.Monitor = types.StringNull()
	}
	if val, ok := getResponseData["msr"]; ok && val != nil {
		data.Msr = types.StringValue(val.(string))
	} else {
		data.Msr = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["nexthop"]; ok && val != nil {
		data.Nexthop = types.BoolValue(val.(bool))
	} else {
		data.Nexthop = types.BoolNull()
	}
	if val, ok := getResponseData["nexthopval"]; ok && val != nil {
		data.Nexthopval = types.StringValue(val.(string))
	} else {
		data.Nexthopval = types.StringNull()
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		data.Protocol = types.StringValue(val.(string))
	} else {
		data.Protocol = types.StringNull()
	}
	if val, ok := getResponseData["protocolnumber"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Protocolnumber = types.Int64Value(intVal)
		}
	} else {
		data.Protocolnumber = types.Int64Null()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.BoolValue(val.(bool))
	} else {
		data.Srcip = types.BoolNull()
	}
	if val, ok := getResponseData["srcipop"]; ok && val != nil {
		data.Srcipop = types.StringValue(val.(string))
	} else {
		data.Srcipop = types.StringNull()
	}
	if val, ok := getResponseData["srcipval"]; ok && val != nil {
		data.Srcipval = types.StringValue(val.(string))
	} else {
		data.Srcipval = types.StringNull()
	}
	if val, ok := getResponseData["srcmac"]; ok && val != nil {
		data.Srcmac = types.StringValue(val.(string))
	} else {
		data.Srcmac = types.StringNull()
	}
	if val, ok := getResponseData["srcmacmask"]; ok && val != nil {
		data.Srcmacmask = types.StringValue(val.(string))
	} else {
		data.Srcmacmask = types.StringNull()
	}
	if val, ok := getResponseData["srcport"]; ok && val != nil {
		data.Srcport = types.BoolValue(val.(bool))
	} else {
		data.Srcport = types.BoolNull()
	}
	if val, ok := getResponseData["srcportop"]; ok && val != nil {
		data.Srcportop = types.StringValue(val.(string))
	} else {
		data.Srcportop = types.StringNull()
	}
	if val, ok := getResponseData["srcportval"]; ok && val != nil {
		data.Srcportval = types.StringValue(val.(string))
	} else {
		data.Srcportval = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["targettd"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Targettd = types.Int64Value(intVal)
		}
	} else {
		data.Targettd = types.Int64Null()
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
	if val, ok := getResponseData["vxlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlan = types.Int64Value(intVal)
		}
	} else {
		data.Vxlan = types.Int64Null()
	}
	if val, ok := getResponseData["vxlanvlanmap"]; ok && val != nil {
		data.Vxlanvlanmap = types.StringValue(val.(string))
	} else {
		data.Vxlanvlanmap = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Name.ValueString()))

	return data
}
