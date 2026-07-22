package systemcpuparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemcpuparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"pemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set PEmode to DEFAULT/CPUBOUND. Distribute the PE weights equally if PEmode is set to CPUBOUND.",
			},
		},
	}
}
