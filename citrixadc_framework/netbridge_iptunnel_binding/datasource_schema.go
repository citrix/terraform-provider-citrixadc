package netbridge_iptunnel_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NetbridgeIptunnelBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the network bridge.",
			},
			"tunnel": schema.StringAttribute{
				Required:    true,
				Description: "The name of the tunnel that is a part of this bridge.",
			},
		},
	}
}
