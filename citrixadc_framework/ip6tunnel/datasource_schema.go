package ip6tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Ip6tunnelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"local": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An IPv6 address of the local Citrix ADC used to set up the tunnel.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the IPv6 Tunnel. Cannot be changed after the service group is created. Must begin with a number or letter, and can consist of letters, numbers, and the @ _ - . (period) : (colon) # and space ( ) characters.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for the tunnel.",
			},
			"remote": schema.StringAttribute{
				Required:    true,
				Description: "An IPv6 address of the remote Citrix ADC used to set up the tunnel.",
			},
		},
	}
}
