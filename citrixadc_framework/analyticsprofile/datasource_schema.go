package analyticsprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AnalyticsprofileDataSourceModel is the data-source-specific model, decoupled
// from AnalyticsprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (e.g. refcnt). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type AnalyticsprofileDataSourceModel struct {
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
	Mcpsummary                   types.String `tfsdk:"mcpsummary"`
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

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/analyticsprofile.json). Never settable; populated from GET.
	Refcnt types.Int64 `tfsdk:"refcnt"`
}

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
				Sensitive:   true,
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
			"mcpsummary": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/disable appflow logging for MCP (Model Context Protocol) traffic.",
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

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"refcnt": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of references to the profile.",
			},
		},
	}
}

// analyticsprofileDataSourceSetAttrFromGet projects a NITRO analyticsprofile GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) - no unknown->null resolution or plan preservation is
// required. The shared utils.MapGet* helpers implement that projection.
func analyticsprofileDataSourceSetAttrFromGet(ctx context.Context, data *AnalyticsprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In analyticsprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Allhttpheaders = utils.MapGetString(g, "allhttpheaders")
	data.Analyticsendpointcontenttype = utils.MapGetString(g, "analyticsendpointcontenttype")
	data.Analyticsendpointmetadata = utils.MapGetString(g, "analyticsendpointmetadata")
	data.Analyticsendpointurl = utils.MapGetString(g, "analyticsendpointurl")
	data.Auditlogs = utils.MapGetString(g, "auditlogs")
	data.Collectors = utils.MapGetString(g, "collectors")
	data.Cqareporting = utils.MapGetString(g, "cqareporting")
	data.Dataformatfile = utils.MapGetString(g, "dataformatfile")
	data.Events = utils.MapGetString(g, "events")
	data.Grpcstatus = utils.MapGetString(g, "grpcstatus")
	data.Httpauthentication = utils.MapGetString(g, "httpauthentication")
	data.Httpclientsidemeasurements = utils.MapGetString(g, "httpclientsidemeasurements")
	data.Httpcontenttype = utils.MapGetString(g, "httpcontenttype")
	data.Httpcookie = utils.MapGetString(g, "httpcookie")
	data.Httpcustomheaders = utils.MapGetStringList(g, "httpcustomheaders")
	data.Httpdomainname = utils.MapGetString(g, "httpdomainname")
	data.Httphost = utils.MapGetString(g, "httphost")
	data.Httplocation = utils.MapGetString(g, "httplocation")
	data.Httpmethod = utils.MapGetString(g, "httpmethod")
	data.Httppagetracking = utils.MapGetString(g, "httppagetracking")
	data.Httpreferer = utils.MapGetString(g, "httpreferer")
	data.Httpsetcookie = utils.MapGetString(g, "httpsetcookie")
	data.Httpsetcookie2 = utils.MapGetString(g, "httpsetcookie2")
	data.Httpurl = utils.MapGetString(g, "httpurl")
	data.Httpurlquery = utils.MapGetString(g, "httpurlquery")
	data.Httpuseragent = utils.MapGetString(g, "httpuseragent")
	data.Httpvia = utils.MapGetString(g, "httpvia")
	data.Httpxforwardedforheader = utils.MapGetString(g, "httpxforwardedforheader")
	data.Integratedcache = utils.MapGetString(g, "integratedcache")
	data.Managementlog = utils.MapGetStringList(g, "managementlog")
	data.Mcpsummary = utils.MapGetString(g, "mcpsummary")
	data.Metrics = utils.MapGetString(g, "metrics")
	data.Metricsexportfrequency = utils.MapGetInt64(g, "metricsexportfrequency")
	data.Outputmode = utils.MapGetString(g, "outputmode")
	data.Schemafile = utils.MapGetString(g, "schemafile")
	data.Servemode = utils.MapGetString(g, "servemode")
	data.Tcpburstreporting = utils.MapGetString(g, "tcpburstreporting")
	data.Topn = utils.MapGetString(g, "topn")
	data.Type = utils.MapGetString(g, "type")
	data.Urlcategory = utils.MapGetString(g, "urlcategory")

	// Read-only metadata.
	data.Refcnt = utils.MapGetInt64(g, "refcnt")

	// analyticsauthtoken is a secret input the GET never returns -> Null.
	data.Analyticsauthtoken = types.StringNull()
}
