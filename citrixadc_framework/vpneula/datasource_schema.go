package vpneula

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpneulaDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the eula",
			},
		},
	}
}
