package vridparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VridparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"deadinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of seconds after which a peer node in active-active mode is marked down if vrrp advertisements are not received from the peer node.",
			},
			"hellointerval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in milliseconds, between vrrp advertisement messages sent to the peer node in active-active mode.",
			},
			"sendtomaster": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Forward packets to the master node, in an active-active mode configuration, if the virtual server is in the backup state and sharing is disabled.",
			},
		},
	}
}
