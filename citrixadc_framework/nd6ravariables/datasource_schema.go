package nd6ravariables

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Nd6ravariablesDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ceaserouteradv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cease router advertisements on this vlan.",
			},
			"currhoplimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Current Hop limit.",
			},
			"defaultlifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Default life time, in seconds.",
			},
			"linkmtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The Link MTU.",
			},
			"managedaddrconfig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value to be placed in the Managed address configuration flag field.",
			},
			"maxrtadvinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time allowed between unsolicited multicast RAs, in seconds.",
			},
			"minrtadvinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum time interval between RA messages, in seconds.",
			},
			"onlyunicastrtadvresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send only Unicast Router Advertisements in respond to Router Solicitations.",
			},
			"otheraddrconfig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value to be placed in the Other configuration flag field.",
			},
			"reachabletime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Reachable time, in milliseconds.",
			},
			"retranstime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Retransmission time, in milliseconds.",
			},
			"sendrouteradv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "whether the router sends periodic RAs and responds to Router Solicitations.",
			},
			"srclinklayeraddroption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include source link layer address option in RA messages.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "The VLAN number.",
			},
		},
	}
}
