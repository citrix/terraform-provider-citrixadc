package ssldtlsprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SsldtlsprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"helloverifyrequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send a Hello Verify request to validate the client.",
			},
			"initialretrytimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial time out value to retransmit the last flight sent from the NetScaler.",
			},
			"maxbadmacignorecount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of bad MAC errors to ignore for a connection prior disconnect. Disabling parameter terminateSession terminates session immediately when bad MAC is detected in the connection.",
			},
			"maxholdqlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of datagrams that can be queued at DTLS layer for processing",
			},
			"maxpacketsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of packets to reassemble. This value helps protect against a fragmented packet attack.",
			},
			"maxrecordsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of records that can be sent if PMTU is disabled.",
			},
			"maxretrytime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Wait for the specified time, in seconds, before resending the request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DTLS profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"pmtudiscovery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source for the maximum record size value. If ENABLED, the value is taken from the PMTU table. If DISABLED, the value is taken from the profile.",
			},
			"terminatesession": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Terminate the session if the message authentication code (MAC) of the client and server do not match.",
			},
		},
	}
}
