package lsnstatic

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnstaticDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"destip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination IP address for the LSN mapping entry.",
			},
			"dsttd": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the traffic domain through which the destination IP address for this LSN mapping entry is reachable from the Citrix ADC.\n\nIf you do not specify an ID, the destination IP address is assumed to be reachable through the default traffic domain, which has an ID of 0.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN static mapping entry. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn static1\" or 'lsn static1').",
			},
			"natip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 address, already existing on the Citrix ADC as type LSN, to be used as NAT IP address for this mapping entry.",
			},
			"natport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NAT port for this LSN mapping entry. * represents all ports being used. Used in case of static wildcard",
			},
			"network6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "B4 address in DS-Lite setup",
			},
			"subscrip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4(NAT44 & DS-Lite)/IPv6(NAT64) address of an LSN subscriber for the LSN static mapping entry.",
			},
			"subscrport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port of the LSN subscriber for the LSN mapping entry. * represents all ports being used. Used in case of static wildcard",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the traffic domain to which the subscriber belongs. \n\nIf you do not specify an ID, the subscriber is assumed to be a part of the default traffic domain.",
			},
			"transportprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol for the LSN mapping entry.",
			},
		},
	}
}
