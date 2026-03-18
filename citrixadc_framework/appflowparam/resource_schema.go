package appflowparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppflowparamResourceModel describes the resource data model.
type AppflowparamResourceModel struct {
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
}

func (r *AppflowparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appflowparam resource.",
			},
			"aaausername": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable AppFlow AAA Username logging.",
			},
			"analyticsauthtoken": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication token to be set by the agent.",
			},
			"appnamerefresh": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(600),
				Description: "Interval, in seconds, at which to send Appnames to the configured collectors. Appname refers to the name of an entity (virtual server, service, or service group) in the Citrix ADC.",
			},
			"auditlogs": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable Auditlogs to be sent to the Telemetry Agent",
			},
			"cacheinsight": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Flag to determine whether cache records need to be exported or not. If this flag is true and IC is enabled, cache records are exported instead of L7 HTTP records",
			},
			"clienttrafficonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Generate AppFlow records for only the traffic from the client.",
			},
			"connectionchaining": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable connection chaining so that the client server flows of a connection are linked. Also the connection chain ID is propagated across Citrix ADCs, so that in a multi-hop environment the flows belonging to the same logical connection are linked. This id is also logged as part of appflow record",
			},
			"cqareporting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "TCP CQA reporting enable/disable knob.",
			},
			"distributedtracing": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable generation of the distributed tracing templates in the Appflow records",
			},
			"disttracingsamplingrate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Sampling rate for Distributed Tracing",
			},
			"emailaddress": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable AppFlow user email-id logging.",
			},
			"events": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable Events to be sent to the Telemetry Agent",
			},
			"flowrecordinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(60),
				Description: "Interval, in seconds, at which to send flow records to the configured collectors.",
			},
			"gxsessionreporting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable this option for Gx session reporting",
			},
			"httpauthorization": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the HTTP Authorization header information.",
			},
			"httpcontenttype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the HTTP Content-Type header sent from the server to the client to determine the type of the content sent.",
			},
			"httpcookie": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the cookie that was in the HTTP request the appliance received from the client.",
			},
			"httpdomain": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the http domain request to be exported.",
			},
			"httphost": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the host identified in the HTTP request that the appliance received from the client.",
			},
			"httplocation": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the HTTP location headers returned from the HTTP responses.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the method that was specified in the HTTP request that the appliance received from the client.",
			},
			"httpquerywithurl": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the HTTP query segment along with the URL that the Citrix ADC received from the client.",
			},
			"httpreferer": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the web page that was last visited by the client.",
			},
			"httpsetcookie": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the Set-cookie header sent from the server to the client in response to a HTTP request.",
			},
			"httpsetcookie2": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the Set-cookie header sent from the server to the client in response to a HTTP request.",
			},
			"httpurl": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the http URL that the Citrix ADC received from the client.",
			},
			"httpuseragent": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the client application through which the HTTP request was received by the Citrix ADC.",
			},
			"httpvia": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the httpVia header which contains the IP address of proxy server through which the client accessed the server.",
			},
			"httpxforwardedfor": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the httpXForwardedFor header, which contains the original IP Address of the client using a proxy server to access the server.",
			},
			"identifiername": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the stream identifier name to be exported.",
			},
			"identifiersessionname": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the stream identifier session name to be exported.",
			},
			"logstreamovernsip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "To use the Citrix ADC IP to send Logstream records instead of the SNIP",
			},
			"lsnlogging": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will send the Large Scale Nat(LSN) records to the configured collectors.",
			},
			"metrics": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     int64default.StaticInt64(600),
				Description: "Interval, in seconds, at which to send security insight flow records to the configured collectors.",
			},
			"securityinsighttraffic": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/disable the feature individually on appflow action.",
			},
			"skipcacheredirectionhttptransaction": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Skip Cache http transaction. This HTTP transaction is specific to Cache Redirection module. In Case of Cache Miss there will be another HTTP transaction initiated by the cache server.",
			},
			"subscriberawareness": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable this option for logging end user MSISDN in L4/L7 appflow records",
			},
			"subscriberidobfuscation": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable this option for obfuscating MSISDN in L4/L7 appflow records",
			},
			"subscriberidobfuscationalgo": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("MD5"),
				Description: "Algorithm(MD5 or SHA256) to be used for obfuscating MSISDN",
			},
			"tcpattackcounterinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in seconds, at which to send tcp attack counters to the configured collectors. If 0 is configured, the record is not sent.",
			},
			"templaterefresh": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(600),
				Description: "Refresh interval, in seconds, at which to export the template data. Because data transmission is in UDP, the templates must be resent at regular intervals.",
			},
			"timeseriesovernsip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "To use the Citrix ADC IP to send Time series data such as metrics and events, instead of the SNIP",
			},
			"udppmtu": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1472),
				Description: "MTU, in bytes, for IPFIX UDP packets.",
			},
			"urlcategory": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Include the URL category record.",
			},
			"usagerecordinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the NGS will send bandwidth usage record to configured collectors.",
			},
			"videoinsight": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/disable the feature individually on appflow action.",
			},
			"websaasappusagereporting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, NGS will send data used by Web/saas app at the end of every HTTP transaction to configured collectors.",
			},
		},
	}
}

