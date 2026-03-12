package clusternodegroup_clusternode_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ClusternodegroupClusternodeBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
			"node": schema.Int64Attribute{
				Required:    true,
				Description: "Nodes in the nodegroup",
			},
		},
	}
}
