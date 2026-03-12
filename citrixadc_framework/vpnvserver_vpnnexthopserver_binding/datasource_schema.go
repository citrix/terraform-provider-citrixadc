package vpnvserver_vpnnexthopserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverVpnnexthopserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"nexthopserver": schema.StringAttribute{
				Required:    true,
				Description: "The name of the next hop server bound to the VPN virtual server.",
			},
		},
	}
}
