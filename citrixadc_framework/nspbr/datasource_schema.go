package nspbr

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NspbrDataSourceModel is the data-source-specific model, decoupled from
// NspbrResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime/statistics attributes the resource
// deliberately omits (hits, kernelstate, curstate, probe counters, ...). Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type NspbrDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Interface       types.String `tfsdk:"interface"`
	Action          types.String `tfsdk:"action"`
	Destip          types.Bool   `tfsdk:"destip"`
	Destipdataset   types.String `tfsdk:"destipdataset"`
	Destipop        types.String `tfsdk:"destipop"`
	Destipval       types.String `tfsdk:"destipval"`
	Destport        types.Bool   `tfsdk:"destport"`
	Destportdataset types.String `tfsdk:"destportdataset"`
	Destportop      types.String `tfsdk:"destportop"`
	Destportval     types.String `tfsdk:"destportval"`
	Detail          types.Bool   `tfsdk:"detail"`
	Iptunnel        types.Bool   `tfsdk:"iptunnel"`
	Iptunnelname    types.String `tfsdk:"iptunnelname"`
	Monitor         types.String `tfsdk:"monitor"`
	Msr             types.String `tfsdk:"msr"`
	Name            types.String `tfsdk:"name"` // Required lookup key
	Nexthop         types.Bool   `tfsdk:"nexthop"`
	Nexthopval      types.String `tfsdk:"nexthopval"`
	Ownergroup      types.String `tfsdk:"ownergroup"`
	Priority        types.Int64  `tfsdk:"priority"`
	Protocol        types.String `tfsdk:"protocol"`
	Protocolnumber  types.Int64  `tfsdk:"protocolnumber"`
	Srcip           types.Bool   `tfsdk:"srcip"`
	Srcipdataset    types.String `tfsdk:"srcipdataset"`
	Srcipop         types.String `tfsdk:"srcipop"`
	Srcipval        types.String `tfsdk:"srcipval"`
	Srcmac          types.String `tfsdk:"srcmac"`
	Srcmacmask      types.String `tfsdk:"srcmacmask"`
	Srcport         types.Bool   `tfsdk:"srcport"`
	Srcportdataset  types.String `tfsdk:"srcportdataset"`
	Srcportop       types.String `tfsdk:"srcportop"`
	Srcportval      types.String `tfsdk:"srcportval"`
	State           types.String `tfsdk:"state"`
	Targettd        types.Int64  `tfsdk:"targettd"`
	Td              types.Int64  `tfsdk:"td"`
	Vlan            types.Int64  `tfsdk:"vlan"`
	Vxlan           types.Int64  `tfsdk:"vxlan"`
	Vxlanvlanmap    types.String `tfsdk:"vxlanvlanmap"`

	// Read-only (GET-only) runtime/statistics metadata from the NITRO doc
	// read-only set (zion73x_readonly/nspbr.json). Never settable; from GET.
	Hits              types.Int64  `tfsdk:"hits"`
	Kernelstate       types.String `tfsdk:"kernelstate"`
	Curstate          types.Int64  `tfsdk:"curstate"`
	Totalprobes       types.Int64  `tfsdk:"totalprobes"`
	Totalfailedprobes types.Int64  `tfsdk:"totalfailedprobes"`
	Failedprobes      types.Int64  `tfsdk:"failedprobes"`
	Monstatcode       types.Int64  `tfsdk:"monstatcode"`
	Monstatparam1     types.Int64  `tfsdk:"monstatparam1"`
	Monstatparam2     types.Int64  `tfsdk:"monstatparam2"`
	Monstatparam3     types.Int64  `tfsdk:"monstatparam3"`
	Data              types.Bool   `tfsdk:"data"`
	Pbrchildcount     types.Int64  `tfsdk:"pbrchildcount"`
}

func NspbrDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of an interface. The Citrix ADC compares the PBR only to the outgoing packets on the specified interface. If you do not specify any value, the appliance compares the PBR to the outgoing packets on all interfaces.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform on the outgoing IPv4 packets that match the PBR.\n\nAvailable settings function as follows:\n* ALLOW - The Citrix ADC sends the packet to the designated next-hop router.\n* DENY - The Citrix ADC applies the routing table for normal destination-based routing.",
			},
			"destip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an outgoing IPv4 packet.  In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.",
			},
			"destipdataset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Policy dataset which can have multiple IP ranges bound to it.",
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
			"destportdataset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Policy dataset which can have multiple port ranges bound to it.",
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
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
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
				Optional:    true,
				Computed:    true,
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
			"srcipdataset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Policy dataset which can have multiple IP ranges bound to it.",
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
				Computed:    true,
				Description: "Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value \"000000111111\".",
			},
			"srcport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the source port number of an outgoing IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
			},
			"srcportdataset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Policy dataset which can have multiple port ranges bound to it.",
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
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the PBR. After you apply the PBRs, the Citrix ADC compares outgoing packets to the enabled PBRs.",
			},
			"targettd": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain to which you want to send packet to.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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

			// Read-only (GET-only) runtime/statistics metadata surfaced by the
			// data source (these are intentionally NOT modeled on the resource).
			// All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The hits of this PBR.",
			},
			"kernelstate": schema.StringAttribute{
				Computed:    true,
				Description: "The commit status of the PBR. Possible values = APPLIED, NOTAPPLIED, RE-APPLY, SFAPPLIED, SFNOTAPPLIED, SFAPPLIED61, SFNOTAPPLIED61.",
			},
			"curstate": schema.Int64Attribute{
				Computed:    true,
				Description: "If this route is UP/DOWN.",
			},
			"totalprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of probes sent.",
			},
			"totalfailedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of failed probes.",
			},
			"failedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of the current failed monitoring probes.",
			},
			"monstatcode": schema.Int64Attribute{
				Computed:    true,
				Description: "The code indicating the monitor response.",
			},
			"monstatparam1": schema.Int64Attribute{
				Computed:    true,
				Description: "First parameter for use with message code.",
			},
			"monstatparam2": schema.Int64Attribute{
				Computed:    true,
				Description: "Second parameter for use with message code.",
			},
			"monstatparam3": schema.Int64Attribute{
				Computed:    true,
				Description: "Third parameter for use with message code.",
			},
			"data": schema.BoolAttribute{
				Computed:    true,
				Description: "Internal data of this route.",
			},
			"pbrchildcount": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of childs for this PBR, in case of dataset this will be number of ips bound to that dataset.",
			},
		},
	}
}

// nspbrDataSourceSetAttrFromGet projects a NITRO nspbr GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func nspbrDataSourceSetAttrFromGet(ctx context.Context, data *NspbrDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nspbrDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs. NITRO returns the interface
	// under the "Interface" key (capitalized).
	data.Interface = utils.MapGetString(g, "Interface")
	data.Action = utils.MapGetString(g, "action")
	data.Destip = utils.MapGetBool(g, "destip")
	data.Destipdataset = utils.MapGetString(g, "destipdataset")
	data.Destipop = utils.MapGetString(g, "destipop")
	data.Destipval = utils.MapGetString(g, "destipval")
	data.Destport = utils.MapGetBool(g, "destport")
	data.Destportdataset = utils.MapGetString(g, "destportdataset")
	data.Destportop = utils.MapGetString(g, "destportop")
	data.Destportval = utils.MapGetString(g, "destportval")
	data.Detail = utils.MapGetBool(g, "detail")
	data.Iptunnel = utils.MapGetBool(g, "iptunnel")
	data.Iptunnelname = utils.MapGetString(g, "iptunnelname")
	data.Monitor = utils.MapGetString(g, "monitor")
	data.Msr = utils.MapGetString(g, "msr")
	data.Nexthop = utils.MapGetBool(g, "nexthop")
	data.Nexthopval = utils.MapGetString(g, "nexthopval")
	data.Ownergroup = utils.MapGetString(g, "ownergroup")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Protocol = utils.MapGetString(g, "protocol")
	data.Protocolnumber = utils.MapGetInt64(g, "protocolnumber")
	data.Srcip = utils.MapGetBool(g, "srcip")
	data.Srcipdataset = utils.MapGetString(g, "srcipdataset")
	data.Srcipop = utils.MapGetString(g, "srcipop")
	data.Srcipval = utils.MapGetString(g, "srcipval")
	data.Srcmac = utils.MapGetString(g, "srcmac")
	data.Srcmacmask = utils.MapGetString(g, "srcmacmask")
	data.Srcport = utils.MapGetBool(g, "srcport")
	data.Srcportdataset = utils.MapGetString(g, "srcportdataset")
	data.Srcportop = utils.MapGetString(g, "srcportop")
	data.Srcportval = utils.MapGetString(g, "srcportval")
	data.State = utils.MapGetString(g, "state")
	data.Targettd = utils.MapGetInt64(g, "targettd")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Vxlan = utils.MapGetInt64(g, "vxlan")
	data.Vxlanvlanmap = utils.MapGetString(g, "vxlanvlanmap")

	// Read-only runtime/statistics metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Kernelstate = utils.MapGetString(g, "kernelstate")
	data.Curstate = utils.MapGetInt64(g, "curstate")
	data.Totalprobes = utils.MapGetInt64(g, "totalprobes")
	data.Totalfailedprobes = utils.MapGetInt64(g, "totalfailedprobes")
	data.Failedprobes = utils.MapGetInt64(g, "failedprobes")
	data.Monstatcode = utils.MapGetInt64(g, "monstatcode")
	data.Monstatparam1 = utils.MapGetInt64(g, "monstatparam1")
	data.Monstatparam2 = utils.MapGetInt64(g, "monstatparam2")
	data.Monstatparam3 = utils.MapGetInt64(g, "monstatparam3")
	data.Data = utils.MapGetBool(g, "data")
	data.Pbrchildcount = utils.MapGetInt64(g, "pbrchildcount")
}
