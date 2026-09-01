package videooptimizationdetectionpolicylabel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VideooptimizationdetectionpolicylabelDataSourceModel is the
// data-source-specific model, decoupled from
// VideooptimizationdetectionpolicylabelResourceModel. A data source is a pure
// read surface (Read only; no plan/apply lifecycle), so it can expose the FULL
// GET projection: the read/write attributes (as Computed outputs) AND the
// read-only attributes the resource deliberately omits (numpol, hits, priority,
// ...). Every non-key attribute is Computed.
type VideooptimizationdetectionpolicylabelDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Labelname types.String `tfsdk:"labelname"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Comment         types.String `tfsdk:"comment"`
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/videooptimizationdetectionpolicylabel.json). Never settable; populated from GET.
	Numpol                 types.Int64  `tfsdk:"numpol"`
	Hits                   types.Int64  `tfsdk:"hits"`
	Priority               types.Int64  `tfsdk:"priority"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Labeltype              types.String `tfsdk:"labeltype"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
}

func VideooptimizationdetectionpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this videooptimization detection policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Video optimization detection policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (\n.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the videooptimization detection policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my videooptimization detection policy label\" or my videooptimization detection policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the videooptimization detection policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (\n-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of responses sent by the policies bound to this policy label. Types are:\n* HTTP - HTTP responses.\n* OTHERTCP - NON-HTTP TCP responses.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of policies bound to the label.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times the policy label was invoked.",
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
				Description: "Type of policy label to invoke. Available settings function as follows:\n* vserver - Invoke an unnamed policy label associated with a virtual server.\n* policylabel - Invoke a user-defined policy label.",
			},
			"invoke_labelname": schema.StringAttribute{
				Computed:    true,
				Description: "* If labelType is policylabel, name of the policy label to invoke.\n* If labelType is reqvserver or resvserver, name of the virtual server.",
			},
		},
	}
}

// videooptimizationdetectionpolicylabelDataSourceSetAttrFromGet projects a NITRO
// videooptimizationdetectionpolicylabel GET response onto the data-source model.
// Because a data source has no plan/apply reconciliation, attributes are simply
// filled from the GET (or left Null when the GET omits them). The shared
// utils.MapGet* helpers implement that projection.
func videooptimizationdetectionpolicylabelDataSourceSetAttrFromGet(ctx context.Context, data *VideooptimizationdetectionpolicylabelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In videooptimizationdetectionpolicylabelDataSourceSetAttrFromGet Function")

	if v, ok := g["labelname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Labelname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Newname = utils.MapGetString(g, "newname")
	data.Policylabeltype = utils.MapGetString(g, "policylabeltype")

	// Read-only metadata.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.InvokeLabelname = utils.MapGetString(g, "invoke_labelname")
}
