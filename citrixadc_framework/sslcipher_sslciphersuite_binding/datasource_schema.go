package sslcipher_sslciphersuite_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslcipherSslciphersuiteBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ciphergroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the user-defined cipher group.",
			},
			"ciphername": schema.StringAttribute{
				Required:    true,
				Description: "Cipher name.",
			},
			"cipheroperation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The operation that is performed when adding the cipher-suite.\n\nPossible cipher operations are:\n	ADD - Appends the given cipher-suite to the existing one configured for the virtual server.\n	REM - Removes the given cipher-suite from the existing one configured for the virtual server.\n	ORD - Overrides the current configured cipher-suite for the virtual server with the given cipher-suite.",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This indicates priority assigned to the particular cipher",
			},
			"ciphgrpals": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or user defined cipher-group name.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cipher suite description.",
			},
		},
	}
}
