package appfwwsdl

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwwsdlDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the WSDL.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the WSDL file to remove.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing WSDL of the same name.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and name) of the WSDL file to be imported is stored.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}
