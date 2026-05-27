package apiprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ApiprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"apivisibility": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable the schema lookup for the requests/apispecs that are bounded to the API profile. The default value of this parameter is DISABLED.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the API profile to add",
			},
		},
	}
}
