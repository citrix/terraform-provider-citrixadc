package appflowparam

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppflowparamDataSourceModel is the data-source-specific model, decoupled from
// AppflowparamResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type AppflowparamDataSourceModel struct {
	Id                                  types.String `tfsdk:"id"`
	Aaausername                         types.String `tfsdk:"aaausername"`
	Analyticsauthtoken                  types.String `tfsdk:"analyticsauthtoken"`
	Appnamerefresh                      types.Int64  `tfsdk:"appnamerefresh"`
	Auditlogs                           types.String `tfsdk:"auditlogs"`
	Cacheinsight                        types.String `tfsdk:"cacheinsight"`
	Clienttrafficonly                   types.String `tfsdk:"clienttrafficonly"`
	Connectionchaining                  types.String `tfsdk:"connectionchaining"`
	Cqareporting                        types.String `tfsdk:"cqareporting"`
	Distributedtracing                  types.String `tfsdk:"distributedtracing"`
	Disttracingsamplingrate             types.Int64  `tfsdk:"disttracingsamplingrate"`
	Emailaddress                        types.String `tfsdk:"emailaddress"`
	Events                              types.String `tfsdk:"events"`
	Flowrecordinterval                  types.Int64  `tfsdk:"flowrecordinterval"`
	Gxsessionreporting                  types.String `tfsdk:"gxsessionreporting"`
	Httpauthorization                   types.String `tfsdk:"httpauthorization"`
	Httpcontenttype                     types.String `tfsdk:"httpcontenttype"`
	Httpcookie                          types.String `tfsdk:"httpcookie"`
	Httpdomain                          types.String `tfsdk:"httpdomain"`
	Httphost                            types.String `tfsdk:"httphost"`
	Httplocation                        types.String `tfsdk:"httplocation"`
	Httpmethod                          types.String `tfsdk:"httpmethod"`
	Httpquerywithurl                    types.String `tfsdk:"httpquerywithurl"`
	Httpreferer                         types.String `tfsdk:"httpreferer"`
	Httpsetcookie                       types.String `tfsdk:"httpsetcookie"`
	Httpsetcookie2                      types.String `tfsdk:"httpsetcookie2"`
	Httpurl                             types.String `tfsdk:"httpurl"`
	Httpuseragent                       types.String `tfsdk:"httpuseragent"`
	Httpvia                             types.String `tfsdk:"httpvia"`
	Httpxforwardedfor                   types.String `tfsdk:"httpxforwardedfor"`
	Identifiername                      types.String `tfsdk:"identifiername"`
	Identifiersessionname               types.String `tfsdk:"identifiersessionname"`
	Logalljsonfields                    types.String `tfsdk:"logalljsonfields"`
	Logstreamovernsip                   types.String `tfsdk:"logstreamovernsip"`
	Lsnlogging                          types.String `tfsdk:"lsnlogging"`
	Metrics                             types.String `tfsdk:"metrics"`
	Observationdomainid                 types.Int64  `tfsdk:"observationdomainid"`
	Observationdomainname               types.String `tfsdk:"observationdomainname"`
	Observationpointid                  types.Int64  `tfsdk:"observationpointid"`
	Securityinsightrecordinterval       types.Int64  `tfsdk:"securityinsightrecordinterval"`
	Securityinsighttraffic              types.String `tfsdk:"securityinsighttraffic"`
	Skipcacheredirectionhttptransaction types.String `tfsdk:"skipcacheredirectionhttptransaction"`
	Subscriberawareness                 types.String `tfsdk:"subscriberawareness"`
	Subscriberidobfuscation             types.String `tfsdk:"subscriberidobfuscation"`
	Subscriberidobfuscationalgo         types.String `tfsdk:"subscriberidobfuscationalgo"`
	Tcpattackcounterinterval            types.Int64  `tfsdk:"tcpattackcounterinterval"`
	Templaterefresh                     types.Int64  `tfsdk:"templaterefresh"`
	Timeseriesovernsip                  types.String `tfsdk:"timeseriesovernsip"`
	Udppmtu                             types.Int64  `tfsdk:"udppmtu"`
	Urlcategory                         types.String `tfsdk:"urlcategory"`
	Usagerecordinterval                 types.Int64  `tfsdk:"usagerecordinterval"`
	Videoinsight                        types.String `tfsdk:"videoinsight"`
	Websaasappusagereporting            types.String `tfsdk:"websaasappusagereporting"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appflowparam.json). Never settable; populated from GET.
	Builtin                    types.List   `tfsdk:"builtin"`
	Feature                    types.String `tfsdk:"feature"`
	Tcpburstreporting          types.String `tfsdk:"tcpburstreporting"`
	Tcpburstreportingthreshold types.Int64  `tfsdk:"tcpburstreportingthreshold"`
}

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
				Sensitive:   true,
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
			"logalljsonfields": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overrides the field filtering for all analytics profiles, and sends all the fields for the configured insights.",
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

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if the appflow param is built-in or not. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ].",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"tcpburstreporting": schema.StringAttribute{
				Computed:    true,
				Description: "TCP burst reporting enable/disable knob. Possible values: [ ENABLED, DISABLED ].",
			},
			"tcpburstreportingthreshold": schema.Int64Attribute{
				Computed:    true,
				Description: "TCP burst reporting threshold.",
			},
		},
	}
}

// appflowparamDataSourceSetAttrFromGet projects a NITRO appflowparam GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func appflowparamDataSourceSetAttrFromGet(ctx context.Context, data *AppflowparamDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appflowparamDataSourceSetAttrFromGet Function")

	// appflowparam is a singleton config object with no lookup key.
	data.Id = types.StringValue("appflowparam-config")

	// Read/write attributes as read-back outputs.
	data.Aaausername = utils.MapGetString(g, "aaausername")
	data.Appnamerefresh = utils.MapGetInt64(g, "appnamerefresh")
	data.Auditlogs = utils.MapGetString(g, "auditlogs")
	data.Cacheinsight = utils.MapGetString(g, "cacheinsight")
	data.Clienttrafficonly = utils.MapGetString(g, "clienttrafficonly")
	data.Connectionchaining = utils.MapGetString(g, "connectionchaining")
	data.Cqareporting = utils.MapGetString(g, "cqareporting")
	data.Distributedtracing = utils.MapGetString(g, "distributedtracing")
	data.Disttracingsamplingrate = utils.MapGetInt64(g, "disttracingsamplingrate")
	data.Emailaddress = utils.MapGetString(g, "emailaddress")
	data.Events = utils.MapGetString(g, "events")
	data.Flowrecordinterval = utils.MapGetInt64(g, "flowrecordinterval")
	data.Gxsessionreporting = utils.MapGetString(g, "gxsessionreporting")
	data.Httpauthorization = utils.MapGetString(g, "httpauthorization")
	data.Httpcontenttype = utils.MapGetString(g, "httpcontenttype")
	data.Httpcookie = utils.MapGetString(g, "httpcookie")
	data.Httpdomain = utils.MapGetString(g, "httpdomain")
	data.Httphost = utils.MapGetString(g, "httphost")
	data.Httplocation = utils.MapGetString(g, "httplocation")
	data.Httpmethod = utils.MapGetString(g, "httpmethod")
	data.Httpquerywithurl = utils.MapGetString(g, "httpquerywithurl")
	data.Httpreferer = utils.MapGetString(g, "httpreferer")
	data.Httpsetcookie = utils.MapGetString(g, "httpsetcookie")
	data.Httpsetcookie2 = utils.MapGetString(g, "httpsetcookie2")
	data.Httpurl = utils.MapGetString(g, "httpurl")
	data.Httpuseragent = utils.MapGetString(g, "httpuseragent")
	data.Httpvia = utils.MapGetString(g, "httpvia")
	data.Httpxforwardedfor = utils.MapGetString(g, "httpxforwardedfor")
	data.Identifiername = utils.MapGetString(g, "identifiername")
	data.Identifiersessionname = utils.MapGetString(g, "identifiersessionname")
	data.Logalljsonfields = utils.MapGetString(g, "logalljsonfields")
	data.Logstreamovernsip = utils.MapGetString(g, "logstreamovernsip")
	data.Lsnlogging = utils.MapGetString(g, "lsnlogging")
	data.Metrics = utils.MapGetString(g, "metrics")
	data.Observationdomainid = utils.MapGetInt64(g, "observationdomainid")
	data.Observationdomainname = utils.MapGetString(g, "observationdomainname")
	data.Observationpointid = utils.MapGetInt64(g, "observationpointid")
	data.Securityinsightrecordinterval = utils.MapGetInt64(g, "securityinsightrecordinterval")
	data.Securityinsighttraffic = utils.MapGetString(g, "securityinsighttraffic")
	data.Skipcacheredirectionhttptransaction = utils.MapGetString(g, "skipcacheredirectionhttptransaction")
	data.Subscriberawareness = utils.MapGetString(g, "subscriberawareness")
	data.Subscriberidobfuscation = utils.MapGetString(g, "subscriberidobfuscation")
	data.Subscriberidobfuscationalgo = utils.MapGetString(g, "subscriberidobfuscationalgo")
	data.Tcpattackcounterinterval = utils.MapGetInt64(g, "tcpattackcounterinterval")
	data.Templaterefresh = utils.MapGetInt64(g, "templaterefresh")
	data.Timeseriesovernsip = utils.MapGetString(g, "timeseriesovernsip")
	data.Udppmtu = utils.MapGetInt64(g, "udppmtu")
	data.Urlcategory = utils.MapGetString(g, "urlcategory")
	data.Usagerecordinterval = utils.MapGetInt64(g, "usagerecordinterval")
	data.Videoinsight = utils.MapGetString(g, "videoinsight")
	data.Websaasappusagereporting = utils.MapGetString(g, "websaasappusagereporting")

	// analyticsauthtoken is a secret input the GET never returns -> Null.
	data.Analyticsauthtoken = types.StringNull()

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Tcpburstreporting = utils.MapGetString(g, "tcpburstreporting")
	data.Tcpburstreportingthreshold = utils.MapGetInt64(g, "tcpburstreportingthreshold")
}
