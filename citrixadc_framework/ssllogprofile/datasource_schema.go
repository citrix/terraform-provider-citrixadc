package ssllogprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SsllogprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the ssllogprofile.",
			},
			"ssllogclauth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "log all SSL ClAuth events.",
			},
			"ssllogclauthfailures": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "log all SSL ClAuth error events.",
			},
			"sslloghs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "log all SSL HS events.",
			},
			"sslloghsfailures": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "log all SSL HS error events.",
			},
		},
	}
}
