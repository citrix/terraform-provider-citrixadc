package appflowpolicylabel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppflowpolicylabelDataSourceModel is the data-source-specific model, decoupled
// from AppflowpolicylabelResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type AppflowpolicylabelDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Labelname       types.String `tfsdk:"labelname"` // Required lookup key
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appflowpolicylabel.json). Never settable; populated from GET.
	Numpol                 types.Int64  `tfsdk:"numpol"`
	Hits                   types.Int64  `tfsdk:"hits"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Labeltype              types.String `tfsdk:"labeltype"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
	Flowtype               types.Int64  `tfsdk:"flowtype"`
	Description            types.String `tfsdk:"description"`
}

func AppflowpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the AppFlow policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow policylabel\" or 'my appflow policylabel').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\n                    The following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow policylabel\" or 'my appflow policylabel')",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of traffic evaluated by the policies bound to the policy label.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of polices bound to the policy label.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times the policy label was invoked.",
			},
			"policyname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the AppFlow policy.",
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
				Description: "Type of policy label to be invoked. Possible values: [ vserver, policylabel ].",
			},
			"invoke_labelname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the label to invoke if the current policy evaluates to TRUE.",
			},
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "Flowtype of the bound AppFlow policy.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the policylabel.",
			},
		},
	}
}

// appflowpolicylabelDataSourceSetAttrFromGet projects a NITRO appflowpolicylabel
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func appflowpolicylabelDataSourceSetAttrFromGet(ctx context.Context, data *AppflowpolicylabelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appflowpolicylabelDataSourceSetAttrFromGet Function")

	if v, ok := g["labelname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Labelname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Policylabeltype = utils.MapGetString(g, "policylabeltype")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.InvokeLabelname = utils.MapGetString(g, "invoke_labelname")
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Description = utils.MapGetString(g, "description")
}
