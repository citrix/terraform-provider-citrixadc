package onlinkipv6prefix

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Onlinkipv6prefixDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"autonomusprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RA Prefix Autonomus flag.",
			},
			"decrementprefixlifetimes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RA Prefix Autonomus flag.",
			},
			"depricateprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Depricate the prefix.",
			},
			"ipv6prefix": schema.StringAttribute{
				Required:    true,
				Description: "Onlink prefixes for RA messages.",
			},
			"onlinkprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RA Prefix onlink flag.",
			},
			"prefixpreferredlifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Preferred life time of the prefix, in seconds.",
			},
			"prefixvalidelifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Valide life time of the prefix, in seconds.",
			},
		},
	}
}
