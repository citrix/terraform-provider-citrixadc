package vpnvserver_staserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverStaserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"staaddresstype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the STA server address(ipv4/v6).",
			},
			"staserver": schema.StringAttribute{
				Required:    true,
				Description: "Configured Secure Ticketing Authority (STA) server.",
			},
		},
	}
}
