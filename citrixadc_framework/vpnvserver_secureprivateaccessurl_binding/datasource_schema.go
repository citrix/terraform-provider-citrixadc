package vpnvserver_secureprivateaccessurl_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverSecureprivateaccessurlBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"secureprivateaccessurl": schema.StringAttribute{
				Required:    true,
				Description: "Configured Secure Private Access URL",
			},
		},
	}
}
