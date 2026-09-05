package sslcipher

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslcipherDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ciphergroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the user-defined cipher group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the cipher group is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ciphergroup\" or 'my ciphergroup').",
			},
			"ciphersuitebinding": schema.SetNestedAttribute{
				Computed:    true,
				Description: "The ciphersuites (ciphername + cipherpriority) bound to this cipher group.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ciphername": schema.StringAttribute{
							Computed:    true,
							Description: "Cipher name.",
						},
						"cipherpriority": schema.Int64Attribute{
							Computed:    true,
							Description: "This indicates priority assigned to the particular cipher",
						},
					},
				},
			},
		},
	}
}
