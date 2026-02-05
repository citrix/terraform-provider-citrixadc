package sslfipskey

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslfipskeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"curve": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Only p_256 (prime256v1) and P_384 (secp384r1) are supported.",
			},
			"exponent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exponent value for the FIPS key to be created. Available values function as follows:\n 3=3 (hexadecimal)\nF4=10001 (hexadecimal)",
			},
			"fipskeyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the FIPS key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the FIPS key is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my fipskey\" or 'my fipskey').",
			},
			"inform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Input format of the key file. Available formats are:\nSIM - Secure Information Management; select when importing a FIPS key. If the external FIPS key is encrypted, first decrypt it, and then import it.\nPEM - Privacy Enhanced Mail; select when importing a non-FIPS key.",
			},
			"iv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initialization Vector (IV) to use for importing the key. Required for importing a non-FIPS key.",
			},
			"key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the key file to be imported.\n /nsconfig/ssl/ is the default path.",
			},
			"keytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Only RSA key and ECDSA Key are supported.",
			},
			"modulus": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Modulus, in multiples of 64, of the FIPS key to be created.",
			},
			"wrapkeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the wrap key to use for importing the key. Required for importing a non-FIPS key.",
			},
		},
	}
}
