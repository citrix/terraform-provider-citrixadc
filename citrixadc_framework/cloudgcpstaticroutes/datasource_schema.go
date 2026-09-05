package cloudgcpstaticroutes

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudgcpstaticroutesDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "status to push routes or not. Possible values = ENABLED, DISABLED",
			},
			"project": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "GCP project name for which static routes functionality is enabled.",
			},
		},
	}
}
