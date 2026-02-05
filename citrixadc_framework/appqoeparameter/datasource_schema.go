package appqoeparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppqoeparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"avgwaitingclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "average number of client connections, that can sit in service waiting queue",
			},
			"dosattackthresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "average number of client connection that can queue up on vserver level without triggering DoS mitigation module",
			},
			"maxaltrespbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "maximum bandwidth which will determine whether to send alternate content response",
			},
			"sessionlife": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, between the first time and the next time the AppQoE alternative content window is displayed. The alternative content window is displayed only once during a session for the same browser accessing a configured URL, so this parameter determines the length of a session.",
			},
		},
	}
}
