package quicparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func QuicparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"quicsecrettimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Rotation frequency, in seconds, for the secret used to generate address validation tokens that will be issued in QUIC Retry packets and QUIC NEW_TOKEN frames sent by the Citrix ADC. A value of 0 can be configured if secret rotation is not desired.",
			},
		},
	}
}