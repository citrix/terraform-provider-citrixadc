package lbsipparameters

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbsipparametersDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addrportvip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add the rport parameter to the VIA headers of SIP requests that virtual servers receive from clients or servers.",
			},
			"retrydur": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which a client must wait before initiating a connection after receiving a 503 Service Unavailable response from the SIP server. The time value is sent in the \"Retry-After\" header in the 503 response.",
			},
			"rnatdstport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the destination port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"rnatsecuredstport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the destination port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"rnatsecuresrcport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the source port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"rnatsrcport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the source port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"sip503ratethreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of 503 Service Unavailable responses to generate, once every 10 milliseconds, when a SIP virtual server becomes unavailable.",
			},
		},
	}
}
