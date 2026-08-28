package hasecureheartbeats

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func HasecureheartbeatsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By enabling this option, HA heartbeats are securely exchanged between nodes. Possible values = ENABLED, DISABLED",
			},
			"hapsk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Pre shared key to be used for securing HA heartbeats.",
			},
			"hapsk_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Pre shared key to be used for securing HA heartbeats.",
			},
			"hapsk_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a hapsk_wo update.",
			},
		},
	}
}
