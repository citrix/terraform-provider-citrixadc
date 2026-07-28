package nschannelparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NschannelparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vfautorecover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "VF autorecover mode",
			},
		},
	}
}
