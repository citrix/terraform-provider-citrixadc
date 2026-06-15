package vpnvserver_intranetip_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverIntranetipBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"intranetip": schema.StringAttribute{
				Required:    true,
				Description: "The network ID for the range of intranet IP addresses or individual intranet IP addresses to be bound to the virtual server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask of the intranet IP address or range.",
			},
		},
	}
}
