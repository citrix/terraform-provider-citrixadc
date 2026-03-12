package dnssuffix

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnssuffixDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dnssuffix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Suffix to be appended when resolving domain names that are not fully qualified.",
			},
		},
	}
}
