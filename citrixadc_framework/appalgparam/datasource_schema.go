package appalgparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppalgparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"pptpgreidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval in sec, after which data sessions of PPTP GRE is cleared.",
			},
		},
	}
}
