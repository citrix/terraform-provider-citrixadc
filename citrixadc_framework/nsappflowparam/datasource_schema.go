package nsappflowparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsappflowparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clienttrafficonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control whether AppFlow records should be generated only for client-side traffic.",
			},
			"httpcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP cookie logging.",
			},
			"httphost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP host logging.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP method logging.",
			},
			"httpreferer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP referer logging.",
			},
			"httpurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP URL logging.",
			},
			"httpuseragent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP user-agent logging.",
			},
			"templaterefresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IPFIX template refresh interval (in seconds).",
			},
			"udppmtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MTU to be used for IPFIX UDP packets.",
			},
		},
	}
}
