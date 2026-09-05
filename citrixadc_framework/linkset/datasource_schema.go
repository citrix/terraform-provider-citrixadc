package linkset

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LinksetDataSourceModel is the data-source-specific model, decoupled from
// LinksetResourceModel. A data source is a pure read surface, so it exposes the
// existing datasource attributes as Computed outputs PLUS the read-only
// attributes the NITRO GET returns (zion73x_readonly/linkset.json) that the
// resource intentionally omits.
type LinksetDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Linksetid        types.String `tfsdk:"linkset_id"` // Required lookup key
	Interfacebinding types.Set    `tfsdk:"interfacebinding"`

	// Read-only (GET-only) attribute from the NITRO read-only set.
	Ifnum types.String `tfsdk:"ifnum"`
}

func LinksetDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"linkset_id": schema.StringAttribute{
				Required:    true,
				Description: "Unique identifier for the linkset. Must be of the form LS/x, where x can be an integer from 1 to 32.",
			},
			"interfacebinding": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Set of interface bindings for the linkset.",
			},

			// Read-only (GET-only) attribute surfaced by the data source.
			"ifnum": schema.StringAttribute{
				Computed:    true,
				Description: "The interfaces bound to the linkset, as returned by the appliance.",
			},
		},
	}
}

// linksetDataSourceSetAttrFromGet projects a NITRO linkset GET response onto the
// data-source model. The interfacebinding set is populated separately by the
// Read function via a binding query, so it is left Null here.
func linksetDataSourceSetAttrFromGet(ctx context.Context, data *LinksetDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In linksetDataSourceSetAttrFromGet Function")

	if v, ok := g["id"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Linksetid = types.StringValue(utils.AnyToString(v))
	}

	// Read-only attribute from the GET.
	data.Ifnum = utils.MapGetString(g, "ifnum")

	// interfacebinding is filled by the Read path from a separate binding query.
	data.Interfacebinding = types.SetNull(types.StringType)
}
