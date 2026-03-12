package lbgroup_lbvserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbgroupLbvserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the load balancing virtual server group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lbgroup\" or 'my lbgroup').",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Virtual server name.",
			},
		},
	}
}