func appflowparamGetThePayloadFromtheConfig(ctx context.Context, data *AppflowparamResourceModel) appflow.Appflowparam {
	tflog.Debug(ctx, "In appflowparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appflowparam := appflow.Appflowparam{}
	if !data.Aaausername.IsNull() {
		appflowparam.Aaausername = data.Aaausername.ValueString()
	}
	if !data.Analyticsauthtoken.IsNull() {
		appflowparam.Analyticsauthtoken = data.Analyticsauthtoken.ValueString()
	}
	if !data.Appnamerefresh.IsNull() {
		appflowparam.Appnamerefresh = utils.IntPtr(int(data.Appnamerefresh.ValueInt64()))
	}
	if !data.Auditlogs.IsNull() {
		appflowparam.Auditlogs = data.Auditlogs.ValueString()
	}
	if !data.Cacheinsight.IsNull() {
		appflowparam.Cacheinsight = data.Cacheinsight.ValueString()
	}
	if !data.Clienttrafficonly.IsNull() {
		appflowparam.Clienttrafficonly = data.Clienttrafficonly.ValueString()
	}
	if !data.Connectionchaining.IsNull() {
		appflowparam.Connectionchaining = data.Connectionchaining.ValueString()
	}
	if !data.Cqareporting.IsNull() {
		appflowparam.Cqareporting = data.Cqareporting.ValueString()
	}
	if !data.Distributedtracing.IsNull() {
		appflowparam.Distributedtracing = data.Distributedtracing.ValueString()
	}
	if !data.Disttracingsamplingrate.IsNull() {
		appflowparam.Disttracingsamplingrate = utils.IntPtr(int(data.Disttracingsamplingrate.ValueInt64()))
	}
	if !data.Emailaddress.IsNull() {
		appflowparam.Emailaddress = data.Emailaddress.ValueString()
	}
	if !data.Events.IsNull() {
		appflowparam.Events = data.Events.ValueString()
	}
	if !data.Flowrecordinterval.IsNull() {
		appflowparam.Flowrecordinterval = utils.IntPtr(int(data.Flowrecordinterval.ValueInt64()))
	}
	if !data.Gxsessionreporting.IsNull() {
		appflowparam.Gxsessionreporting = data.Gxsessionreporting.ValueString()
	}
	if !data.Httpauthorization.IsNull() {
		appflowparam.Httpauthorization = data.Httpauthorization.ValueString()
	}
	if !data.Httpcontenttype.IsNull() {
		appflowparam.Httpcontenttype = data.Httpcontenttype.ValueString()
	}
	if !data.Httpcookie.IsNull() {
		appflowparam.Httpcookie = data.Httpcookie.ValueString()
	}
	if !data.Httpdomain.IsNull() {
		appflowparam.Httpdomain = data.Httpdomain.ValueString()
	}
	if !data.Httphost.IsNull() {
		appflowparam.Httphost = data.Httphost.ValueString()
	}
	if !data.Httplocation.IsNull() {
		appflowparam.Httplocation = data.Httplocation.ValueString()
	}
	if !data.Httpmethod.IsNull() {
		appflowparam.Httpmethod = data.Httpmethod.ValueString()
	}
	if !data.Httpquerywithurl.IsNull() {
		appflowparam.Httpquerywithurl = data.Httpquerywithurl.ValueString()
	}
	if !data.Httpreferer.IsNull() {
		appflowparam.Httpreferer = data.Httpreferer.ValueString()
	}
	if !data.Httpsetcookie.IsNull() {
		appflowparam.Httpsetcookie = data.Httpsetcookie.ValueString()
	}
	if !data.Httpsetcookie2.IsNull() {
		appflowparam.Httpsetcookie2 = data.Httpsetcookie2.ValueString()
	}
	if !data.Httpurl.IsNull() {
		appflowparam.Httpurl = data.Httpurl.ValueString()
	}
	if !data.Httpuseragent.IsNull() {
		appflowparam.Httpuseragent = data.Httpuseragent.ValueString()
	}
	if !data.Httpvia.IsNull() {
		appflowparam.Httpvia = data.Httpvia.ValueString()
	}
	if !data.Httpxforwardedfor.IsNull() {
		appflowparam.Httpxforwardedfor = data.Httpxforwardedfor.ValueString()
	}
	if !data.Identifiername.IsNull() {
		appflowparam.Identifiername = data.Identifiername.ValueString()
	}
	if !data.Identifiersessionname.IsNull() {
		appflowparam.Identifiersessionname = data.Identifiersessionname.ValueString()
	}
	if !data.Logstreamovernsip.IsNull() {
		appflowparam.Logstreamovernsip = data.Logstreamovernsip.ValueString()
	}
	if !data.Lsnlogging.IsNull() {
		appflowparam.Lsnlogging = data.Lsnlogging.ValueString()
	}
	if !data.Metrics.IsNull() {
		appflowparam.Metrics = data.Metrics.ValueString()
	}
	if !data.Observationdomainid.IsNull() {
		appflowparam.Observationdomainid = utils.IntPtr(int(data.Observationdomainid.ValueInt64()))
	}
	if !data.Observationdomainname.IsNull() {
		appflowparam.Observationdomainname = data.Observationdomainname.ValueString()
	}
	if !data.Observationpointid.IsNull() {
		appflowparam.Observationpointid = utils.IntPtr(int(data.Observationpointid.ValueInt64()))
	}
	if !data.Securityinsightrecordinterval.IsNull() {
		appflowparam.Securityinsightrecordinterval = utils.IntPtr(int(data.Securityinsightrecordinterval.ValueInt64()))
	}
	if !data.Securityinsighttraffic.IsNull() {
		appflowparam.Securityinsighttraffic = data.Securityinsighttraffic.ValueString()
	}
	if !data.Skipcacheredirectionhttptransaction.IsNull() {
		appflowparam.Skipcacheredirectionhttptransaction = data.Skipcacheredirectionhttptransaction.ValueString()
	}
	if !data.Subscriberawareness.IsNull() {
		appflowparam.Subscriberawareness = data.Subscriberawareness.ValueString()
	}
	if !data.Subscriberidobfuscation.IsNull() {
		appflowparam.Subscriberidobfuscation = data.Subscriberidobfuscation.ValueString()
	}
	if !data.Subscriberidobfuscationalgo.IsNull() {
		appflowparam.Subscriberidobfuscationalgo = data.Subscriberidobfuscationalgo.ValueString()
	}
	if !data.Tcpattackcounterinterval.IsNull() {
		appflowparam.Tcpattackcounterinterval = utils.IntPtr(int(data.Tcpattackcounterinterval.ValueInt64()))
	}
	if !data.Templaterefresh.IsNull() {
		appflowparam.Templaterefresh = utils.IntPtr(int(data.Templaterefresh.ValueInt64()))
	}
	if !data.Timeseriesovernsip.IsNull() {
		appflowparam.Timeseriesovernsip = data.Timeseriesovernsip.ValueString()
	}
	if !data.Udppmtu.IsNull() {
		appflowparam.Udppmtu = utils.IntPtr(int(data.Udppmtu.ValueInt64()))
	}
	if !data.Urlcategory.IsNull() {
		appflowparam.Urlcategory = data.Urlcategory.ValueString()
	}
	if !data.Usagerecordinterval.IsNull() {
		appflowparam.Usagerecordinterval = utils.IntPtr(int(data.Usagerecordinterval.ValueInt64()))
	}
	if !data.Videoinsight.IsNull() {
		appflowparam.Videoinsight = data.Videoinsight.ValueString()
	}
	if !data.Websaasappusagereporting.IsNull() {
		appflowparam.Websaasappusagereporting = data.Websaasappusagereporting.ValueString()
	}

	return appflowparam
}

func appflowparamSetAttrFromGet(ctx context.Context, data *AppflowparamResourceModel, getResponseData map[string]interface{}) *AppflowparamResourceModel {
	tflog.Debug(ctx, "In appflowparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["aaausername"]; ok && val != nil {
		data.Aaausername = types.StringValue(val.(string))
	} else {
		data.Aaausername = types.StringNull()
	}
	if val, ok := getResponseData["analyticsauthtoken"]; ok && val != nil {
		data.Analyticsauthtoken = types.StringValue(val.(string))
	} else {
		data.Analyticsauthtoken = types.StringNull()
	}
	if val, ok := getResponseData["appnamerefresh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Appnamerefresh = types.Int64Value(intVal)
		}
	} else {
		data.Appnamerefresh = types.Int64Null()
	}
	if val, ok := getResponseData["auditlogs"]; ok && val != nil {
		data.Auditlogs = types.StringValue(val.(string))
	} else {
		data.Auditlogs = types.StringNull()
	}
	if val, ok := getResponseData["cacheinsight"]; ok && val != nil {
		data.Cacheinsight = types.StringValue(val.(string))
	} else {
		data.Cacheinsight = types.StringNull()
	}
	if val, ok := getResponseData["clienttrafficonly"]; ok && val != nil {
		data.Clienttrafficonly = types.StringValue(val.(string))
	} else {
		data.Clienttrafficonly = types.StringNull()
	}
	if val, ok := getResponseData["connectionchaining"]; ok && val != nil {
		data.Connectionchaining = types.StringValue(val.(string))
	} else {
		data.Connectionchaining = types.StringNull()
	}
	if val, ok := getResponseData["cqareporting"]; ok && val != nil {
		data.Cqareporting = types.StringValue(val.(string))
	} else {
		data.Cqareporting = types.StringNull()
	}
	if val, ok := getResponseData["distributedtracing"]; ok && val != nil {
		data.Distributedtracing = types.StringValue(val.(string))
	} else {
		data.Distributedtracing = types.StringNull()
	}
	if val, ok := getResponseData["disttracingsamplingrate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Disttracingsamplingrate = types.Int64Value(intVal)
		}
	} else {
		data.Disttracingsamplingrate = types.Int64Null()
	}
	if val, ok := getResponseData["emailaddress"]; ok && val != nil {
		data.Emailaddress = types.StringValue(val.(string))
	} else {
		data.Emailaddress = types.StringNull()
	}
	if val, ok := getResponseData["events"]; ok && val != nil {
		data.Events = types.StringValue(val.(string))
	} else {
		data.Events = types.StringNull()
	}
	if val, ok := getResponseData["flowrecordinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Flowrecordinterval = types.Int64Value(intVal)
		}
	} else {
		data.Flowrecordinterval = types.Int64Null()
	}
	if val, ok := getResponseData["gxsessionreporting"]; ok && val != nil {
		data.Gxsessionreporting = types.StringValue(val.(string))
	} else {
		data.Gxsessionreporting = types.StringNull()
	}
	if val, ok := getResponseData["httpauthorization"]; ok && val != nil {
		data.Httpauthorization = types.StringValue(val.(string))
	} else {
		data.Httpauthorization = types.StringNull()
	}
	if val, ok := getResponseData["httpcontenttype"]; ok && val != nil {
		data.Httpcontenttype = types.StringValue(val.(string))
	} else {
		data.Httpcontenttype = types.StringNull()
	}
	if val, ok := getResponseData["httpcookie"]; ok && val != nil {
		data.Httpcookie = types.StringValue(val.(string))
	} else {
		data.Httpcookie = types.StringNull()
	}
	if val, ok := getResponseData["httpdomain"]; ok && val != nil {
		data.Httpdomain = types.StringValue(val.(string))
	} else {
		data.Httpdomain = types.StringNull()
	}
	if val, ok := getResponseData["httphost"]; ok && val != nil {
		data.Httphost = types.StringValue(val.(string))
	} else {
		data.Httphost = types.StringNull()
	}
	if val, ok := getResponseData["httplocation"]; ok && val != nil {
		data.Httplocation = types.StringValue(val.(string))
	} else {
		data.Httplocation = types.StringNull()
	}
	if val, ok := getResponseData["httpmethod"]; ok && val != nil {
		data.Httpmethod = types.StringValue(val.(string))
	} else {
		data.Httpmethod = types.StringNull()
	}
	if val, ok := getResponseData["httpquerywithurl"]; ok && val != nil {
		data.Httpquerywithurl = types.StringValue(val.(string))
	} else {
		data.Httpquerywithurl = types.StringNull()
	}
	if val, ok := getResponseData["httpreferer"]; ok && val != nil {
		data.Httpreferer = types.StringValue(val.(string))
	} else {
		data.Httpreferer = types.StringNull()
	}
	if val, ok := getResponseData["httpsetcookie"]; ok && val != nil {
		data.Httpsetcookie = types.StringValue(val.(string))
	} else {
		data.Httpsetcookie = types.StringNull()
	}
	if val, ok := getResponseData["httpsetcookie2"]; ok && val != nil {
		data.Httpsetcookie2 = types.StringValue(val.(string))
	} else {
		data.Httpsetcookie2 = types.StringNull()
	}
	if val, ok := getResponseData["httpurl"]; ok && val != nil {
		data.Httpurl = types.StringValue(val.(string))
	} else {
		data.Httpurl = types.StringNull()
	}
	if val, ok := getResponseData["httpuseragent"]; ok && val != nil {
		data.Httpuseragent = types.StringValue(val.(string))
	} else {
		data.Httpuseragent = types.StringNull()
	}
	if val, ok := getResponseData["httpvia"]; ok && val != nil {
		data.Httpvia = types.StringValue(val.(string))
	} else {
		data.Httpvia = types.StringNull()
	}
	if val, ok := getResponseData["httpxforwardedfor"]; ok && val != nil {
		data.Httpxforwardedfor = types.StringValue(val.(string))
	} else {
		data.Httpxforwardedfor = types.StringNull()
	}
	if val, ok := getResponseData["identifiername"]; ok && val != nil {
		data.Identifiername = types.StringValue(val.(string))
	} else {
		data.Identifiername = types.StringNull()
	}
	if val, ok := getResponseData["identifiersessionname"]; ok && val != nil {
		data.Identifiersessionname = types.StringValue(val.(string))
	} else {
		data.Identifiersessionname = types.StringNull()
	}
	if val, ok := getResponseData["logstreamovernsip"]; ok && val != nil {
		data.Logstreamovernsip = types.StringValue(val.(string))
	} else {
		data.Logstreamovernsip = types.StringNull()
	}
	if val, ok := getResponseData["lsnlogging"]; ok && val != nil {
		data.Lsnlogging = types.StringValue(val.(string))
	} else {
		data.Lsnlogging = types.StringNull()
	}
	if val, ok := getResponseData["metrics"]; ok && val != nil {
		data.Metrics = types.StringValue(val.(string))
	} else {
		data.Metrics = types.StringNull()
	}
	if val, ok := getResponseData["observationdomainid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Observationdomainid = types.Int64Value(intVal)
		}
	} else {
		data.Observationdomainid = types.Int64Null()
	}
	if val, ok := getResponseData["observationdomainname"]; ok && val != nil {
		data.Observationdomainname = types.StringValue(val.(string))
	} else {
		data.Observationdomainname = types.StringNull()
	}
	if val, ok := getResponseData["observationpointid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Observationpointid = types.Int64Value(intVal)
		}
	} else {
		data.Observationpointid = types.Int64Null()
	}
	if val, ok := getResponseData["securityinsightrecordinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Securityinsightrecordinterval = types.Int64Value(intVal)
		}
	} else {
		data.Securityinsightrecordinterval = types.Int64Null()
	}
	if val, ok := getResponseData["securityinsighttraffic"]; ok && val != nil {
		data.Securityinsighttraffic = types.StringValue(val.(string))
	} else {
		data.Securityinsighttraffic = types.StringNull()
	}
	if val, ok := getResponseData["skipcacheredirectionhttptransaction"]; ok && val != nil {
		data.Skipcacheredirectionhttptransaction = types.StringValue(val.(string))
	} else {
		data.Skipcacheredirectionhttptransaction = types.StringNull()
	}
	if val, ok := getResponseData["subscriberawareness"]; ok && val != nil {
		data.Subscriberawareness = types.StringValue(val.(string))
	} else {
		data.Subscriberawareness = types.StringNull()
	}
	if val, ok := getResponseData["subscriberidobfuscation"]; ok && val != nil {
		data.Subscriberidobfuscation = types.StringValue(val.(string))
	} else {
		data.Subscriberidobfuscation = types.StringNull()
	}
	if val, ok := getResponseData["subscriberidobfuscationalgo"]; ok && val != nil {
		data.Subscriberidobfuscationalgo = types.StringValue(val.(string))
	} else {
		data.Subscriberidobfuscationalgo = types.StringNull()
	}
	if val, ok := getResponseData["tcpattackcounterinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpattackcounterinterval = types.Int64Value(intVal)
		}
	} else {
		data.Tcpattackcounterinterval = types.Int64Null()
	}
	if val, ok := getResponseData["templaterefresh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Templaterefresh = types.Int64Value(intVal)
		}
	} else {
		data.Templaterefresh = types.Int64Null()
	}
	if val, ok := getResponseData["timeseriesovernsip"]; ok && val != nil {
		data.Timeseriesovernsip = types.StringValue(val.(string))
	} else {
		data.Timeseriesovernsip = types.StringNull()
	}
	if val, ok := getResponseData["udppmtu"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Udppmtu = types.Int64Value(intVal)
		}
	} else {
		data.Udppmtu = types.Int64Null()
	}
	if val, ok := getResponseData["urlcategory"]; ok && val != nil {
		data.Urlcategory = types.StringValue(val.(string))
	} else {
		data.Urlcategory = types.StringNull()
	}
	if val, ok := getResponseData["usagerecordinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Usagerecordinterval = types.Int64Value(intVal)
		}
	} else {
		data.Usagerecordinterval = types.Int64Null()
	}
	if val, ok := getResponseData["videoinsight"]; ok && val != nil {
		data.Videoinsight = types.StringValue(val.(string))
	} else {
		data.Videoinsight = types.StringNull()
	}
	if val, ok := getResponseData["websaasappusagereporting"]; ok && val != nil {
		data.Websaasappusagereporting = types.StringValue(val.(string))
	} else {
		data.Websaasappusagereporting = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("appflowparam-config")

	return data
}
