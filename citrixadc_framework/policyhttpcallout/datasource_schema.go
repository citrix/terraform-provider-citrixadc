package policyhttpcallout

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicyhttpcalloutDataSourceModel is the data-source-specific model, decoupled
// from PolicyhttpcalloutResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime attributes the resource deliberately omits
// (hits, undefhits, svrstate, effectivestate, undefreason, recursivecallout).
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type PolicyhttpcalloutDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"` // Required lookup key
	Bodyexpr     types.String `tfsdk:"bodyexpr"`
	Cacheforsecs types.Int64  `tfsdk:"cacheforsecs"`
	Comment      types.String `tfsdk:"comment"`
	Fullreqexpr  types.String `tfsdk:"fullreqexpr"`
	Headers      types.List   `tfsdk:"headers"`
	Hostexpr     types.String `tfsdk:"hostexpr"`
	Httpmethod   types.String `tfsdk:"httpmethod"`
	Ipaddress    types.String `tfsdk:"ipaddress"`
	Parameters   types.List   `tfsdk:"parameters"`
	Port         types.Int64  `tfsdk:"port"`
	Resultexpr   types.String `tfsdk:"resultexpr"`
	Returntype   types.String `tfsdk:"returntype"`
	Scheme       types.String `tfsdk:"scheme"`
	Urlstemexpr  types.String `tfsdk:"urlstemexpr"`
	Vserver      types.String `tfsdk:"vserver"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/policyhttpcallout.json). Never settable; populated from GET.
	Hits             types.Int64  `tfsdk:"hits"`
	Undefhits        types.Int64  `tfsdk:"undefhits"`
	Svrstate         types.String `tfsdk:"svrstate"`
	Effectivestate   types.String `tfsdk:"effectivestate"`
	Undefreason      types.String `tfsdk:"undefreason"`
	Recursivecallout types.Int64  `tfsdk:"recursivecallout"`
}

func PolicyhttpcalloutDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bodyexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An advanced string expression for generating the body of the request. The expression can contain a literal string or an expression that derives the value (for example, client.ip.src). Mutually exclusive with -fullReqExpr.",
			},
			"cacheforsecs": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Duration, in seconds, for which the callout response is cached. The cached responses are stored in an integrated caching content group named \"calloutContentGroup\". If no duration is configured, the callout responses will not be cached unless normal caching configuration is used to cache them. This parameter takes precedence over any normal caching configuration that would otherwise apply to these responses.\n	   Note that the calloutContentGroup definition may not be modified or removed nor may it be used with other cache policies.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this HTTP callout.",
			},
			"fullreqexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exact HTTP request, in the form of an expression, which the Citrix ADC sends to the callout agent. If you set this parameter, you must not include HTTP method, host expression, URL stem expression, headers, or parameters.\nThe request expression is constrained by the feature for which the callout is used. For example, an HTTP.RES expression cannot be used in a request-time policy bank or in a TCP content switching policy bank.\nThe Citrix ADC does not check the validity of this request. You must manually validate the request.",
			},
			"headers": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more headers to insert into the HTTP request. Each header is specified as \"name(expr)\", where expr is an expression that is evaluated at runtime to provide the value for the named header. You can configure a maximum of eight headers for an HTTP callout. Mutually exclusive with the full HTTP request expression.",
			},
			"hostexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String expression to configure the Host header. Can contain a literal value (for example, 10.101.10.11) or a derived value (for example, http.req.header(\"Host\")). The literal value can be an IP address or a fully qualified domain name. Mutually exclusive with the full HTTP request expression.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Method used in the HTTP request that this callout sends.  Mutually exclusive with the full HTTP request expression.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP Address of the server (callout agent) to which the callout is sent. Can be an IPv4 or IPv6 address.\nMutually exclusive with the Virtual Server parameter. Therefore, you cannot set the <IP Address, Port> and the Virtual Server in the same HTTP callout.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the HTTP callout. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or HTTP callout.",
			},
			"parameters": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more query parameters to insert into the HTTP request URL (for a GET request) or into the request body (for a POST request). Each parameter is specified as \"name(expr)\", where expr is an expression that is evaluated at run time to provide the value for the named parameter (name=value). The parameter values are URL encoded. Mutually exclusive with the full HTTP request expression.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Server port to which the HTTP callout agent is mapped. Mutually exclusive with the Virtual Server parameter. Therefore, you cannot set the <IP Address, Port> and the Virtual Server in the same HTTP callout.",
			},
			"resultexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that extracts the callout results from the response sent by the HTTP callout agent. Must be a response based expression, that is, it must begin with HTTP.RES. The operations in this expression must match the return type. For example, if you configure a return type of TEXT, the result expression must be a text based expression. If the return type is NUM, the result expression (resultExpr) must return a numeric value, as in the following example: http.res.body(10000).length.",
			},
			"returntype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of data that the target callout agent returns in response to the callout. \nAvailable settings function as follows:\n* TEXT - Treat the returned value as a text string. \n* NUM - Treat the returned value as a number.\n* BOOL - Treat the returned value as a Boolean value. \nNote: You cannot change the return type after it is set.",
			},
			"scheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of scheme for the callout server.",
			},
			"urlstemexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String expression for generating the URL stem. Can contain a literal string (for example, \"/mysite/index.html\") or an expression that derives the value (for example, http.req.url). Mutually exclusive with the full HTTP request expression.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing, content switching, or cache redirection virtual server (the callout agent) to which the HTTP callout is sent. The service type of the virtual server must be HTTP. Mutually exclusive with the IP address and port parameters. Therefore, you cannot set the <IP Address, Port> and the Virtual Server in the same HTTP callout.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Total hits.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Total undefs.",
			},
			"svrstate": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the service. Possible values: UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.",
			},
			"effectivestate": schema.StringAttribute{
				Computed:    true,
				Description: "The effective state of the service. Possible values: UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.",
			},
			"undefreason": schema.StringAttribute{
				Computed:    true,
				Description: "Reason for last undef.",
			},
			"recursivecallout": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of recursive callouts.",
			},
		},
	}
}

// policyhttpcalloutDataSourceSetAttrFromGet projects a NITRO policyhttpcallout
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func policyhttpcalloutDataSourceSetAttrFromGet(ctx context.Context, data *PolicyhttpcalloutDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In policyhttpcalloutDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Bodyexpr = utils.MapGetString(g, "bodyexpr")
	data.Cacheforsecs = utils.MapGetInt64(g, "cacheforsecs")
	data.Comment = utils.MapGetString(g, "comment")
	data.Fullreqexpr = utils.MapGetString(g, "fullreqexpr")
	data.Headers = utils.MapGetStringList(g, "headers")
	data.Hostexpr = utils.MapGetString(g, "hostexpr")
	data.Httpmethod = utils.MapGetString(g, "httpmethod")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Parameters = utils.MapGetStringList(g, "parameters")
	data.Port = utils.MapGetInt64(g, "port")
	data.Resultexpr = utils.MapGetString(g, "resultexpr")
	data.Returntype = utils.MapGetString(g, "returntype")
	data.Scheme = utils.MapGetString(g, "scheme")
	data.Urlstemexpr = utils.MapGetString(g, "urlstemexpr")
	data.Vserver = utils.MapGetString(g, "vserver")

	// Read-only runtime metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Svrstate = utils.MapGetString(g, "svrstate")
	data.Effectivestate = utils.MapGetString(g, "effectivestate")
	data.Undefreason = utils.MapGetString(g, "undefreason")
	data.Recursivecallout = utils.MapGetInt64(g, "recursivecallout")
}
