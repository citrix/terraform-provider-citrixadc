package videooptimizationpacingpolicylabel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VideooptimizationpacingpolicylabelDataSourceModel is the data-source-specific
// model, decoupled from VideooptimizationpacingpolicylabelResourceModel. A data
// source is a pure read surface, so it exposes the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only counters/metadata
// the resource deliberately omits (numpol, hits, priority,
// gotopriorityexpression, labeltype, invoke_labelname). Every non-key attribute
// is Computed.
type VideooptimizationpacingpolicylabelDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Comment         types.String `tfsdk:"comment"`
	Labelname       types.String `tfsdk:"labelname"`
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`

	// Read-only (GET-only) attributes from zion73x_readonly. Never settable.
	Numpol                 types.Int64  `tfsdk:"numpol"`
	Hits                   types.Int64  `tfsdk:"hits"`
	Priority               types.Int64  `tfsdk:"priority"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Labeltype              types.String `tfsdk:"labeltype"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
}

func VideooptimizationpacingpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this videooptimization pacing policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Video optimization pacing policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (\n.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the videooptimization pacing policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my videooptimization pacing policy label\" or my videooptimization pacing policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the videooptimization pacing policy label (rename-only field).",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of responses sent by the policies bound to this policy label. Types are:\n* HTTP - HTTP responses.\n* OTHERTCP - NON-HTTP TCP responses.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of policies bound to label.",
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
				Description: "Type of policy label to invoke. Possible values = vserver, policylabel.",
			},
			"invoke_labelname": schema.StringAttribute{
				Computed:    true,
				Description: "If labelType is policylabel, name of the policy label to invoke. If labelType is reqvserver or resvserver, name of the virtual server.",
			},
		},
	}
}

// videooptimizationpacingpolicylabelDataSourceSetAttrFromGet projects a NITRO
// videooptimizationpacingpolicylabel GET response onto the data-source model.
// Every attribute is filled from the GET (or left Null when the GET omits it) via
// the shared utils.MapGet* helpers.
func videooptimizationpacingpolicylabelDataSourceSetAttrFromGet(ctx context.Context, data *VideooptimizationpacingpolicylabelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In videooptimizationpacingpolicylabelDataSourceSetAttrFromGet Function")

	if v, ok := g["labelname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Labelname = types.StringValue(utils.AnyToString(v))
	}

	data.Comment = utils.MapGetString(g, "comment")
	data.Policylabeltype = utils.MapGetString(g, "policylabeltype")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.InvokeLabelname = utils.MapGetString(g, "invoke_labelname")
}
