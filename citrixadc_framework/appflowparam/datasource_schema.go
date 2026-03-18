package appflowparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppflowparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aaausername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow AAA Username logging.",
			},
			"analyticsauthtoken": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication token to be set by the agent.",
			},
			"appnamerefresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in seconds, at which to send Appnames to the configured collectors. Appname refers to the name of an entity (virtual server, service, or service group) in the Citrix ADC.",
			},
			"auditlogs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable Auditlogs to be sent to the Telemetry Agent",
			},
			"cacheinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag to determine whether cache records need to be exported or not. If this flag is true and IC is enabled, cache records are exported instead of L7 HTTP records",
			},
			"clienttrafficonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Generate AppFlow records for only the traffic from the client.",
			},
			"connectionchaining": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable connection chaining so that the client server flows of a connection are linked. Also the connection chain ID is propagated across Citrix ADCs, so that in a multi-hop environment the flows belonging to the same logical connection are linked. This id is also logged as part of appflow record",
			},
			"cqareporting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP CQA reporting enable/disable knob.",
			},
			"distributedtracing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable generation of the distributed tracing templates in the Appflow records",
			},
			"disttracingsamplingrate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Sampling rate for Distributed Tracing",
			},
			"emailaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow user email-id logging.",
			},
			"events": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable Events to be sent to the Telemetry Agent",
			},
			"flowrecordinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in seconds, at which to send flow records to the configured collectors.",
			},
			"gxsessionreporting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable this option for Gx session reporting",
			},
			"httpauthorization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the HTTP Authorization header information.",
			},
			"httpcontenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the HTTP Content-Type header sent from the server to the client to determine the type of the content sent.",
			},
			"httpcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the cookie that was in the HTTP request the appliance received from the client.",
			},
			"httpdomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the http domain request to be exported.",
			},
			"httphost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the host identified in the HTTP request that the appliance received from the client.",
			},
			"httplocation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the HTTP location headers returned from the HTTP responses.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the method that was specified in the HTTP request that the appliance received from the client.",
			},
			"httpquerywithurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the HTTP query segment along with the URL that the Citrix ADC received from the client.",
			},
			"httpreferer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the web page that was last visited by the client.",
			},
			"httpsetcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the Set-cookie header sent from the server to the client in response to a HTTP request.",
			},
			"httpsetcookie2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the Set-cookie header sent from the server to the client in response to a HTTP request.",
			},
			"httpurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the http URL that the Citrix ADC received from the client.",
			},
			"httpuseragent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the client application through which the HTTP request was received by the Citrix ADC.",
			},
			"httpvia": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the httpVia header which contains the IP address of proxy server through which the client accessed the server.",
			},
			"httpxforwardedfor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the httpXForwardedFor header, which contains the original IP Address of the client using a proxy server to access the server.",
			},
			"identifiername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the stream identifier name to be exported.",
			},
			"identifiersessionname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the stream identifier session name to be exported.",
			},
			"logstreamovernsip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To use the Citrix ADC IP to send Logstream records instead of the SNIP",
			},
			"lsnlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the Large Scale Nat(LSN) records to the configured collectors.",
			},
			"metrics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable Citrix ADC Stats to be sent to the Telemetry Agent",
			},
			"observationdomainid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An observation domain groups a set of Citrix ADCs based on deployment: cluster, HA etc. A unique Observation Domain ID is required to be assigned to each such group.",
			},
			"observationdomainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Observation Domain defined by the observation domain ID.",
			},
			"observationpointid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An observation point ID is identifier for the NetScaler from which appflow records are being exported. By default, the NetScaler IP is the observation point ID.",
			},
			"securityinsightrecordinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in seconds, at which to send security insight flow records to the configured collectors.",
			},
			"securityinsighttraffic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/disable the feature individually on appflow action.",
			},
			"skipcacheredirectionhttptransaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Skip Cache http transaction. This HTTP transaction is specific to Cache Redirection module. In Case of Cache Miss there will be another HTTP transaction initiated by the cache server.",
			},
			"subscriberawareness": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable this option for logging end user MSISDN in L4/L7 appflow records",
			},
			"subscriberidobfuscation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable this option for obfuscating MSISDN in L4/L7 appflow records",
			},
			"subscriberidobfuscationalgo": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Algorithm(MD5 or SHA256) to be used for obfuscating MSISDN",
			},
			"tcpattackcounterinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in seconds, at which to send tcp attack counters to the configured collectors. If 0 is configured, the record is not sent.",
			},
			"templaterefresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Refresh interval, in seconds, at which to export the template data. Because data transmission is in UDP, the templates must be resent at regular intervals.",
			},
			"timeseriesovernsip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To use the Citrix ADC IP to send Time series data such as metrics and events, instead of the SNIP",
			},
			"udppmtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MTU, in bytes, for IPFIX UDP packets.",
			},
			"urlcategory": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the URL category record.",
			},
			"usagerecordinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the NGS will send bandwidth usage record to configured collectors.",
			},
			"videoinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/disable the feature individually on appflow action.",
			},
			"websaasappusagereporting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, NGS will send data used by Web/saas app at the end of every HTTP transaction to configured collectors.",
			},
		},
	}
}
