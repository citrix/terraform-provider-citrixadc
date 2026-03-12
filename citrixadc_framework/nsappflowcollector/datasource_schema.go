package nsappflowcollector

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsappflowcollectorDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IPv4 address of the AppFlow collector.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the AppFlow collector.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The UDP port on which the AppFlow collector is listening.",
			},
		},
	}
}
