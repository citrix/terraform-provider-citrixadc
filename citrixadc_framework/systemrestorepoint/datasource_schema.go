package systemrestorepoint

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemrestorepointDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"filename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the restore point.",
			},
		},
	}
}
