package vpnvserver_vpnurl_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverVpnurlBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"urlname": schema.StringAttribute{
				Required:    true,
				Description: "The intranet URL.",
			},
		},
	}
}
