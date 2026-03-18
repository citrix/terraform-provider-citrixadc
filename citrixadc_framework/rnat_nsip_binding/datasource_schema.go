package rnat_nsip_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RnatNsipBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the RNAT rule to which to bind NAT IPs.",
			},
			"natip": schema.StringAttribute{
				Required:    true,
				Description: "Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range.",
			},
		},
	}
}
