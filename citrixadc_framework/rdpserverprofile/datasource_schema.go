package rdpserverprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RdpserverprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the rdp server profile",
			},
			"psk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pre shared key value",
			},
			"rdpip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address of RDP listener. This terminates client RDP connections.",
			},
			"rdpport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP port on which the RDP connection is established.",
			},
			"rdpredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable RDP redirection support. This needs to be enabled in presence of connection broker or session directory with IP cookie(msts cookie) based redirection support",
			},
		},
	}
}
