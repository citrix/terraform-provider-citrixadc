package callhome

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CallhomeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"emailaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Email address of the contact administrator.",
			},
			"hbcustominterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval (in days) between CallHome heartbeats",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the proxy server.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CallHome mode of operation",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP port on the Proxy server. This is a mandatory parameter for both IP address and service name based configuration.",
			},
			"proxyauthservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service that represents the proxy server.",
			},
			"proxymode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables or disables the proxy mode. The proxy server can be set by either specifying the IP address of the server or the name of the service representing the proxy server.",
			},
		},
	}
}
