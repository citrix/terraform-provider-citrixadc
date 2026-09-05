package policyexpression

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicyexpressionDataSourceModel is the data-source-specific model, decoupled
// from PolicyexpressionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (hits, pihits, type1, isdefault, builtin, feature). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares,
// which is why it cannot reuse the resource model.
type PolicyexpressionDataSourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"` // Required lookup key
	Clientsecuritymessage types.String `tfsdk:"clientsecuritymessage"`
	Comment               types.String `tfsdk:"comment"`
	Value                 types.String `tfsdk:"value"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/policyexpression.json). Never settable; populated from GET.
	Hits      types.Int64  `tfsdk:"hits"`
	Pihits    types.Int64  `tfsdk:"pihits"`
	Type1     types.String `tfsdk:"type1"`
	Isdefault types.Bool   `tfsdk:"isdefault"`
	Builtin   types.List   `tfsdk:"builtin"`
	Feature   types.String `tfsdk:"feature"`
}

func PolicyexpressionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientsecuritymessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Message to display if the expression fails. Allowed for classic end-point check expressions only.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the expression. Displayed upon viewing the policy expression.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Unique name for the expression. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or HTTP callout.",
			},
			"value": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression string. For example: http.req.body(100).contains(\"this\").",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of hits.",
			},
			"pihits": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of hits.",
			},
			"type1": schema.StringAttribute{
				Computed:    true,
				Description: "The type of expression. This is for output only. Possible values: CLASSIC, ADVANCED.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default policy expression.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// policyexpressionDataSourceSetAttrFromGet projects a NITRO policyexpression GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func policyexpressionDataSourceSetAttrFromGet(ctx context.Context, data *PolicyexpressionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In policyexpressionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Clientsecuritymessage = utils.MapGetString(g, "clientsecuritymessage")
	data.Comment = utils.MapGetString(g, "comment")
	data.Value = utils.MapGetString(g, "value")

	// Read-only metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Pihits = utils.MapGetInt64(g, "pihits")
	data.Type1 = utils.MapGetString(g, "type1")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
