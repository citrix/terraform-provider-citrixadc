package analyticsprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/analytics"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AnalyticsprofileResourceModel describes the resource data model.
type AnalyticsprofileResourceModel struct {
	Id                           types.String `tfsdk:"id"`
	Allhttpheaders               types.String `tfsdk:"allhttpheaders"`
	Analyticsauthtoken           types.String `tfsdk:"analyticsauthtoken"`
	Analyticsendpointcontenttype types.String `tfsdk:"analyticsendpointcontenttype"`
	Analyticsendpointmetadata    types.String `tfsdk:"analyticsendpointmetadata"`
	Analyticsendpointurl         types.String `tfsdk:"analyticsendpointurl"`
	Auditlogs                    types.String `tfsdk:"auditlogs"`
	Collectors                   types.String `tfsdk:"collectors"`
	Cqareporting                 types.String `tfsdk:"cqareporting"`
	Dataformatfile               types.String `tfsdk:"dataformatfile"`
	Events                       types.String `tfsdk:"events"`
	Grpcstatus                   types.String `tfsdk:"grpcstatus"`
	Httpauthentication           types.String `tfsdk:"httpauthentication"`
	Httpclientsidemeasurements   types.String `tfsdk:"httpclientsidemeasurements"`
	Httpcontenttype              types.String `tfsdk:"httpcontenttype"`
	Httpcookie                   types.String `tfsdk:"httpcookie"`
	Httpcustomheaders            types.List   `tfsdk:"httpcustomheaders"`
	Httpdomainname               types.String `tfsdk:"httpdomainname"`
	Httphost                     types.String `tfsdk:"httphost"`
	Httplocation                 types.String `tfsdk:"httplocation"`
	Httpmethod                   types.String `tfsdk:"httpmethod"`
	Httppagetracking             types.String `tfsdk:"httppagetracking"`
	Httpreferer                  types.String `tfsdk:"httpreferer"`
	Httpsetcookie                types.String `tfsdk:"httpsetcookie"`
	Httpsetcookie2               types.String `tfsdk:"httpsetcookie2"`
	Httpurl                      types.String `tfsdk:"httpurl"`
	Httpurlquery                 types.String `tfsdk:"httpurlquery"`
	Httpuseragent                types.String `tfsdk:"httpuseragent"`
	Httpvia                      types.String `tfsdk:"httpvia"`
	Httpxforwardedforheader      types.String `tfsdk:"httpxforwardedforheader"`
	Integratedcache              types.String `tfsdk:"integratedcache"`
	Managementlog                types.List   `tfsdk:"managementlog"`
	Metrics                      types.String `tfsdk:"metrics"`
	Metricsexportfrequency       types.Int64  `tfsdk:"metricsexportfrequency"`
	Name                         types.String `tfsdk:"name"`
	Outputmode                   types.String `tfsdk:"outputmode"`
	Schemafile                   types.String `tfsdk:"schemafile"`
	Servemode                    types.String `tfsdk:"servemode"`
	Tcpburstreporting            types.String `tfsdk:"tcpburstreporting"`
	Topn                         types.String `tfsdk:"topn"`
	Type                         types.String `tfsdk:"type"`
	Urlcategory                  types.String `tfsdk:"urlcategory"`
}

