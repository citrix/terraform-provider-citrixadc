package rnat

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RnatDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aclname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An extended ACL defined for the RNAT entry.",
			},
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Synchronize all connection-related information for the RNAT sessions with the secondary ADC in a high availability (HA) pair.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the RNAT4 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the RNAT4 rule.",
			},
			"natip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The subnet mask for the network address.",
			},
			"network": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The network address defined for the RNAT entry.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the RNAT4 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain       only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this rnat rule.",
			},
			"redirectport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number to which the IPv4 packets are redirected. Applicable to TCP and UDP protocols.",
			},
			"srcippersistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables the Citrix ADC to use the same NAT IP address for all RNAT sessions initiated from a particular server.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable source port proxying, which enables the Citrix ADC to use the RNAT ips using proxied source port.",
			},
		},
	}
}
