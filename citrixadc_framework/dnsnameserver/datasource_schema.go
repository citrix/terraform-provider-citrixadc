package dnsnameserver

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnsnameserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the name server",
			},
			"dnsvservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of a DNS virtual server. Overrides any IP address-based name servers configured on the Citrix ADC. Either dnsvservername or ip must be specified.",
			},
			"ip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of an external name server or, if the Local parameter is set, IP address of a local DNS server (LDNS). Either dnsvservername or ip must be specified.",
			},
			"local": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark the IP address as one that belongs to a local recursive DNS server on the Citrix ADC. The appliance recursively resolves queries received on an IP address that is marked as being local. For recursive resolution to work, the global DNS parameter, Recursion, must also be set.\n\nIf no name server is marked as being local, the appliance functions as a stub resolver and load balances the name servers.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Administrative state of the name server.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by the name server. UDP_TCP is not valid if the name server is a DNS virtual server configured on the appliance.",
			},
		},
	}
}
