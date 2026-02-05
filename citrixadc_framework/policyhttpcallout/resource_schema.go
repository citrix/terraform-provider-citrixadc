package policyhttpcallout

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// PolicyhttpcalloutResourceModel describes the resource data model.
type PolicyhttpcalloutResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Bodyexpr     types.String `tfsdk:"bodyexpr"`
	Cacheforsecs types.Int64  `tfsdk:"cacheforsecs"`
	Comment      types.String `tfsdk:"comment"`
	Fullreqexpr  types.String `tfsdk:"fullreqexpr"`
	Headers      types.List   `tfsdk:"headers"`
	Hostexpr     types.String `tfsdk:"hostexpr"`
	Httpmethod   types.String `tfsdk:"httpmethod"`
	Ipaddress    types.String `tfsdk:"ipaddress"`
	Name         types.String `tfsdk:"name"`
	Parameters   types.List   `tfsdk:"parameters"`
	Port         types.Int64  `tfsdk:"port"`
	Resultexpr   types.String `tfsdk:"resultexpr"`
	Returntype   types.String `tfsdk:"returntype"`
	Scheme       types.String `tfsdk:"scheme"`
	Urlstemexpr  types.String `tfsdk:"urlstemexpr"`
	Vserver      types.String `tfsdk:"vserver"`
}

func (r *PolicyhttpcalloutResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policyhttpcallout resource.",
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
		},
	}
}

func policyhttpcalloutGetThePayloadFromtheConfig(ctx context.Context, data *PolicyhttpcalloutResourceModel) policy.Policyhttpcallout {
	tflog.Debug(ctx, "In policyhttpcalloutGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	policyhttpcallout := policy.Policyhttpcallout{}
	if !data.Bodyexpr.IsNull() {
		policyhttpcallout.Bodyexpr = data.Bodyexpr.ValueString()
	}
	if !data.Cacheforsecs.IsNull() {
		policyhttpcallout.Cacheforsecs = utils.IntPtr(int(data.Cacheforsecs.ValueInt64()))
	}
	if !data.Comment.IsNull() {
		policyhttpcallout.Comment = data.Comment.ValueString()
	}
	if !data.Fullreqexpr.IsNull() {
		policyhttpcallout.Fullreqexpr = data.Fullreqexpr.ValueString()
	}
	if !data.Hostexpr.IsNull() {
		policyhttpcallout.Hostexpr = data.Hostexpr.ValueString()
	}
	if !data.Httpmethod.IsNull() {
		policyhttpcallout.Httpmethod = data.Httpmethod.ValueString()
	}
	if !data.Ipaddress.IsNull() {
		policyhttpcallout.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() {
		policyhttpcallout.Name = data.Name.ValueString()
	}
	if !data.Port.IsNull() {
		policyhttpcallout.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Resultexpr.IsNull() {
		policyhttpcallout.Resultexpr = data.Resultexpr.ValueString()
	}
	if !data.Returntype.IsNull() {
		policyhttpcallout.Returntype = data.Returntype.ValueString()
	}
	if !data.Scheme.IsNull() {
		policyhttpcallout.Scheme = data.Scheme.ValueString()
	}
	if !data.Urlstemexpr.IsNull() {
		policyhttpcallout.Urlstemexpr = data.Urlstemexpr.ValueString()
	}
	if !data.Vserver.IsNull() {
		policyhttpcallout.Vserver = data.Vserver.ValueString()
	}

	return policyhttpcallout
}

func policyhttpcalloutSetAttrFromGet(ctx context.Context, data *PolicyhttpcalloutResourceModel, getResponseData map[string]interface{}) *PolicyhttpcalloutResourceModel {
	tflog.Debug(ctx, "In policyhttpcalloutSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bodyexpr"]; ok && val != nil {
		data.Bodyexpr = types.StringValue(val.(string))
	} else {
		data.Bodyexpr = types.StringNull()
	}
	if val, ok := getResponseData["cacheforsecs"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cacheforsecs = types.Int64Value(intVal)
		}
	} else {
		data.Cacheforsecs = types.Int64Null()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["fullreqexpr"]; ok && val != nil {
		data.Fullreqexpr = types.StringValue(val.(string))
	} else {
		data.Fullreqexpr = types.StringNull()
	}
	if val, ok := getResponseData["hostexpr"]; ok && val != nil {
		data.Hostexpr = types.StringValue(val.(string))
	} else {
		data.Hostexpr = types.StringNull()
	}
	if val, ok := getResponseData["httpmethod"]; ok && val != nil {
		data.Httpmethod = types.StringValue(val.(string))
	} else {
		data.Httpmethod = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["resultexpr"]; ok && val != nil {
		data.Resultexpr = types.StringValue(val.(string))
	} else {
		data.Resultexpr = types.StringNull()
	}
	if val, ok := getResponseData["returntype"]; ok && val != nil {
		data.Returntype = types.StringValue(val.(string))
	} else {
		data.Returntype = types.StringNull()
	}
	if val, ok := getResponseData["scheme"]; ok && val != nil {
		data.Scheme = types.StringValue(val.(string))
	} else {
		data.Scheme = types.StringNull()
	}
	if val, ok := getResponseData["urlstemexpr"]; ok && val != nil {
		data.Urlstemexpr = types.StringValue(val.(string))
	} else {
		data.Urlstemexpr = types.StringNull()
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	} else {
		data.Vserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
