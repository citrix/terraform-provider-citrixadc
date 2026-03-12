package nat64param

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Nat64paramDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nat64fragheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When disabled, translator will not insert IPv6 fragmentation header for non fragmented IPv4 packets",
			},
			"nat64ignoretos": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore TOS.",
			},
			"nat64v6mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MTU setting for the IPv6 side. If the incoming IPv4 packet greater than this, either fragment or send icmp need fragmentation error.",
			},
			"nat64zerochecksum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Calculate checksum for UDP packets with zero checksum",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}
