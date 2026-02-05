package appfwjsonerrorpage

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwjsonerrorpageDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the JSON error object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Indicates name of the imported json error page to be removed.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing JSON error object of the same name.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and name) for the location at which to store the imported JSON error object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}
