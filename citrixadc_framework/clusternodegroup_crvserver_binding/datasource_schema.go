package clusternodegroup_crvserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ClusternodegroupCrvserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
			"vserver": schema.StringAttribute{
				Required:    true,
				Description: "vserver that need to be bound to this nodegroup.",
			},
		},
	}
}
