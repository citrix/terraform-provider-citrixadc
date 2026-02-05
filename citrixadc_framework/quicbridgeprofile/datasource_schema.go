package quicbridgeprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func QuicbridgeprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"routingalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Routing algorithm to generate routable connection IDs.",
			},
			"serveridlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Length of serverid to encode/decode server information",
			},
		},
	}
}
