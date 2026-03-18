package nsmode

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsmodeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsmode datasource.",
			},
			"fr": schema.BoolAttribute{
				Computed:    true,
				Description: "Fast Ramp mode.",
			},
			"l2": schema.BoolAttribute{
				Computed:    true,
				Description: "Layer 2 mode.",
			},
			"usip": schema.BoolAttribute{
				Computed:    true,
				Description: "Use Source IP mode.",
			},
			"cka": schema.BoolAttribute{
				Computed:    true,
				Description: "Client Keep-Alive mode.",
			},
			"tcpb": schema.BoolAttribute{
				Computed:    true,
				Description: "TCP Buffering mode.",
			},
			"mbf": schema.BoolAttribute{
				Computed:    true,
				Description: "MAC-based forwarding mode.",
			},
			"edge": schema.BoolAttribute{
				Computed:    true,
				Description: "Edge configuration mode.",
			},
			"usnip": schema.BoolAttribute{
				Computed:    true,
				Description: "Use Subnet IP mode.",
			},
			"l3": schema.BoolAttribute{
				Computed:    true,
				Description: "Layer 3 mode.",
			},
			"pmtud": schema.BoolAttribute{
				Computed:    true,
				Description: "Path MTU Discovery mode.",
			},
			"mediaclassification": schema.BoolAttribute{
				Computed:    true,
				Description: "Media classification mode.",
			},
			"sradv": schema.BoolAttribute{
				Computed:    true,
				Description: "Static route advertisement mode.",
			},
			"dradv": schema.BoolAttribute{
				Computed:    true,
				Description: "Dynamic route advertisement mode.",
			},
			"iradv": schema.BoolAttribute{
				Computed:    true,
				Description: "Intranet route advertisement mode.",
			},
			"sradv6": schema.BoolAttribute{
				Computed:    true,
				Description: "IPv6 static route advertisement mode.",
			},
			"dradv6": schema.BoolAttribute{
				Computed:    true,
				Description: "IPv6 dynamic route advertisement mode.",
			},
			"bridgebpdus": schema.BoolAttribute{
				Computed:    true,
				Description: "Bridge BPDUs mode.",
			},
			"ulfd": schema.BoolAttribute{
				Computed:    true,
				Description: "Use Layer 2 mode for IPv4 packets.",
			},
		},
	}
}
