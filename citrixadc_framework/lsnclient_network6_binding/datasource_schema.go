package lsnclient_network6_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnclientNetwork6BindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn client1\" or 'lsn client1').",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask for the IPv4 address specified in the Network parameter.",
			},
			"network": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 address(es) of the LSN subscriber(s) or subscriber network(s) on whose traffic you want the Citrix ADC to perform Large Scale NAT.",
			},
			"network6": schema.StringAttribute{
				Required:    true,
				Description: "IPv6 address(es) of the LSN subscriber(s) or subscriber network(s) on whose traffic you want the Citrix ADC to perform Large Scale NAT.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs. \nIf you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.",
			},
		},
	}
}
