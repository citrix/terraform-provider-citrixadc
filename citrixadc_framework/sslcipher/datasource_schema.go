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
			"ciphername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cipher name.",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This indicates priority assigned to the particular cipher",
			},
			"ciphgrpalias": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The individual cipher name(s), a user-defined cipher group, or a system predefined cipher alias that will be added to the  predefined cipher alias that will be added to the group cipherGroupName.\nIf a cipher alias or a cipher group is specified, all the individual ciphers in the cipher alias or group will be added to the user-defined cipher group.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the profile to which cipher is attached.",
			},
		},
	}
}
