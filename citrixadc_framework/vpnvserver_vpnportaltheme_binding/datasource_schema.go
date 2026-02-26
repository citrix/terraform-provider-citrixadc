package vpnvserver_vpnportaltheme_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverVpnportalthemeBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"portaltheme": schema.StringAttribute{
				Required:    true,
				Description: "Name of the portal theme bound to VPN vserver",
			},
		},
	}
}
