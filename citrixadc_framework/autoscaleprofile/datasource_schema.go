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
			"apikey_wo": schema.StringAttribute{
				Optional:    true,
				Description: "api key for authentication with service",
			},
			"apikey_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a apikey_wo update.",
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
			"sharedsecret_wo": schema.StringAttribute{
				Optional:    true,
				Description: "shared secret for authentication with service",
			},
			"sharedsecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a sharedsecret_wo update.",
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
