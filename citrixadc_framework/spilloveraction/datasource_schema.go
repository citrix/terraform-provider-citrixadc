package spilloveraction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SpilloveractionDataSourceModel is the data-source-specific model, decoupled
// from SpilloveractionResourceModel. A data source is a pure read surface, so it
// exposes the read/write attributes (as Computed outputs) AND the read-only
// attributes the GET returns (builtin, feature) that the resource deliberately
// omits.
type SpilloveractionDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Action  types.String `tfsdk:"action"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/spilloveraction.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func SpilloveractionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Spillover action. Currently only type SPILLOVER is supported",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the spillover action.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the spillover action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters. \nChoose a name that can be correlated with the function that the action performs. \n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether the spillover action is default or not. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ].",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this configuration.",
			},
		},
	}
}

// spilloveractionDataSourceSetAttrFromGet projects a NITRO spilloveraction GET
// response onto the data-source model. Attributes are filled from the GET (or
// left Null when the GET omits them) via the shared utils.MapGet* helpers.
func spilloveractionDataSourceSetAttrFromGet(ctx context.Context, data *SpilloveractionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In spilloveractionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}
	data.Action = utils.MapGetString(g, "action")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
