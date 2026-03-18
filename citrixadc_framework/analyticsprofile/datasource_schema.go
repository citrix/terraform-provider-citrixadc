package analyticsprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AnalyticsprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allhttpheaders": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log all the request and response headers.",
			},
			"analyticsauthtoken": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorizaiton header is required to be of the form - Splunk <auth-token>.",
			},
			"analyticsendpointcontenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By default, application/json content-type is used. If this needs to be overridden, specify the value.",
			},
			"analyticsendpointmetadata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the endpoint requires some metadata to be present before the actual json data, specify the same.",
			},
			"analyticsendpointurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The URL at which to upload the analytics data on the endpoint",
			},
			"auditlogs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option indicates the whether auditlog should be sent to the REST collector.",
			},
			"collectors": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The collector can be an IP, an appflow collector name, a service or a vserver. If IP is specified, the transport is considered as logstream and default port of 5557 is taken. If collector name is specified, the collector properties are taken from the configured collector. If service is specified, the configured service is assumed as the collector. If vserver is specified, the services bound to it are considered as collectors and the records are load balanced.",
			},
			"cqareporting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log TCP CQA parameters.",
			},
			"dataformatfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is for configuring the file containing the data format and metadata required by the analytics endpoint.",
			},
			"events": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option indicates the whether events should be sent to the REST collector.",
			},
			"grpcstatus": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log the gRPC status headers",
			},
			"httpauthentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log Authentication header.",
			},
			"httpclientsidemeasurements": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will insert a javascript into the HTTP response to collect the client side page-timings and will send the same to the configured collectors.",
			},
			"httpcontenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log content-length header.",
			},
			"httpcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log cookie header.",
			},
			"httpcustomheaders": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Specify the list of custom headers to be exported in web transaction records.",
			},
			"httpdomainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log domain name.",
			},
			"httphost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log the Host header in appflow records",
			},
			"httplocation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log location header.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log the method header in appflow records",
			},
			"httppagetracking": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will link the embedded objects of a page together.",
			},
			"httpreferer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log the referer header in appflow records",
			},
			"httpsetcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log set-cookie header.",
			},
			"httpsetcookie2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log set-cookie2 header.",
			},
			"httpurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log the URL in appflow records",
			},
			"httpurlquery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log URL Query.",
			},
			"httpuseragent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log User-Agent header.",
			},
			"httpvia": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will Via header.",
			},
			"httpxforwardedforheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log X-Forwarded-For header.",
			},
			"integratedcache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log the Integrated Caching appflow records",
			},
			"managementlog": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "This option indicates the whether managementlog should be sent to the REST collector.",
			},
			"metrics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option indicates the whether metrics should be sent to the REST collector.",
			},
			"metricsexportfrequency": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is for configuring the metrics export frequency in seconds, frequency value must be in [30,300] seconds range",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the analytics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow profile\" or 'my appflow profile').",
			},
			"outputmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option indicates the format of REST API POST body. It depends on the consumer of the analytics data.",
			},
			"schemafile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is for configuring json schema file containing a list of counters to be exported by metricscollector",
			},
			"servemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is for setting the mode of how data is provided",
			},
			"tcpburstreporting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will log TCP burst parameters.",
			},
			"topn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this topn support, the topn information of the stream identifier this profile is bound to will be exported to the analytics endpoint.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option indicates what information needs to be collected and exported.",
			},
			"urlcategory": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the URL category record.",
			},
		},
	}
}
