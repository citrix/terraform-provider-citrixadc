package lsnhttphdrlogprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnhttphdrlogprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"httphdrlogprofilename": schema.StringAttribute{
				Required:    true,
				Description: "The name of the HTTP header logging Profile.",
			},
			"loghost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host information is logged if option is enabled.",
			},
			"logmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP method information is logged if option is enabled.",
			},
			"logurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL information is logged if option is enabled.",
			},
			"logversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Version information is logged if option is enabled.",
			},
		},
	}
}
