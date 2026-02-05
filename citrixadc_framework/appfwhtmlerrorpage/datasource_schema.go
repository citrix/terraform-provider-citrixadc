package appfwhtmlerrorpage

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwhtmlerrorpageDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the HTML error object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the XML error object to remove.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing HTML error object of the same name.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and name) for the location at which to store the imported HTML error object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}
