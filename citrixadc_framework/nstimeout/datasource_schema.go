package nstimeout

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NstimeoutDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"anyclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for non-TCP client connections. This value is over ridden by the client timeout that is configured on individual entities.",
			},
			"anyserver": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for non TCP server connections. This value is over ridden by the server timeout that is configured on individual entities.",
			},
			"anytcpclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for TCP client connections. This value takes precedence over  entity level timeout settings (vserver/service). This is applicable only to transport protocol TCP.",
			},
			"anytcpserver": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for TCP server connections. This value takes precedence over entity level timeout settings ( vserver/service). This is applicable only to transport protocol TCP.",
			},
			"client": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Client idle timeout (in seconds). If zero, the service-type default value is taken when service is created.",
			},
			"halfclose": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle timeout, in seconds, for connections that are in TCP half-closed state.",
			},
			"httpclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for client connections of HTTP service type. This value is over ridden by the client timeout that is configured on individual entities.",
			},
			"httpserver": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for server connections of HTTP service type. This value is over ridden by the server timeout that is configured on individual entities.",
			},
			"newconnidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timer interval, in seconds, for new TCP NATPCB connections on which no data was received.",
			},
			"nontcpzombie": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval at which the zombie clean-up process for non-TCP connections should run. Inactive IP NAT connections will be cleaned up.",
			},
			"reducedfintimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Alternative idle timeout, in seconds, for closed TCP NATPCB connections.",
			},
			"reducedrsttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timer interval, in seconds, for abruptly terminated TCP NATPCB connections.",
			},
			"server": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Server idle timeout (in seconds).  If zero, the service-type default value is taken when service is created.",
			},
			"tcpclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for non-HTTP client connections of TCP service type. This value is over ridden by the client timeout that is configured on individual entities.",
			},
			"tcpserver": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for non-HTTP server connections of TCP service type. This value is over ridden by the server timeout that is configured on entities.",
			},
			"zombie": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in seconds, at which the Citrix ADC zombie cleanup process must run. This process cleans up inactive TCP connections.",
			},
		},
	}
}
