package vpnvserver_vpnsecureprivateaccessprofile_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverVpnsecureprivateaccessprofileBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"secureprivateaccessprofile": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Secure Private Access profile bound to the vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
		},
	}
}
