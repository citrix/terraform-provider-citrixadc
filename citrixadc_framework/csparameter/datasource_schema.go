package csparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CsparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"stateupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether the virtual server checks the attached load balancing server for state information.",
			},
		},
	}
}
