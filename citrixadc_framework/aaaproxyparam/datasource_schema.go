package aaaproxyparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaaproxyparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"proxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address and Port of the proxy server to be used for HTTP access for this request. Configure in ipaddress:port format (a.b.c.d:e) or as a URL (http://a.b.c.d or http://a.b.c.d:8080).",
			},
			"proxyauthorization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This indicates whether Proxy-Authorization header will be sent or not. Possible values = disabled, basic",
			},
			"proxyusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username that will be sent as part of Basic Proxy-Authorization header.",
			},
			"proxypassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Password that will be sent as part of Basic Proxy-Authorization header.",
			},
			"proxypassword_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password that will be sent as part of Basic Proxy-Authorization header.",
			},
			"proxypassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a proxypassword_wo update.",
			},
		},
	}
}
