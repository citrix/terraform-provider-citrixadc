package tmtrafficpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TmtrafficpolicyDataSourceModel is the data-source-specific model, decoupled
// from TmtrafficpolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (hits, ...). Every non-key attribute is Computed.
type TmtrafficpolicyDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"` // Required lookup key
	Rule   types.String `tfsdk:"rule"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/tmtrafficpolicy.json). Never settable; populated from GET.
	Hits types.Int64 `tfsdk:"hits"`
}

func TmtrafficpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the action to apply to requests or connections that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the traffic policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named expression, or an expression, that the policy uses to determine whether to apply certain action on the current traffic.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
		},
	}
}

// tmtrafficpolicyDataSourceSetAttrFromGet projects a NITRO tmtrafficpolicy GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func tmtrafficpolicyDataSourceSetAttrFromGet(ctx context.Context, data *TmtrafficpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In tmtrafficpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Rule = utils.MapGetString(g, "rule")

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
}
