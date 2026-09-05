package spilloverpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SpilloverpolicyDataSourceModel is the data-source-specific model, decoupled
// from SpilloverpolicyResourceModel. A data source is a pure read surface, so it
// exposes the read/write attributes (as Computed outputs) AND the read-only
// attributes the GET returns (hits, undefhits, builtin, feature) that the
// resource deliberately omits.
type SpilloverpolicyDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Action  types.String `tfsdk:"action"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Rule    types.String `tfsdk:"rule"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/spilloverpolicy.json). Never settable; populated from GET.
	Hits      types.Int64  `tfsdk:"hits"`
	Undefhits types.Int64  `tfsdk:"undefhits"`
	Builtin   types.List   `tfsdk:"builtin"`
	Feature   types.String `tfsdk:"feature"`
}

func SpilloverpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action for the spillover policy. Action is created using add spillover action command",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments that you might want to associate with the spillover policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the spillover policy.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the spillover policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nChoose a name that reflects the function that the policy performs. \n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression to be used by the spillover policy.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the policy has been hit.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of policy UNDEF hits.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if the spillover policy is builtin or not. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ].",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this configuration.",
			},
		},
	}
}

// spilloverpolicyDataSourceSetAttrFromGet projects a NITRO spilloverpolicy GET
// response onto the data-source model. Attributes are filled from the GET (or
// left Null when the GET omits them) via the shared utils.MapGet* helpers.
func spilloverpolicyDataSourceSetAttrFromGet(ctx context.Context, data *SpilloverpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In spilloverpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}
	data.Action = utils.MapGetString(g, "action")
	data.Comment = utils.MapGetString(g, "comment")
	data.Rule = utils.MapGetString(g, "rule")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
