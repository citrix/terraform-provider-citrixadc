package authorizationpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthorizationpolicyDataSourceModel is the data-source-specific model,
// decoupled from AuthorizationpolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime attributes that the resource deliberately
// omits. Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type AuthorizationpolicyDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Action  types.String `tfsdk:"action"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Rule    types.String `tfsdk:"rule"`

	// Read-only (GET-only) runtime attributes from the NITRO doc read-only set
	// (zion73x_readonly/authorizationpolicy.json). Never settable; populated from
	// GET.
	Activepolicy   types.Int64  `tfsdk:"activepolicy"`
	Expressiontype types.String `tfsdk:"expressiontype"`
	Hits           types.Int64  `tfsdk:"hits"`
}

func AuthorizationpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the policy matches: either allow or deny the request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new authorization policy. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authorization policy\" or 'my authorization policy').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the author policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to perform the authentication.",
			},

			// Read-only (GET-only) runtime attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"activepolicy": schema.Int64Attribute{
				Computed:    true,
				Description: "Indicates whether policy is bound or not.",
			},
			"expressiontype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of policy (Classic/Advanced).",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
		},
	}
}

// authorizationpolicyDataSourceSetAttrFromGet projects a NITRO authorizationpolicy
// GET response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when the
// GET omits them) — no unknown->null resolution or plan preservation is required.
// The shared utils.MapGet* helpers implement that projection.
func authorizationpolicyDataSourceSetAttrFromGet(ctx context.Context, data *AuthorizationpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authorizationpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Rule = utils.MapGetString(g, "rule")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only runtime attributes.
	data.Activepolicy = utils.MapGetInt64(g, "activepolicy")
	data.Expressiontype = utils.MapGetString(g, "expressiontype")
	data.Hits = utils.MapGetInt64(g, "hits")
}
