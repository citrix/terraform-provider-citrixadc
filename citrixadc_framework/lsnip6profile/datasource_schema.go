package lsnip6profile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Lsnip6profileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN ip6 profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN ip6 profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn ip6 profile1\" or 'lsn ip6 profile1').",
			},
			"natprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 address(es) of the LSN subscriber(s) or subscriber network(s) on whose traffic you want the Citrix ADC to perform Large Scale NAT.",
			},
			"network6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 address of the Citrix ADC AFTR device",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 translation type for which to set the LSN IP6 profile parameters.",
			},
		},
	}
}
