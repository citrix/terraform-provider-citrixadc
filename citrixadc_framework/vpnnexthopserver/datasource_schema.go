package vpnnexthopserver

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnnexthopserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Citrix Gateway appliance in the first DMZ.",
			},
			"nexthopfqdn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FQDN of the Citrix Gateway proxy in the second DMZ.",
			},
			"nexthopip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the Citrix Gateway proxy in the second DMZ.",
			},
			"nexthopport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number of the Citrix Gateway proxy in the second DMZ.",
			},
			"resaddresstype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Address Type (IPV4/IPv6) of DNS name of nextHopServer FQDN.",
			},
			"secure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use of a secure port, such as 443, for the double-hop configuration.",
			},
		},
	}
}
