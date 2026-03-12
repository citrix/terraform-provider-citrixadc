package nssimpleacl6

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Nssimpleacl6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aclaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop incoming IPv6 packets that match the simple ACL6 rule.",
			},
			"aclname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the simple ACL6 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the simple ACL6 rule is created.",
			},
			"destport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number to match against the destination port number of an incoming IPv6 packet.\n\nDestPort is mandatory while setting Protocol. Omitting the port number and protocol creates an all-ports  and all protocol simple ACL6 rule, which matches any port and any protocol. In that case, you cannot create another simple ACL6 rule specifying a specific port and the same source IPv6 address.",
			},
			"estsessions": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol to match against the protocol of an incoming IPv6 packet. You must set this parameter if you set the Destination Port parameter.",
			},
			"srcipv6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address to match against the source IP address of an incoming IPv6 packet.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of seconds, in multiples of four, after which the simple ACL6 rule expires. If you do not want the simple ACL6 rule to expire, do not specify a TTL value.",
			},
		},
	}
}
