package iptunnelparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IptunnelparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dropfrag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop any IP packet that requires fragmentation before it is sent through the tunnel.",
			},
			"dropfragcputhreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold value, as a percentage of CPU usage, at which to drop packets that require fragmentation to use the IP tunnel. Applies only if dropFragparameter is set to NO. The default value, 0, specifies that this parameter is not set.",
			},
			"enablestrictrx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Strict PBR check for IPSec packets received through tunnel",
			},
			"enablestricttx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Strict PBR check for packets to be sent IPSec protected",
			},
			"mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The shared MAC used for shared IP between cluster nodes/HA peers",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Common source-IP address for all tunnels. For a specific tunnel, this global setting is overridden if you have specified another source IP address. Must be a MIP or SNIP address.",
			},
			"srciproundrobin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use a different source IP address for each new session through a particular IP tunnel, as determined by round robin selection of one of the SNIP addresses. This setting is ignored if a common global source IP address has been specified for all the IP tunnels. This setting does not apply to a tunnel for which a source IP address has been specified.",
			},
			"useclientsourceip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use client source IP as source IP for outer tunnel IP header",
			},
		},
	}
}
