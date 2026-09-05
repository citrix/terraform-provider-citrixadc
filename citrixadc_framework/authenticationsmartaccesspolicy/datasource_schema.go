package authenticationsmartaccesspolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationsmartaccesspolicyDataSourceModel is the data-source-specific
// model, decoupled from AuthenticationsmartaccesspolicyResourceModel. Every
// non-key attribute is Computed, and it additionally exposes read-only
// (GET-only) attributes the resource deliberately omits (hits).
type AuthenticationsmartaccesspolicyDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Action  types.String `tfsdk:"action"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Rule    types.String `tfsdk:"rule"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/authenticationsmartaccesspolicy.json). Never settable;
	// populated from GET.
	Hits types.Int64 `tfsdk:"hits"`
}

func AuthenticationsmartaccesspolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Smartaccess profile to use if the policy matches.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Smartaccess policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after Smartaccess policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression.",
			},

			// Read-only (GET-only) attribute surfaced by the data source
			// (intentionally NOT modeled on the resource). Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
		},
	}
}

// authenticationsmartaccesspolicyDataSourceSetAttrFromGet projects a NITRO
// authenticationsmartaccesspolicy GET response onto the data-source model.
// Because a data source has no plan/apply reconciliation, attributes are simply
// filled from the GET (or left Null when the GET omits them). The shared
// utils.MapGet* helpers implement that projection.
func authenticationsmartaccesspolicyDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationsmartaccesspolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationsmartaccesspolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Action = utils.MapGetString(g, "action")
	data.Comment = utils.MapGetString(g, "comment")
	data.Rule = utils.MapGetString(g, "rule")

	// Read-only attribute.
	data.Hits = utils.MapGetInt64(g, "hits")
}
