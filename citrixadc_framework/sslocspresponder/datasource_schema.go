package sslocspresponder

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslocspresponderDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"batchingdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time, in milliseconds, to wait to accumulate OCSP requests to batch.  Does not apply if the Batching Depth is 1.",
			},
			"batchingdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of client certificates to batch together into one OCSP request. Batching avoids overloading the OCSP responder. A value of 1 signifies that each request is queried independently. For a value greater than 1, specify a timeout (batching delay) to avoid inordinately delaying the processing of a single certificate.",
			},
			"cache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable caching of responses. Caching of responses received from the OCSP responder enables faster responses to the clients and reduces the load on the OCSP responder.",
			},
			"cachetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout for caching the OCSP response. After the timeout, the Citrix ADC sends a fresh request to the OCSP responder for the certificate status. If a timeout is not specified, the timeout provided in the OCSP response applies.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP method used to send ocsp request. POST is the default httpmethod. If request length is > 255, POST wil be used even if GET is set as httpMethod",
			},
			"insertclientcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the complete client certificate in the OCSP request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the OCSP responder. Cannot begin with a hash (#) or space character and must contain only ASCII alphanumeric, underscore (_), hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the responder is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder\" or 'my responder').",
			},
			"ocspurlresolvetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, to wait for an OCSP URL Resolution. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server.",
			},
			"producedattimeskew": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which the Citrix ADC waits before considering the response as invalid. The response is considered invalid if the Produced At time stamp in the OCSP response exceeds or precedes the current Citrix ADC clock time by the amount of time specified.",
			},
			"respondercert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"resptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, to wait for an OCSP response. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server. Includes Batching Delay time.",
			},
			"signingcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Certificate-key pair that is used to sign OCSP requests. If this parameter is not set, the requests are not signed.",
			},
			"trustresponder": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A certificate to use to validate OCSP responses.  Alternatively, if -trustResponder is specified, no verification will be done on the reponse.  If both are omitted, only the response times (producedAt, lastUpdate, nextUpdate) will be verified.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the OCSP responder.",
			},
			"usenonce": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the OCSP nonce extension, which is designed to prevent replay attacks.",
			},
		},
	}
}
