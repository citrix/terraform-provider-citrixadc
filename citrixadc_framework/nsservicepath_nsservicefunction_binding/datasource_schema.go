package nsservicepath_nsservicefunction_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsservicepathNsservicefunctionBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"index": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The serviceindex of each servicefunction in path.",
			},
			"servicefunction": schema.StringAttribute{
				Required:    true,
				Description: "List of service functions constituting the chain.",
			},
			"servicepathname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Service path. Must begin with an ASCII alphanumeric or underscore (_) character, and must\n      contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-)\n      characters.",
			},
		},
	}
}
