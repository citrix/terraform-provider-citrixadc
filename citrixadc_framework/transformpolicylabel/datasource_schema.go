package transformpolicylabel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TransformpolicylabelDataSourceModel is the data-source-specific model,
// decoupled from TransformpolicylabelResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (numpol, hits, priority, gotopriorityexpression, labeltype, invoke_labelname,
// description, ...). Every non-key attribute is Computed.
type TransformpolicylabelDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Labelname       types.String `tfsdk:"labelname"` // Required lookup key
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/transformpolicylabel.json). Never settable; populated from GET.
	Numpol                 types.Int64  `tfsdk:"numpol"`
	Hits                   types.Int64  `tfsdk:"hits"`
	Priority               types.Int64  `tfsdk:"priority"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Labeltype              types.String `tfsdk:"labeltype"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
	Description            types.String `tfsdk:"description"`
}

func TransformpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the URL Transformation policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policylabel or my transform policylabel).",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the policy label.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policylabel or my transform policylabel).",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Types of transformations allowed by the policies bound to the label. For URL transformation, always http_req (HTTP Request).",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of polices bound to label.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times policy label was invoked.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of invocation. Available settings function as follows:\n* reqvserver - Forward the request to the specified request virtual server.\n* policylabel - Invoke the specified policy label. Possible values = reqvserver, policylabel",
			},
			"invoke_labelname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the policy label.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the policylabel.",
			},
		},
	}
}

// transformpolicylabelDataSourceSetAttrFromGet projects a NITRO
// transformpolicylabel GET response onto the data-source model. Because a data
// source has no plan/apply reconciliation, attributes are simply filled from the
// GET (or left Null when the GET omits them). The shared utils.MapGet* helpers
// implement that projection.
func transformpolicylabelDataSourceSetAttrFromGet(ctx context.Context, data *TransformpolicylabelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In transformpolicylabelDataSourceSetAttrFromGet Function")

	if v, ok := g["labelname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Labelname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Policylabeltype = utils.MapGetString(g, "policylabeltype")

	// newname is a rename-only input that the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.InvokeLabelname = utils.MapGetString(g, "invoke_labelname")
	data.Description = utils.MapGetString(g, "description")
}
