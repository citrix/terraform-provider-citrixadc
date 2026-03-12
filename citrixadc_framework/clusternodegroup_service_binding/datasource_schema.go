package clusternodegroup_service_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ClusternodegroupServiceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
			"service": schema.StringAttribute{
				Required:    true,
				Description: "name of the service bound to this nodegroup.",
			},
		},
	}
}
