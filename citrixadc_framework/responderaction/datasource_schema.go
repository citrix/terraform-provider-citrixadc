package responderaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ResponderactionDataSourceModel is the data-source-specific model, decoupled
// from ResponderactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (hits,
// referencecount, undefhits, builtin, feature). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares,
// which is why it cannot reuse the resource model.
type ResponderactionDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Bypasssafetycheck  types.String `tfsdk:"bypasssafetycheck"`
	Comment            types.String `tfsdk:"comment"`
	Headers            types.List   `tfsdk:"headers"`
	Htmlpage           types.String `tfsdk:"htmlpage"`
	Name               types.String `tfsdk:"name"` // Required lookup key
	Newname            types.String `tfsdk:"newname"`
	Reasonphrase       types.String `tfsdk:"reasonphrase"`
	Responsestatuscode types.Int64  `tfsdk:"responsestatuscode"`
	Target             types.String `tfsdk:"target"`
	Type               types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/responderaction.json). Never settable; populated from GET.
	Hits           types.Int64  `tfsdk:"hits"`
	Referencecount types.Int64  `tfsdk:"referencecount"`
	Undefhits      types.Int64  `tfsdk:"undefhits"`
	Builtin        types.List   `tfsdk:"builtin"`
	Feature        types.String `tfsdk:"feature"`
}

func ResponderactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bypasssafetycheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bypass the safety check, allowing potentially unsafe expressions. An unsafe expression in a response is one that contains references to request elements that might not be present in all requests. If a response refers to a missing request element, an empty string is used instead.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Any type of information about this responder action.",
			},
			"headers": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more headers to insert into the HTTP response. Each header is specified as \"name(expr)\", where expr is an expression that is evaluated at runtime to provide the value for the named header. You can configure a maximum of eight headers for a responder action.",
			},
			"htmlpage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "For respondwithhtmlpage policies, name of the HTML page object to use as the response. You must first import the page object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the responder action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the responder policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder action\" or 'my responder action').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the responder action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder action\" or my responder action').",
			},
			"reasonphrase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the reason phrase of the HTTP response. The reason phrase may be a string literal with quotes or a PI expression. For example: \"Invalid URL: \" + HTTP.REQ.URL",
			},
			"responsestatuscode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP response status code, for example 200, 302, 404, etc. The default value for the redirect action type is 302 and for respondwithhtmlpage is 200",
			},
			"target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying what to respond with. Typically a URL for redirect policies or a default-syntax expression.  In addition to Citrix ADC default-syntax expressions that refer to information in the request, a stringbuilder expression can contain text and HTML, and simple escape codes that define new lines and paragraphs. Enclose each stringbuilder expression element (either a Citrix ADC default-syntax expression or a string) in double quotation marks. Use the plus (+) character to join the elements.\n\nExamples:\n1) Respondwith expression that sends an HTTP 1.1 200 OK response:\n\"HTTP/1.1 200 OK\\r\\n\\r\\n\"\n\n2) Redirect expression that redirects user to the specified web host and appends the request URL to the redirect.\n\"http://backupsite2.com\" + HTTP.REQ.URL\n\n3) Respondwith expression that sends an HTTP 1.1 404 Not Found response with the request URL included in the response:\n\"HTTP/1.1 404 Not Found\\r\\n\\r\\n\"+ \"HTTP.REQ.URL.HTTP_URL_SAFE\" + \"does not exist on the web server.\"\n\nThe following requirement applies only to the Citrix ADC CLI:\nEnclose the entire expression in single quotation marks. (Citrix ADC expression elements should be included inside the single quotation marks for the entire expression, but do not need to be enclosed in double quotation marks.)",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of responder action. Available settings function as follows:\n* respondwith <target> - Respond to the request with the expression specified as the target.\n* respondwithhtmlpage - Respond to the request with the uploaded HTML page object specified as the target.\n* redirect - Redirect the request to the URL specified as the target.\n* sqlresponse_ok - Send an SQL OK response.\n* sqlresponse_error - Send an SQL ERROR response.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action has been taken.",
			},
			"referencecount": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of references to the action.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action resulted in UNDEF.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether responder action is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// responderactionDataSourceSetAttrFromGet projects a NITRO responderaction GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func responderactionDataSourceSetAttrFromGet(ctx context.Context, data *ResponderactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In responderactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Bypasssafetycheck = utils.MapGetString(g, "bypasssafetycheck")
	data.Comment = utils.MapGetString(g, "comment")
	data.Headers = utils.MapGetStringList(g, "headers")
	data.Htmlpage = utils.MapGetString(g, "htmlpage")
	data.Reasonphrase = utils.MapGetString(g, "reasonphrase")
	data.Responsestatuscode = utils.MapGetInt64(g, "responsestatuscode")
	data.Target = utils.MapGetString(g, "target")
	data.Type = utils.MapGetString(g, "type")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Referencecount = utils.MapGetInt64(g, "referencecount")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
