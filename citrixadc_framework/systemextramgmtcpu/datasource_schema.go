package systemextramgmtcpu

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemextramgmtcpuDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Boolean value indicating the effective state of the extra management CPU.",
			},
		},
	}
}
