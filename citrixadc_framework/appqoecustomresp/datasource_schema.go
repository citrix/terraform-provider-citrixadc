package appqoecustomresp

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppqoecustomrespDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Indicates name of the custom response HTML page to import/update.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
		},
	}
}
