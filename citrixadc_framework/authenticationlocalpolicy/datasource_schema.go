package authenticationlocalpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationlocalpolicyDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationlocalpolicyResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the read/write
// attributes (as Computed outputs) AND the read-only attributes the resource
// deliberately omits (reqaction). The Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type AuthenticationlocalpolicyDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Rule types.String `tfsdk:"rule"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/authenticationlocalpolicy.json). Never settable.
	Reqaction types.String `tfsdk:"reqaction"`
}

func AuthenticationlocalpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the local authentication policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after local policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to perform the authentication.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"reqaction": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the RADIUS action the policy uses.",
			},
		},
	}
}

// authenticationlocalpolicyDataSourceSetAttrFromGet projects a NITRO
// authenticationlocalpolicy GET response onto the data-source model. Attributes
// are simply filled from the GET (or left Null when the GET omits them) via the
// shared utils.MapGet* helpers.
func authenticationlocalpolicyDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationlocalpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationlocalpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Rule = utils.MapGetString(g, "rule")

	// Read-only metadata.
	data.Reqaction = utils.MapGetString(g, "reqaction")
}
