package quicprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func QuicprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ackdelayexponent": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, indicating an exponent that the remote QUIC endpoint should use, to decode the ACK Delay field in QUIC ACK frames sent by the Citrix ADC.",
			},
			"activeconnectionidlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum number of QUIC connection IDs from the remote QUIC endpoint, that the Citrix ADC is willing to store.",
			},
			"activeconnectionmigration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify whether the Citrix ADC should allow the remote QUIC endpoint to perform active QUIC connection migration.",
			},
			"congestionctrlalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the congestion control algorithm to be used for QUIC connections. The default congestion control algorithm is CUBIC.",
			},
			"initialmaxdata": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial value, in bytes, for the maximum amount of data that can be sent on a QUIC connection.",
			},
			"initialmaxstreamdatabidilocal": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the Citrix ADC.",
			},
			"initialmaxstreamdatabidiremote": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the remote QUIC endpoint.",
			},
			"initialmaxstreamdatauni": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for unidirectional streams initiated by the remote QUIC endpoint.",
			},
			"initialmaxstreamsbidi": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial maximum number of bidirectional streams the remote QUIC endpoint may initiate.",
			},
			"initialmaxstreamsuni": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial maximum number of unidirectional streams the remote QUIC endpoint may initiate.",
			},
			"maxackdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum amount of time, in milliseconds, by which the Citrix ADC will delay sending acknowledgments.",
			},
			"maxidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum idle timeout, in seconds, for a QUIC connection. A QUIC connection will be silently discarded by the Citrix ADC if it remains idle for longer than the minimum of the idle timeout values advertised by the Citrix ADC and the remote QUIC endpoint, and three times the current Probe Timeout (PTO).",
			},
			"maxudpdatagramsperburst": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value, specifying the maximum number of UDP datagrams that can be transmitted by the Citrix ADC in a single transmission burst on a QUIC connection.",
			},
			"maxudppayloadsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the size of the largest UDP datagram payload, in bytes, that the Citrix ADC is willing to receive on a QUIC connection.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"newtokenvalidityperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value, specifying the validity period, in seconds, of address validation tokens issued through QUIC NEW_TOKEN frames sent by the Citrix ADC.",
			},
			"retrytokenvalidityperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value, specifying the validity period, in seconds, of address validation tokens issued through QUIC Retry packets sent by the Citrix ADC.",
			},
			"statelessaddressvalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify whether the Citrix ADC should perform stateless address validation for QUIC clients, by sending tokens in QUIC Retry packets during QUIC connection establishment, and by sending tokens in QUIC NEW_TOKEN frames after QUIC connection establishment.",
			},
		},
	}
}