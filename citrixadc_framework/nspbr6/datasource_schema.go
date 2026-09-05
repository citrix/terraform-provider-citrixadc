package nspbr6

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Nspbr6DataSourceModel is the data-source-specific model, decoupled from
// Nspbr6ResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime/statistics attributes the resource
// deliberately omits (kernelstate, hits, curstate, probe counters, ...). Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type Nspbr6DataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Interface      types.String `tfsdk:"interface"`
	Action         types.String `tfsdk:"action"`
	Destipop       types.String `tfsdk:"destipop"`
	Destipv6       types.Bool   `tfsdk:"destipv6"`
	Destipv6val    types.String `tfsdk:"destipv6val"`
	Destport       types.Bool   `tfsdk:"destport"`
	Destportop     types.String `tfsdk:"destportop"`
	Destportval    types.String `tfsdk:"destportval"`
	Detail         types.Bool   `tfsdk:"detail"` // Required query flag
	Iptunnel       types.String `tfsdk:"iptunnel"`
	Monitor        types.String `tfsdk:"monitor"`
	Msr            types.String `tfsdk:"msr"`
	Name           types.String `tfsdk:"name"` // Required lookup key
	Nexthop        types.Bool   `tfsdk:"nexthop"`
	Nexthopval     types.String `tfsdk:"nexthopval"`
	Nexthopvlan    types.Int64  `tfsdk:"nexthopvlan"`
	Ownergroup     types.String `tfsdk:"ownergroup"`
	Priority       types.Int64  `tfsdk:"priority"`
	Protocol       types.String `tfsdk:"protocol"`
	Protocolnumber types.Int64  `tfsdk:"protocolnumber"`
	Srcipop        types.String `tfsdk:"srcipop"`
	Srcipv6        types.Bool   `tfsdk:"srcipv6"`
	Srcipv6val     types.String `tfsdk:"srcipv6val"`
	Srcmac         types.String `tfsdk:"srcmac"`
	Srcmacmask     types.String `tfsdk:"srcmacmask"`
	Srcport        types.Bool   `tfsdk:"srcport"`
	Srcportop      types.String `tfsdk:"srcportop"`
	Srcportval     types.String `tfsdk:"srcportval"`
	State          types.String `tfsdk:"state"`
	Td             types.Int64  `tfsdk:"td"`
	Vlan           types.Int64  `tfsdk:"vlan"`
	Vxlan          types.Int64  `tfsdk:"vxlan"`
	Vxlanvlanmap   types.String `tfsdk:"vxlanvlanmap"`

	// Read-only (GET-only) runtime/statistics metadata from the NITRO doc
	// read-only set (zion73x_readonly/nspbr6.json). Never settable; from GET.
	Kernelstate       types.String `tfsdk:"kernelstate"`
	Hits              types.Int64  `tfsdk:"hits"`
	Curstate          types.Int64  `tfsdk:"curstate"`
	Totalprobes       types.Int64  `tfsdk:"totalprobes"`
	Totalfailedprobes types.Int64  `tfsdk:"totalfailedprobes"`
	Failedprobes      types.Int64  `tfsdk:"failedprobes"`
	Monstatcode       types.Int64  `tfsdk:"monstatcode"`
	Monstatparam1     types.Int64  `tfsdk:"monstatparam1"`
	Monstatparam2     types.Int64  `tfsdk:"monstatparam2"`
	Monstatparam3     types.Int64  `tfsdk:"monstatparam3"`
	Data              types.Bool   `tfsdk:"data"`
}