func (r *AnalyticsprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the analyticsprofile resource.",
			},
			"allhttpheaders": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option indicates the whether auditlog should be sent to the REST collector.",
			},
			"collectors": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The collector can be an IP, an appflow collector name, a service or a vserver. If IP is specified, the transport is considered as logstream and default port of 5557 is taken. If collector name is specified, the collector properties are taken from the configured collector. If service is specified, the configured service is assumed as the collector. If vserver is specified, the services bound to it are considered as collectors and the records are load balanced.",
			},
			"cqareporting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log TCP CQA parameters.",
			},
			"dataformatfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is for configuring the file containing the data format and metadata required by the analytics endpoint.",
			},
			"events": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option indicates the whether events should be sent to the REST collector.",
			},
			"grpcstatus": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log the gRPC status headers",
			},
			"httpauthentication": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log Authentication header.",
			},
			"httpclientsidemeasurements": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will insert a javascript into the HTTP response to collect the client side page-timings and will send the same to the configured collectors.",
			},
			"httpcontenttype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log content-length header.",
			},
			"httpcookie": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log domain name.",
			},
			"httphost": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log the Host header in appflow records",
			},
			"httplocation": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log location header.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log the method header in appflow records",
			},
			"httppagetracking": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will link the embedded objects of a page together.",
			},
			"httpreferer": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log the referer header in appflow records",
			},
			"httpsetcookie": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log set-cookie header.",
			},
			"httpsetcookie2": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log set-cookie2 header.",
			},
			"httpurl": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log the URL in appflow records",
			},
			"httpurlquery": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log URL Query.",
			},
			"httpuseragent": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log User-Agent header.",
			},
			"httpvia": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will Via header.",
			},
			"httpxforwardedforheader": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log X-Forwarded-For header.",
			},
			"integratedcache": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will log the Integrated Caching appflow records",
			},
			"managementlog": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "This option indicates the whether managementlog should be sent to the REST collector.",
			},
			"metrics": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option indicates the whether metrics should be sent to the REST collector.",
			},
			"metricsexportfrequency": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(30),
				Description: "This option is for configuring the metrics export frequency in seconds, frequency value must be in [30,300] seconds range",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the analytics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow profile\" or 'my appflow profile').",
			},
			"outputmode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("avro"),
				Description: "This option indicates the format of REST API POST body. It depends on the consumer of the analytics data.",
			},
			"schemafile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is for configuring json schema file containing a list of counters to be exported by metricscollector",
			},
			"servemode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Push"),
				Description: "This option is for setting the mode of how data is provided",
			},
			"tcpburstreporting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "On enabling this option, the Citrix ADC will log TCP burst parameters.",
			},
			"topn": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this topn support, the topn information of the stream identifier this profile is bound to will be exported to the analytics endpoint.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "This option indicates what information needs to be collected and exported.",
			},
			"urlcategory": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will send the URL category record.",
			},
		},
	}
}

