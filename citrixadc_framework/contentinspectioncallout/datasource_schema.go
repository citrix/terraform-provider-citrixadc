package contentinspectioncallout

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ContentinspectioncalloutDataSourceModel is the data-source-specific model,
// decoupled from ContentinspectioncalloutResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only counter/status attributes that the resource
// deliberately omits (hits, undefhits, undefreason). Every non-key attribute is
// Computed.
type ContentinspectioncalloutDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Comment     types.String `tfsdk:"comment"`
	Name        types.String `tfsdk:"name"` // Required lookup key
	Profilename types.String `tfsdk:"profilename"`
	Resultexpr  types.String `tfsdk:"resultexpr"`
	Returntype  types.String `tfsdk:"returntype"`
	Serverip    types.String `tfsdk:"serverip"`
	Servername  types.String `tfsdk:"servername"`
	Serverport  types.Int64  `tfsdk:"serverport"`
	Type        types.String `tfsdk:"type"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/contentinspectioncallout.json). Never settable.
	Hits        types.Int64  `tfsdk:"hits"`
	Undefhits   types.Int64  `tfsdk:"undefhits"`
	Undefreason types.String `tfsdk:"undefreason"`
}

func ContentinspectioncalloutDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this Content Inspection callout.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Content Inspection callout. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or callout.",
			},
			"profilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Content Inspection profile. The type of the configured profile must match the type specified using -type argument.",
			},
			"resultexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that extracts the callout results from the response sent by the CI callout agent. Must be a response based expression, that is, it must begin with ICAP.RES. The operations in this expression must match the return type. For example, if you configure a return type of TEXT, the result expression must be a text based expression, as in the following example: icap.res.header(\"ISTag\")",
			},
			"returntype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of data that the target callout agent returns in response to the callout.\nAvailable settings function as follows:\n* TEXT - Treat the returned value as a text string.\n* NUM - Treat the returned value as a number.\n* BOOL - Treat the returned value as a Boolean value.\nNote: You cannot change the return type after it is set.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of Content Inspection server. Mutually exclusive with the server name parameter.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing or content switching virtual server or service to which the Content Inspection request is issued. Mutually exclusive with server IP address and port parameters. The service type must be TCP or SSL_TCP. If there are vservers and services with the same name, then vserver is selected.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port of the Content Inspection server.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the Content Inspection callout. It must be one of the following:\n* ICAP - Sends ICAP request to the configured ICAP server.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Total hits.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Total undefs.",
			},
			"undefreason": schema.StringAttribute{
				Computed:    true,
				Description: "Reason for last undef.",
			},
		},
	}
}

// contentinspectioncalloutDataSourceSetAttrFromGet projects a NITRO
// contentinspectioncallout GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func contentinspectioncalloutDataSourceSetAttrFromGet(ctx context.Context, data *ContentinspectioncalloutDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In contentinspectioncalloutDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Profilename = utils.MapGetString(g, "profilename")
	data.Resultexpr = utils.MapGetString(g, "resultexpr")
	data.Returntype = utils.MapGetString(g, "returntype")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Servername = utils.MapGetString(g, "servername")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Type = utils.MapGetString(g, "type")

	// Read-only metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Undefreason = utils.MapGetString(g, "undefreason")
}
