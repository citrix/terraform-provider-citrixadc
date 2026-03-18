package clusternodegroup_streamidentifier_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ClusternodegroupStreamidentifierBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"identifiername": schema.StringAttribute{
				Required:    true,
				Description: "stream identifier  and rate limit identifier that need to be bound to this nodegroup.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nodegroup to which you want to bind a cluster node or an entity.",
			},
		},
	}
}
