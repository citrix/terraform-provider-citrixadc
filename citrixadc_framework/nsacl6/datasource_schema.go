package nsacl6

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

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
		},
	}
}
