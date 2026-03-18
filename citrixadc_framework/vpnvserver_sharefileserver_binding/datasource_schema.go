package vpnvserver_sharefileserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverSharefileserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"sharefile": schema.StringAttribute{
				Required:    true,
				Description: "Configured ShareFile server in XenMobile deployment. Format IP:PORT / FQDN:PORT",
			},
		},
	}
}