func analyticsprofileGetThePayloadFromtheConfig(ctx context.Context, data *AnalyticsprofileResourceModel) analytics.Analyticsprofile {
	tflog.Debug(ctx, "In analyticsprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	analyticsprofile := analytics.Analyticsprofile{}
	if !data.Allhttpheaders.IsNull() {
		analyticsprofile.Allhttpheaders = data.Allhttpheaders.ValueString()
	}
	if !data.Analyticsauthtoken.IsNull() {
		analyticsprofile.Analyticsauthtoken = data.Analyticsauthtoken.ValueString()
	}
	if !data.Analyticsendpointcontenttype.IsNull() {
		analyticsprofile.Analyticsendpointcontenttype = data.Analyticsendpointcontenttype.ValueString()
	}
	if !data.Analyticsendpointmetadata.IsNull() {
		analyticsprofile.Analyticsendpointmetadata = data.Analyticsendpointmetadata.ValueString()
	}
	if !data.Analyticsendpointurl.IsNull() {
		analyticsprofile.Analyticsendpointurl = data.Analyticsendpointurl.ValueString()
	}
	if !data.Auditlogs.IsNull() {
		analyticsprofile.Auditlogs = data.Auditlogs.ValueString()
	}
	if !data.Collectors.IsNull() {
		analyticsprofile.Collectors = data.Collectors.ValueString()
	}
	if !data.Cqareporting.IsNull() {
		analyticsprofile.Cqareporting = data.Cqareporting.ValueString()
	}
	if !data.Dataformatfile.IsNull() {
		analyticsprofile.Dataformatfile = data.Dataformatfile.ValueString()
	}
	if !data.Events.IsNull() {
		analyticsprofile.Events = data.Events.ValueString()
	}
	if !data.Grpcstatus.IsNull() {
		analyticsprofile.Grpcstatus = data.Grpcstatus.ValueString()
	}
	if !data.Httpauthentication.IsNull() {
		analyticsprofile.Httpauthentication = data.Httpauthentication.ValueString()
	}
	if !data.Httpclientsidemeasurements.IsNull() {
		analyticsprofile.Httpclientsidemeasurements = data.Httpclientsidemeasurements.ValueString()
	}
	if !data.Httpcontenttype.IsNull() {
		analyticsprofile.Httpcontenttype = data.Httpcontenttype.ValueString()
	}
	if !data.Httpcookie.IsNull() {
		analyticsprofile.Httpcookie = data.Httpcookie.ValueString()
	}
	if !data.Httpdomainname.IsNull() {
		analyticsprofile.Httpdomainname = data.Httpdomainname.ValueString()
	}
	if !data.Httphost.IsNull() {
		analyticsprofile.Httphost = data.Httphost.ValueString()
	}
	if !data.Httplocation.IsNull() {
		analyticsprofile.Httplocation = data.Httplocation.ValueString()
	}
	if !data.Httpmethod.IsNull() {
		analyticsprofile.Httpmethod = data.Httpmethod.ValueString()
	}
	if !data.Httppagetracking.IsNull() {
		analyticsprofile.Httppagetracking = data.Httppagetracking.ValueString()
	}
	if !data.Httpreferer.IsNull() {
		analyticsprofile.Httpreferer = data.Httpreferer.ValueString()
	}
	if !data.Httpsetcookie.IsNull() {
		analyticsprofile.Httpsetcookie = data.Httpsetcookie.ValueString()
	}
	if !data.Httpsetcookie2.IsNull() {
		analyticsprofile.Httpsetcookie2 = data.Httpsetcookie2.ValueString()
	}
	if !data.Httpurl.IsNull() {
		analyticsprofile.Httpurl = data.Httpurl.ValueString()
	}
	if !data.Httpurlquery.IsNull() {
		analyticsprofile.Httpurlquery = data.Httpurlquery.ValueString()
	}
	if !data.Httpuseragent.IsNull() {
		analyticsprofile.Httpuseragent = data.Httpuseragent.ValueString()
	}
	if !data.Httpvia.IsNull() {
		analyticsprofile.Httpvia = data.Httpvia.ValueString()
	}
	if !data.Httpxforwardedforheader.IsNull() {
		analyticsprofile.Httpxforwardedforheader = data.Httpxforwardedforheader.ValueString()
	}
	if !data.Integratedcache.IsNull() {
		analyticsprofile.Integratedcache = data.Integratedcache.ValueString()
	}
	if !data.Metrics.IsNull() {
		analyticsprofile.Metrics = data.Metrics.ValueString()
	}
	if !data.Metricsexportfrequency.IsNull() {
		analyticsprofile.Metricsexportfrequency = utils.IntPtr(int(data.Metricsexportfrequency.ValueInt64()))
	}
	if !data.Name.IsNull() {
		analyticsprofile.Name = data.Name.ValueString()
	}
	if !data.Outputmode.IsNull() {
		analyticsprofile.Outputmode = data.Outputmode.ValueString()
	}
	if !data.Schemafile.IsNull() {
		analyticsprofile.Schemafile = data.Schemafile.ValueString()
	}
	if !data.Servemode.IsNull() {
		analyticsprofile.Servemode = data.Servemode.ValueString()
	}
	if !data.Tcpburstreporting.IsNull() {
		analyticsprofile.Tcpburstreporting = data.Tcpburstreporting.ValueString()
	}
	if !data.Topn.IsNull() {
		analyticsprofile.Topn = data.Topn.ValueString()
	}
	if !data.Type.IsNull() {
		analyticsprofile.Type = data.Type.ValueString()
	}
	if !data.Urlcategory.IsNull() {
		analyticsprofile.Urlcategory = data.Urlcategory.ValueString()
	}

	return analyticsprofile
}

func analyticsprofileSetAttrFromGet(ctx context.Context, data *AnalyticsprofileResourceModel, getResponseData map[string]interface{}) *AnalyticsprofileResourceModel {
	tflog.Debug(ctx, "In analyticsprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["allhttpheaders"]; ok && val != nil {
		data.Allhttpheaders = types.StringValue(val.(string))
	} else {
		data.Allhttpheaders = types.StringNull()
	}
	if val, ok := getResponseData["analyticsauthtoken"]; ok && val != nil {
		data.Analyticsauthtoken = types.StringValue(val.(string))
	} else {
		data.Analyticsauthtoken = types.StringNull()
	}
	if val, ok := getResponseData["analyticsendpointcontenttype"]; ok && val != nil {
		data.Analyticsendpointcontenttype = types.StringValue(val.(string))
	} else {
		data.Analyticsendpointcontenttype = types.StringNull()
	}
	if val, ok := getResponseData["analyticsendpointmetadata"]; ok && val != nil {
		data.Analyticsendpointmetadata = types.StringValue(val.(string))
	} else {
		data.Analyticsendpointmetadata = types.StringNull()
	}
	if val, ok := getResponseData["analyticsendpointurl"]; ok && val != nil {
		data.Analyticsendpointurl = types.StringValue(val.(string))
	} else {
		data.Analyticsendpointurl = types.StringNull()
	}
	if val, ok := getResponseData["auditlogs"]; ok && val != nil {
		data.Auditlogs = types.StringValue(val.(string))
	} else {
		data.Auditlogs = types.StringNull()
	}
	if val, ok := getResponseData["collectors"]; ok && val != nil {
		data.Collectors = types.StringValue(val.(string))
	} else {
		data.Collectors = types.StringNull()
	}
	if val, ok := getResponseData["cqareporting"]; ok && val != nil {
		data.Cqareporting = types.StringValue(val.(string))
	} else {
		data.Cqareporting = types.StringNull()
	}
	if val, ok := getResponseData["dataformatfile"]; ok && val != nil {
		data.Dataformatfile = types.StringValue(val.(string))
	} else {
		data.Dataformatfile = types.StringNull()
	}
	if val, ok := getResponseData["events"]; ok && val != nil {
		data.Events = types.StringValue(val.(string))
	} else {
		data.Events = types.StringNull()
	}
	if val, ok := getResponseData["grpcstatus"]; ok && val != nil {
		data.Grpcstatus = types.StringValue(val.(string))
	} else {
		data.Grpcstatus = types.StringNull()
	}
	if val, ok := getResponseData["httpauthentication"]; ok && val != nil {
		data.Httpauthentication = types.StringValue(val.(string))
	} else {
		data.Httpauthentication = types.StringNull()
	}
	if val, ok := getResponseData["httpclientsidemeasurements"]; ok && val != nil {
		data.Httpclientsidemeasurements = types.StringValue(val.(string))
	} else {
		data.Httpclientsidemeasurements = types.StringNull()
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
	if val, ok := getResponseData["httpdomainname"]; ok && val != nil {
		data.Httpdomainname = types.StringValue(val.(string))
	} else {
		data.Httpdomainname = types.StringNull()
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
	if val, ok := getResponseData["httppagetracking"]; ok && val != nil {
		data.Httppagetracking = types.StringValue(val.(string))
	} else {
		data.Httppagetracking = types.StringNull()
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
	if val, ok := getResponseData["httpurlquery"]; ok && val != nil {
		data.Httpurlquery = types.StringValue(val.(string))
	} else {
		data.Httpurlquery = types.StringNull()
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
	if val, ok := getResponseData["httpxforwardedforheader"]; ok && val != nil {
		data.Httpxforwardedforheader = types.StringValue(val.(string))
	} else {
		data.Httpxforwardedforheader = types.StringNull()
	}
	if val, ok := getResponseData["integratedcache"]; ok && val != nil {
		data.Integratedcache = types.StringValue(val.(string))
	} else {
		data.Integratedcache = types.StringNull()
	}
	if val, ok := getResponseData["metrics"]; ok && val != nil {
		data.Metrics = types.StringValue(val.(string))
	} else {
		data.Metrics = types.StringNull()
	}
	if val, ok := getResponseData["metricsexportfrequency"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Metricsexportfrequency = types.Int64Value(intVal)
		}
	} else {
		data.Metricsexportfrequency = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["outputmode"]; ok && val != nil {
		data.Outputmode = types.StringValue(val.(string))
	} else {
		data.Outputmode = types.StringNull()
	}
	if val, ok := getResponseData["schemafile"]; ok && val != nil {
		data.Schemafile = types.StringValue(val.(string))
	} else {
		data.Schemafile = types.StringNull()
	}
	if val, ok := getResponseData["servemode"]; ok && val != nil {
		data.Servemode = types.StringValue(val.(string))
	} else {
		data.Servemode = types.StringNull()
	}
	if val, ok := getResponseData["tcpburstreporting"]; ok && val != nil {
		data.Tcpburstreporting = types.StringValue(val.(string))
	} else {
		data.Tcpburstreporting = types.StringNull()
	}
	if val, ok := getResponseData["topn"]; ok && val != nil {
		data.Topn = types.StringValue(val.(string))
	} else {
		data.Topn = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["urlcategory"]; ok && val != nil {
		data.Urlcategory = types.StringValue(val.(string))
	} else {
		data.Urlcategory = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
