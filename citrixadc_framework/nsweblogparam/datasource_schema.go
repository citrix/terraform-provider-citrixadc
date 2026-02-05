package nsweblogparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NsweblogparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"buffersizemb": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Buffer size, in MB, allocated for log transaction data on the system. The maximum value is limited to the memory available on the system.",
			},
			"customreqhdrs": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name(s) of HTTP request headers whose values should be exported by the Web Logging feature.",
			},
			"customrsphdrs": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name(s) of HTTP response headers whose values should be exported by the Web Logging feature.",
			},
		},
	}
}
