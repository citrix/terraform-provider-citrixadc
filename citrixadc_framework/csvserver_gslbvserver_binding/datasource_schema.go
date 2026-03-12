package csvserver_gslbvserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CsvserverGslbvserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"vserver": schema.StringAttribute{
				Required:    true,
				Description: "Name of the default gslb or vpn vserver bound to CS vserver of type GSLB/VPN. For Example: bind cs vserver cs1 -vserver gslb1 or bind cs vserver cs1 -vserver vpn1",
			},
		},
	}
}
