package vpnalwaysonprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnalwaysonprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientcontrol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow/Deny user to log off and connect to another Gateway",
			},
			"locationbasedvpn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to decide if tunnel should be established when in enterprise network. When locationBasedVPN is remote, client tries to detect if it is located in enterprise network or not and establishes the tunnel if not in enterprise network. Dns suffixes configured using -add dns suffix- are used to decide if the client is in the enterprise network or not. If the resolution of the DNS suffix results in private IP, client is said to be in enterprise network. When set to EveryWhere, the client skips the check to detect if it is on the enterprise network and tries to establish the tunnel",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "name of AlwaysON profile",
			},
			"networkaccessonvpnfailure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to block network traffic when tunnel is not established(and the config requires that tunnel be established). When set to onlyToGateway, the network traffic to and from the client (except Gateway IP) is blocked. When set to fullAccess, the network traffic is not blocked",
			},
		},
	}
}
