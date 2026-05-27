package appfwprotofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprotofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this gRPC schema file.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the gRPC schema object.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing gRPC schema object of the same name.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates source path of the gRPC schema file.",
			},
		},
	}
}
