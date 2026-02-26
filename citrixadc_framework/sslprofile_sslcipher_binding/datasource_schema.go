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
			"cipheraliasname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the cipher group/alias/individual cipheri bindings.",
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
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The cipher suite description.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL profile.",
			},
		},
	}
}
