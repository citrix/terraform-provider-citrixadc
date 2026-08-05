package lbmetrictable

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbmetrictableDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"metrictable": schema.StringAttribute{
				Required:    true,
				Description: "Name for the metric table. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my metrictable\" or 'my metrictable').",
			},
		},
	}
}
