package clusternode_routemonitor_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ClusternodeRoutemonitorBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"netmask": schema.StringAttribute{
				Required:    true,
				Description: "The netmask.",
			},
			"nodeid": schema.Int64Attribute{
				Required:    true,
				Description: "A number that uniquely identifies the cluster node.",
			},
			"routemonitor": schema.StringAttribute{
				Required:    true,
				Description: "The IP address (IPv4 or IPv6).",
			},
		},
	}
}
