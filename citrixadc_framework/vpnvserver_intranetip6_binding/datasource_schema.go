package vpnvserver_intranetip6_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverIntranetip6BindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"intranetip6": schema.StringAttribute{
				Required:    true,
				Description: "The network id for the range of intranet IP6 addresses or individual intranet ip to be bound to the vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"numaddr": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of ipv6 addresses",
			},
		},
	}
}
