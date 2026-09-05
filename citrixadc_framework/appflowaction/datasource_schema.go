package appflowaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppflowactionDataSourceModel is the data-source-specific model, decoupled from
// AppflowactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (hits, referencecount, description). Every non-key attribute is Computed;
// the Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares, which is why it
// cannot reuse the resource model.
type AppflowactionDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Botinsight             types.String `tfsdk:"botinsight"`
	Ciinsight              types.String `tfsdk:"ciinsight"`
	Clientsidemeasurements types.String `tfsdk:"clientsidemeasurements"`
	Collectors             types.List   `tfsdk:"collectors"`
	Comment                types.String `tfsdk:"comment"`
	Distributionalgorithm  types.String `tfsdk:"distributionalgorithm"`
	Metricslog             types.Bool   `tfsdk:"metricslog"`
	Name                   types.String `tfsdk:"name"`
	Newname                types.String `tfsdk:"newname"`
	Pagetracking           types.String `tfsdk:"pagetracking"`
	Securityinsight        types.String `tfsdk:"securityinsight"`
	Transactionlog         types.String `tfsdk:"transactionlog"`
	Videoanalytics         types.String `tfsdk:"videoanalytics"`
	Webinsight             types.String `tfsdk:"webinsight"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/appflowaction.json). Never settable; populated from GET.
	Hits           types.Int64  `tfsdk:"hits"`
	Referencecount types.Int64  `tfsdk:"referencecount"`
	Description    types.String `tfsdk:"description"`
}

func AppflowactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"botinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the bot insight records to the configured collectors.",
			},
			"ciinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the ContentInspection Insight records to the configured collectors.",
			},
			"clientsidemeasurements": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will collect the time required to load and render the mainpage on the client.",
			},
			"collectors": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name(s) of collector(s) to be associated with the AppFlow action.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this action.  In the CLI, if including spaces between words, enclose the comment in quotation marks. (The quotation marks are not required in the configuration utility.)",
			},
			"distributionalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will distribute records among the collectors. Else, all records will be sent to all the collectors.",
			},
			"metricslog": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If only the stats records are to be exported, turn on this option.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow action\" or 'my appflow action').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the AppFlow action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow action\" or 'my appflow action').",
			},
			"pagetracking": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will start tracking the page for waterfall chart by inserting a NS_ESNS cookie in the response.",
			},
			"securityinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the security insight records to the configured collectors.",
			},
			"transactionlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log ANOMALOUS or ALL transactions",
			},
			"videoanalytics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the videoinsight records to the configured collectors.",
			},
			"webinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the webinsight records to the configured collectors.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action has been taken.",
			},
			"referencecount": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of references to the action.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the action.",
			},
		},
	}
}

// appflowactionDataSourceSetAttrFromGet projects a NITRO appflowaction GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func appflowactionDataSourceSetAttrFromGet(ctx context.Context, data *AppflowactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appflowactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Botinsight = utils.MapGetString(g, "botinsight")
	data.Ciinsight = utils.MapGetString(g, "ciinsight")
	data.Clientsidemeasurements = utils.MapGetString(g, "clientsidemeasurements")
	data.Collectors = utils.MapGetStringList(g, "collectors")
	data.Comment = utils.MapGetString(g, "comment")
	data.Distributionalgorithm = utils.MapGetString(g, "distributionalgorithm")
	data.Metricslog = utils.MapGetBool(g, "metricslog")
	data.Pagetracking = utils.MapGetString(g, "pagetracking")
	data.Securityinsight = utils.MapGetString(g, "securityinsight")
	data.Transactionlog = utils.MapGetString(g, "transactionlog")
	data.Videoanalytics = utils.MapGetString(g, "videoanalytics")
	data.Webinsight = utils.MapGetString(g, "webinsight")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Referencecount = utils.MapGetInt64(g, "referencecount")
	data.Description = utils.MapGetString(g, "description")
}
