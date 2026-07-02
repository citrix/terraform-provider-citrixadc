package cloudallowedngsticketprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudallowedngsticketprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"creator": schema.StringAttribute{
				Computed:    true,
				Description: "Created name for allowed tickets",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Profile name for allowed tickets",
			},
		},
	}
}
