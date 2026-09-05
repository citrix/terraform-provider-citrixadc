package dnssvcbrec

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnssvcbrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Domain name for the SVCB/HTTPS record.",
			},
			"targetname": schema.StringAttribute{
				Required:    true,
				Description: "Target domain name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Service priority (0 for AliasMode, >0 for ServiceMode).",
			},
			"svcbtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service type: SVCB or HTTPS. Possible values = SVCB, HTTPS",
			},
			"alpn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of ALPN protocol identifiers.",
			},
			"encryptedclienthello": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Base64-encoded ECH configuration.",
			},
			"ipv4hint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of IPv4 hint addresses.",
			},
			"ipv6hint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of IPv6 hint addresses.",
			},
			"mandatory": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of mandatory SvcParam keys.",
			},
			"nodefaultalpn": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates no default ALPN protocols.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for the service.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL) in seconds.",
			},
		},
	}
}
