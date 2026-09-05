package csparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsparameterDataSourceModel is the data-source-specific model, decoupled from
// CsparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (builtin, feature). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type CsparameterDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Stateupdate types.String `tfsdk:"stateupdate"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/csparameter.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func CsparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"stateupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether the virtual server checks the attached load balancing server for state information.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if CS param is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// csparameterDataSourceSetAttrFromGet projects a NITRO csparameter GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func csparameterDataSourceSetAttrFromGet(ctx context.Context, data *CsparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In csparameterDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Stateupdate = utils.MapGetString(g, "stateupdate")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")

	// Singleton resource - no unique attributes - static ID.
	data.Id = types.StringValue("csparameter-config")
}
