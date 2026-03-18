package vpnportaltheme

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnportalthemeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"basetheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the uitheme",
			},
		},
	}
}
