package hanode_routemonitor6_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func HanodeRoutemonitor6BindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"hanode_id": schema.Int64Attribute{
				Required:    true,
				Description: "Number that uniquely identifies the local node. The ID of the local node is always 0.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask.",
			},
			"routemonitor": schema.StringAttribute{
				Required:    true,
				Description: "The IP address (IPv4 or IPv6).",
			},
		},
	}
}
