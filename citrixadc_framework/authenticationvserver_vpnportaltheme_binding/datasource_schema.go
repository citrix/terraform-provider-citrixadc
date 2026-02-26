package authenticationvserver_vpnportaltheme_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationvserverVpnportalthemeBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authentication virtual server to which to bind the policy.",
			},
			"portaltheme": schema.StringAttribute{
				Required:    true,
				Description: "Theme for Authentication virtual server Login portal",
			},
		},
	}
}
