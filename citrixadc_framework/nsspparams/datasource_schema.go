package nsspparams

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsspparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"basethreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of server connections that can be opened before surge protection is activated.",
			},
			"throttle": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rate at which the system opens connections to the server.",
			},
		},
	}
}
