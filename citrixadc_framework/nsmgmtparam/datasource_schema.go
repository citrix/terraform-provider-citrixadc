package nsmgmtparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsmgmtparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"httpdmaxclients": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This enables setting the HTTPD Max Clients value in the httpd.conf file. You can configure either Max Clients or Max Request Workers. The allowable range is from a minimum of 1 to a maximum of 255",
			},
			"httpdmaxreqworkers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This enables setting the HTTPD Max Request Workers value in the httpd.conf file. You can configure either Max Clients or Max Request Workers. The allowable range is from a minimum of 1 to a maximum of 255",
			},
			"mgmthttpport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This allow the configuration of management HTTP port.",
			},
			"mgmthttpsport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This allows the configuration of management HTTPS port.",
			},
		},
	}
}
