package reputationsettings

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ReputationsettingsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"proxypassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password with which user logs on.",
			},
			"proxypassword_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Password with which user logs on.",
			},
			"proxypassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a proxypassword_wo update.",
			},
			"proxyport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy server port.",
			},
			"proxyserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy server IP to get Reputation data.",
			},
			"proxyusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Username",
			},
		},
	}
}
