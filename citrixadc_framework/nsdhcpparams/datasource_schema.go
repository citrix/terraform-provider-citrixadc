package nsdhcpparams

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsdhcpparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dhcpclient": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables DHCP client to acquire IP address from the DHCP server in the next boot. When set to OFF, disables the DHCP client in the next boot.",
			},
			"saveroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DHCP acquired routes are saved on the Citrix ADC.",
			},
		},
	}
}
