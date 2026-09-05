package nsacl6

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Nsacl6DataSourceModel is the data-source-specific model, decoupled from
// Nsacl6ResourceModel. A data source is a pure read surface, so it exposes the
// read/write attributes (as Computed outputs) plus the read-only (GET-only)
// attributes the resource deliberately omits.
type Nsacl6DataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Interface      types.String `tfsdk:"interface"`
	Acl6action     types.String `tfsdk:"acl6action"`
	Acl6name       types.String `tfsdk:"acl6name"`
	Aclaction      types.String `tfsdk:"aclaction"`
	Destipop       types.String `tfsdk:"destipop"`
	Destipv6       types.Bool   `tfsdk:"destipv6"`
	Destipv6val    types.String `tfsdk:"destipv6val"`
	Destport       types.Bool   `tfsdk:"destport"`
	Destportop     types.String `tfsdk:"destportop"`
	Destportval    types.String `tfsdk:"destportval"`
	Dfdhash        types.String `tfsdk:"dfdhash"`
	Dfdprefix      types.Int64  `tfsdk:"dfdprefix"`
	Established    types.Bool   `tfsdk:"established"`
	Icmpcode       types.Int64  `tfsdk:"icmpcode"`
	Icmptype       types.Int64  `tfsdk:"icmptype"`
	Logstate       types.String `tfsdk:"logstate"`
	Newname        types.String `tfsdk:"newname"`
	Nodeid         types.Int64  `tfsdk:"nodeid"`
	Priority       types.Int64  `tfsdk:"priority"`
	Protocol       types.String `tfsdk:"protocol"`
	Protocolnumber types.Int64  `tfsdk:"protocolnumber"`
	Ratelimit      types.Int64  `tfsdk:"ratelimit"`
	Srcipop        types.String `tfsdk:"srcipop"`
	Srcipv6        types.Bool   `tfsdk:"srcipv6"`
	Srcipv6val     types.String `tfsdk:"srcipv6val"`
	Srcmac         types.String `tfsdk:"srcmac"`
	Srcmacmask     types.String `tfsdk:"srcmacmask"`
	Srcport        types.Bool   `tfsdk:"srcport"`
	Srcportop      types.String `tfsdk:"srcportop"`
	Srcportval     types.String `tfsdk:"srcportval"`
	State          types.String `tfsdk:"state"`
	Stateful       types.String `tfsdk:"stateful"`
	Td             types.Int64  `tfsdk:"td"`
	Ttl            types.Int64  `tfsdk:"ttl"`
	Type           types.String `tfsdk:"type"`
	Vlan           types.Int64  `tfsdk:"vlan"`
	Vxlan          types.Int64  `tfsdk:"vxlan"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/nsacl6.json). Never settable; populated from GET.
	Kernelstate  types.String `tfsdk:"kernelstate"`
	Hits         types.Int64  `tfsdk:"hits"`
	Aclassociate types.List   `tfsdk:"aclassociate"`
}

func Nsacl6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of an interface. The Citrix ADC applies the ACL6 rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL6 rule to the incoming packets from all interfaces.",
			},
			"acl6action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform on the incoming IPv6 packets that match the ACL6 rule.\nAvailable settings function as follows:\n* ALLOW - The Citrix ADC processes the packet.\n* BRIDGE - The Citrix ADC bridges the packet to the destination without processing it.\n* DENY - The Citrix ADC drops the packet.",
			},
			"acl6name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the ACL6 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"aclaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action associated with the ACL6.",
			},
			"destipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"destipv6": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an incoming IPv6 packet.  In the command line interface, separate the range with a hyphen.",
			},
			"destipv6val": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination IPv6 address (range).",
			},
			"destport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the destination port number of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
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
			"dfdhash": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the type of hashmethod to be applied, to steer the packet to the FP of the packet.",
			},
			"dfdprefix": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "hashprefix to be applied to SIP/DIP to generate rsshash FP.eg 128 => hash calculated on the complete IP",
			},
			"established": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow only incoming TCP packets that have the ACK or RST bit set if the action set for the ACL6 rule is ALLOW and these packets match the other conditions in the ACL6 rule.",
			},
			"icmpcode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Code of a particular ICMP message type to match against the ICMP code of an incoming IPv6 ICMP packet.  For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code.\n\nIf you set this parameter, you must set the ICMP Type parameter.",
			},
			"icmptype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ICMP Message type to match against the message type of an incoming IPv6 ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type.\n\nNote: This parameter can be specified only for the ICMP protocol.",
			},
			"logstate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable logging of events related to the ACL6 rule. The log messages are stored in the configured syslog or auditlog server.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the ACL6 rule. Must begin with an ASCII alphabetic or underscore \\(_\\) character, and must contain only ASCII alphanumeric, underscore, hash \\(\\#\\), period \\(.\\), space, colon \\(:\\), at \\(@\\), equals \\(=\\), and hyphen \\(-\\) characters.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the NodeId to steer the packet to the provided FP.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority for the ACL6 rule, which determines the order in which it is evaluated relative to the other ACL6 rules. If you do not specify priorities while creating ACL6 rules, the ACL6 rules are evaluated in the order in which they are created.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, identified by protocol name, to match against the protocol of an incoming IPv6 packet.",
			},
			"protocolnumber": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, identified by protocol number, to match against the protocol of an incoming IPv6 packet.",
			},
			"ratelimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of log messages to be generated per second. If you set this parameter, you must enable the Log State parameter.",
			},
			"srcipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"srcipv6": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the source IP address of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen.",
			},
			"srcipv6val": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IPv6 address (range).",
			},
			"srcmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address to match against the source MAC address of an incoming IPv6 packet.",
			},
			"srcmacmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value \"000000111111\".",
			},
			"srcport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the source port number of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
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
				Description: "State of the ACL6.",
			},
			"stateful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If stateful option is enabled, transparent sessions are created for the traffic hitting this ACL6 and not hitting any other features like LB, INAT etc.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to expire this ACL6 (in seconds).",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Type of the acl6 ,default will be CLASSIC.\nAvailable options as follows:\n* CLASSIC - specifies the regular extended acls.\n* DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster .",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VLAN. The Citrix ADC applies the ACL6 rule only to the incoming packets on the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL6 rule to the incoming packets on all VLANs.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN. The Citrix ADC applies the ACL6 rule only to the incoming packets on the specified VXLAN. If you do not specify a VXLAN ID, the appliance applies the ACL6 rule to the incoming packets on all VXLANs.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"kernelstate": schema.StringAttribute{
				Computed:    true,
				Description: "Commit status of the ACL6 (APPLIED, NOTAPPLIED, RE-APPLY, ...).",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits of this ACL6.",
			},
			"aclassociate": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "ACL6 linked (NAT, FORWARDINGSESSION, NAT64, LSN).",
			},
		},
	}
}

// nsacl6DataSourceSetAttrFromGet projects a NITRO nsacl6 GET response onto the
// data-source model using the shared utils.MapGet* helpers.
func nsacl6DataSourceSetAttrFromGet(ctx context.Context, data *Nsacl6DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsacl6DataSourceSetAttrFromGet Function")

	// NITRO echoes the interface field under the capitalized "Interface" key.
	data.Interface = utils.MapGetString(g, "Interface")
	data.Acl6action = utils.MapGetString(g, "acl6action")
	data.Acl6name = utils.MapGetString(g, "acl6name")
	data.Aclaction = utils.MapGetString(g, "aclaction")
	data.Destipop = utils.MapGetString(g, "destipop")
	data.Destipv6 = utils.MapGetBool(g, "destipv6")
	data.Destipv6val = utils.MapGetString(g, "destipv6val")
	data.Destport = utils.MapGetBool(g, "destport")
	data.Destportop = utils.MapGetString(g, "destportop")
	data.Destportval = utils.MapGetString(g, "destportval")
	data.Dfdhash = utils.MapGetString(g, "dfdhash")
	data.Dfdprefix = utils.MapGetInt64(g, "dfdprefix")
	data.Established = utils.MapGetBool(g, "established")
	data.Icmpcode = utils.MapGetInt64(g, "icmpcode")
	data.Icmptype = utils.MapGetInt64(g, "icmptype")
	data.Logstate = utils.MapGetString(g, "logstate")
	// newname is rename-only and never echoed by GET.
	data.Newname = types.StringNull()
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Protocol = utils.MapGetString(g, "protocol")
	data.Protocolnumber = utils.MapGetInt64(g, "protocolnumber")
	data.Ratelimit = utils.MapGetInt64(g, "ratelimit")
	data.Srcipop = utils.MapGetString(g, "srcipop")
	data.Srcipv6 = utils.MapGetBool(g, "srcipv6")
	data.Srcipv6val = utils.MapGetString(g, "srcipv6val")
	data.Srcmac = utils.MapGetString(g, "srcmac")
	data.Srcmacmask = utils.MapGetString(g, "srcmacmask")
	data.Srcport = utils.MapGetBool(g, "srcport")
	data.Srcportop = utils.MapGetString(g, "srcportop")
	data.Srcportval = utils.MapGetString(g, "srcportval")
	data.State = utils.MapGetString(g, "state")
	data.Stateful = utils.MapGetString(g, "stateful")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Ttl = utils.MapGetInt64(g, "ttl")
	data.Type = utils.MapGetString(g, "type")
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Vxlan = utils.MapGetInt64(g, "vxlan")

	// Read-only (GET-only) attributes.
	data.Kernelstate = utils.MapGetString(g, "kernelstate")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Aclassociate = utils.MapGetStringList(g, "aclassociate")

	// Set ID from the acl6 name (matching the resource ID format).
	data.Id = types.StringValue(data.Acl6name.ValueString())
}
