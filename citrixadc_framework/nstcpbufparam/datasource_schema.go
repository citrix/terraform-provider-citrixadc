package nstcpbufparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NstcpbufparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum memory, in megabytes, that can be used for buffering.",
			},
			"size": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP buffering size per connection, in kilobytes.",
			},
		},
	}
}