func Nspbr6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of an interface. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified interface. If you do not specify a value, the appliance compares the PBR6 to the outgoing packets on all interfaces.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform on the outgoing IPv6 packets that match the PBR6.\n\nAvailable settings function as follows:\n* ALLOW - The Citrix ADC sends the packet to the designated next-hop router.\n* DENY - The Citrix ADC applies the routing table for normal destination-based routing.",
			},
			"destipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"destipv6": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an outgoing IPv6 packet.  In the command line interface, separate the range with a hyphen.",
			},
			"destipv6val": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an outgoing IPv6 packet.  In the command line interface, separate the range with a hyphen.",
			},
			"destport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the destination port number of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
			},
			"destportop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"destportval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination port (range).",
			},
			"detail": schema.BoolAttribute{
				Required:    true,
				Description: "To get a detailed view.",
			},
			"iptunnel": schema.StringAttribute{
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
				Description: "Monitor the route specified by the Next Hop parameter.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the PBR6. Must begin with an ASCII alphabetic or underscore \\(_\\) character, and must contain only ASCII alphanumeric, underscore, hash \\(\\#\\), period \\(.\\), space, colon \\(:\\), at \\(@\\), equals \\(=\\), and hyphen \\(-\\) characters. Cannot be changed after the PBR6 is created.",
			},
			"nexthop": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the next hop router to which to send matching packets if action is set to ALLOW. This next hop should be directly reachable from the appliance.",
			},
			"nexthopval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Next Hop IPv6 address.",
			},
			"nexthopvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN number to be used for link local nexthop .",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this pbr rule. If owner node group is not specified then the pbr rule is treated as Striped pbr rule.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the PBR6, which determines the order in which it is evaluated relative to the other PBR6s. If you do not specify priorities while creating PBR6s, the PBR6s are evaluated in the order in which they are created.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, identified by protocol name, to match against the protocol of an outgoing IPv6 packet.",
			},
			"protocolnumber": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, identified by protocol number, to match against the protocol of an outgoing IPv6 packet.",
			},
			"srcipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"srcipv6": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the source IP address of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen.",
			},
			"srcipv6val": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the source IP address of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen.",
			},
			"srcmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address to match against the source MAC address of an outgoing IPv6 packet.",
			},
			"srcmacmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value \"000000111111\".",
			},
			"srcport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the source port number of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.",
			},
			"srcportop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"srcportval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source port (range).",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the PBR6. After you apply the PBR6s, the Citrix ADC compares outgoing packets to the enabled PBR6s.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VLAN. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified VLAN. If you do not specify an interface ID, the appliance compares the PBR6 to the outgoing packets on all VLANs.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified VXLAN. If you do not specify an interface ID, the appliance compares the PBR6 to the outgoing packets on all VXLANs.",
			},
			"vxlanvlanmap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The vlan to vxlan mapping to be applied for incoming packets over this pbr tunnel.",
			},

			// Read-only (GET-only) runtime/statistics metadata surfaced by the
			// data source (these are intentionally NOT modeled on the resource).
			// All Computed.
			"kernelstate": schema.StringAttribute{
				Computed:    true,
				Description: "Commit status of the PBR6. Possible values = APPLIED, NOTAPPLIED, RE-APPLY, SFAPPLIED, SFNOTAPPLIED.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits of this PBR6.",
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
		},
	}
}

// nspbr6DataSourceSetAttrFromGet projects a NITRO nspbr6 GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func nspbr6DataSourceSetAttrFromGet(ctx context.Context, data *Nspbr6DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nspbr6DataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs. NITRO returns the interface
	// under the "Interface" key (capitalized).
	data.Interface = utils.MapGetString(g, "Interface")
	data.Action = utils.MapGetString(g, "action")
	data.Destipop = utils.MapGetString(g, "destipop")
	data.Destipv6 = utils.MapGetBool(g, "destipv6")
	data.Destipv6val = utils.MapGetString(g, "destipv6val")
	data.Destport = utils.MapGetBool(g, "destport")
	data.Destportop = utils.MapGetString(g, "destportop")
	data.Destportval = utils.MapGetString(g, "destportval")
	data.Iptunnel = utils.MapGetString(g, "iptunnel")
	data.Monitor = utils.MapGetString(g, "monitor")
	data.Msr = utils.MapGetString(g, "msr")
	data.Nexthop = utils.MapGetBool(g, "nexthop")
	data.Nexthopval = utils.MapGetString(g, "nexthopval")
	data.Nexthopvlan = utils.MapGetInt64(g, "nexthopvlan")
	data.Ownergroup = utils.MapGetString(g, "ownergroup")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Protocol = utils.MapGetString(g, "protocol")
	data.Protocolnumber = utils.MapGetInt64(g, "protocolnumber")
	data.Srcipop = utils.MapGetString(g, "srcipop")
	data.Srcipv6 = utils.MapGetBool(g, "srcipv6")
	data.Srcipv6val = utils.MapGetString(g, "srcipv6val")
	data.Srcmac = utils.MapGetString(g, "srcmac")
	data.Srcmacmask = utils.MapGetString(g, "srcmacmask")
	data.Srcport = utils.MapGetBool(g, "srcport")
	data.Srcportop = utils.MapGetString(g, "srcportop")
	data.Srcportval = utils.MapGetString(g, "srcportval")
	data.State = utils.MapGetString(g, "state")
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

	// detail is a Required query-only input the GET never echoes; preserve the
	// value supplied in configuration (already present in data) rather than
	// clobbering it to Null.

	// Read-only runtime/statistics metadata.
	data.Kernelstate = utils.MapGetString(g, "kernelstate")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Curstate = utils.MapGetInt64(g, "curstate")
	data.Totalprobes = utils.MapGetInt64(g, "totalprobes")
	data.Totalfailedprobes = utils.MapGetInt64(g, "totalfailedprobes")
	data.Failedprobes = utils.MapGetInt64(g, "failedprobes")
	data.Monstatcode = utils.MapGetInt64(g, "monstatcode")
	data.Monstatparam1 = utils.MapGetInt64(g, "monstatparam1")
	data.Monstatparam2 = utils.MapGetInt64(g, "monstatparam2")
	data.Monstatparam3 = utils.MapGetInt64(g, "monstatparam3")
	data.Data = utils.MapGetBool(g, "data")
}
