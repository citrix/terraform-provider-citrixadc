package netbridge_nsip6_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NetbridgeNsip6BindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "The subnet that is extended by this network bridge.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the network bridge.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The network mask for the subnet.",
			},
		},
	}
}
