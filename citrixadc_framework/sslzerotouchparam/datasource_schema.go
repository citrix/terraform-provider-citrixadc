package sslzerotouchparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslzerotouchparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ocspcachetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout(in minutes) for caching the OCSP response.",
			},
			"ocspbatchingdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of certificates to batch together into one OCSP request. Batching avoids overloading the OCSP responder. A value of 1 signifies that each request is queried independently. For a value greater than 1, specify a timeout (batching delay) to avoid inordinately delaying the processing of a single certificate.",
			},
			"ocspbatchingdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time, in milliseconds, to wait to accumulate OCSP requests to batch. Does not apply if the Batching Depth is 1.",
			},
			"ocspresptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, to wait for an OCSP response. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server. Includes Batching Delay time.",
			},
			"ocspurlresolvetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, to wait for an OCSP URL Resolution. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server.",
			},
			"ocsptrustresponder": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If trustResponder is set to YES, signature verification will be skipped for the OCSP response. Possible values = YES, NO",
			},
			"ocspproducedattimeskew": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which the Citrix ADC waits before considering the response as invalid. The response is considered invalid if the Produced At time stamp in the OCSP response exceeds or precedes the current Citrix ADC clock time by the amount of time specified.",
			},
			"ocspusenonce": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the OCSP nonce extension, which is designed to prevent replay attacks. Possible values = ENABLED, DISABLED",
			},
			"ocsphttpmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP method used to send ocsp request. POST is the default httpmethod. If request length is > 255, POST wil be used even if GET is set as httpMethod. Possible values = GET, POST",
			},
		},
	}
}
