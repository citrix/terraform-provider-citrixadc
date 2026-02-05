package vpnintranetapplication

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func VpnintranetapplicationDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientapplication": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Names of the client applications, such as PuTTY and Xshell.",
			},
			"destip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination IP address, IP range, or host name of the intranet application. This address is the server IP address.",
			},
			"destport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination TCP or UDP port number for the intranet application. Use a hyphen to specify a range of port numbers, for example 90-95.",
			},
			"hostname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the host for which to configure interception. The names are resolved during interception when users log on with the Citrix Gateway Plug-in.",
			},
			"interception": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interception mode for the intranet application or resource. Correct value depends on the type of client software used to make connections. If the interception mode is set to TRANSPARENT, users connect with the Citrix Gateway Plug-in for Windows. With the PROXY setting, users connect with the Citrix Gateway Plug-in for Java.",
			},
			"intranetapplication": schema.StringAttribute{
				Required:    true,
				Description: "Name of the intranet application.",
			},
			"iprange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If you have multiple servers in your network, such as web, email, and file shares, configure an intranet application that includes the IP range for all the network applications. This allows users to access all the intranet applications contained in the IP address range.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination subnet mask for the intranet application.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by the intranet application. If protocol is set to BOTH, TCP and UDP traffic is allowed.",
			},
			"spoofiip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address that the intranet application will use to route the connection through the virtual adapter.",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IP address. Required if interception mode is set to PROXY. Default is the loopback address, 127.0.0.1.",
			},
			"srcport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Source port for the application for which the Citrix Gateway virtual server proxies the traffic. If users are connecting from a device that uses the Citrix Gateway Plug-in for Java, applications must be configured manually by using the source IP address and TCP port values specified in the intranet application profile. If a port value is not set, the destination port value is used.",
			},
		},
	}
}
