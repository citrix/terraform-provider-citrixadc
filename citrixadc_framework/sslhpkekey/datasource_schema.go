package sslhpkekey

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslhpkekeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dhkem": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of curve used for HPKE",
			},
			"file": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HPKE key file",
			},
			"hpkekeyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the hpke key configured that is used to decrypt ECH",
			},
		},
	}
}
