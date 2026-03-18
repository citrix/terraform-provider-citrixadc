package ntpparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NtpparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Apply NTP authentication, which enables the NTP client (Citrix ADC) to verify that the server is in fact known and trusted.",
			},
			"autokeylogsec": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Autokey protocol requires the keys to be refreshed periodically. This parameter specifies the interval between regenerations of new session keys. In seconds, expressed as a power of 2.",
			},
			"revokelogsec": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval between re-randomizations of the autokey seeds to prevent brute-force attacks on the autokey algorithms.",
			},
			"trustedkey": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Key identifiers that are trusted for server authentication with symmetric key cryptography in the keys file.",
			},
		},
	}
}
