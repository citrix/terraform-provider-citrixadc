package vpnepaprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnepaprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"data": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "deviceprofile data xml",
			},
			"filename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "filename of the deviceprofile data xml",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "name of device profile",
			},
		},
	}
}
