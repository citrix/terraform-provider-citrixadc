package cloudparaminternal

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudparaminternalDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nonftumode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates if GUI in in FTU mode or not",
			},
		},
	}
}