package nsservicepath

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsservicepathDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"servicepathname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Service path. Must begin with an ASCII alphanumeric or underscore (_) character, and must\n      contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-)\n      characters.",
			},
		},
	}
}
