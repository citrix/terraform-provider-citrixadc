package sslprofile_sslcipher_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslprofileSslcipherBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ciphername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the cipher.",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "cipher priority",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL profile.",
			},
		},
	}
}
