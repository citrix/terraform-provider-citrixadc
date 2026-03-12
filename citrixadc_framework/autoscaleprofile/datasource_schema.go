package autoscaleprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AutoscaleprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"apikey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "api key for authentication with service",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "AutoScale profile name.",
			},
			"sharedsecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "shared secret for authentication with service",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of profile.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL providing the service",
			},
		},
	}
}
