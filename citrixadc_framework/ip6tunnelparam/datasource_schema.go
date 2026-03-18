package ip6tunnelparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Ip6tunnelparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dropfrag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop any packet that requires fragmentation.",
			},
			"dropfragcputhreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold value, as a percentage of CPU usage, at which to drop packets that require fragmentation. Applies only if dropFragparameter is set to NO.",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Common source IPv6 address for all IPv6 tunnels. Must be a SNIP6 or VIP6 address.",
			},
			"srciproundrobin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use a different source IPv6 address for each new session through a particular IPv6 tunnel, as determined by round robin selection of one of the SNIP6 addresses. This setting is ignored if a common global source IPv6 address has been specified for all the IPv6 tunnels. This setting does not apply to a tunnel for which a source IPv6 address has been specified.",
			},
			"useclientsourceipv6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use client source IPv6 address as source IPv6 address for outer tunnel IPv6 header",
			},
		},
	}
}
