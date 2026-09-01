package responderpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ResponderpolicyDataSourceModel is the data-source-specific model, decoupled
// from ResponderpolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (hits, undefhits, builtin, feature). Every non-key attribute is Computed; the
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares, which is why it
// cannot reuse the resource model.
type ResponderpolicyDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Action        types.String `tfsdk:"action"`
	Appflowaction types.String `tfsdk:"appflowaction"`
	Comment       types.String `tfsdk:"comment"`
	Logaction     types.String `tfsdk:"logaction"`
	Rule          types.String `tfsdk:"rule"`
	Undefaction   types.String `tfsdk:"undefaction"`

	// Convenience-block sets, shared with the resource model.
	Globalbinding    types.Set `tfsdk:"globalbinding"`
	Lbvserverbinding types.Set `tfsdk:"lbvserverbinding"`
	Csvserverbinding types.Set `tfsdk:"csvserverbinding"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/responderpolicy.json). Never settable; from GET.
	Hits      types.Int64  `tfsdk:"hits"`
	Undefhits types.Int64  `tfsdk:"undefhits"`
	Builtin   types.List   `tfsdk:"builtin"`
	Feature   types.String `tfsdk:"feature"`
}

// Shared computed-only attribute definitions for the convenience-block nested
// objects (identical to the previous inline "Computed: true" declarations). They
// are referenced rather than written inline so the block bodies stay concise.
var (
	responderpolicyDSComputedString = schema.StringAttribute{Computed: true}
	responderpolicyDSComputedBool   = schema.BoolAttribute{Computed: true}
	responderpolicyDSComputedInt64  = schema.Int64Attribute{Computed: true}
)

func ResponderpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the responder action to perform if the request matches this responder policy.",
			},
			"appflowaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AppFlow action to invoke for requests that match this policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this responder policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the responder policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that the policy uses to determine whether to respond to the specified request.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF).",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of policy UNDEF hits.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if responder policy is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
		// The convenience-block sets are shared with the resource model; they are
		// exposed here as computed outputs so the shared model maps cleanly.
		Blocks: map[string]schema.Block{
			"globalbinding": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"gotopriorityexpression": responderpolicyDSComputedString,
						"invoke":                 responderpolicyDSComputedBool,
						"labelname":              responderpolicyDSComputedString,
						"labeltype":              responderpolicyDSComputedString,
						"policyname":             responderpolicyDSComputedString,
						"priority":               responderpolicyDSComputedInt64,
						"type":                   responderpolicyDSComputedString,
					},
				},
			},
			"lbvserverbinding": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"bindpoint":              responderpolicyDSComputedString,
						"gotopriorityexpression": responderpolicyDSComputedString,
						"invoke":                 responderpolicyDSComputedBool,
						"labelname":              responderpolicyDSComputedString,
						"labeltype":              responderpolicyDSComputedString,
						"name":                   responderpolicyDSComputedString,
						"priority":               responderpolicyDSComputedInt64,
					},
				},
			},
			"csvserverbinding": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"bindpoint":              responderpolicyDSComputedString,
						"gotopriorityexpression": responderpolicyDSComputedString,
						"invoke":                 responderpolicyDSComputedBool,
						"labelname":              responderpolicyDSComputedString,
						"labeltype":              responderpolicyDSComputedString,
						"name":                   responderpolicyDSComputedString,
						"policyname":             responderpolicyDSComputedString,
						"priority":               responderpolicyDSComputedInt64,
						"targetlbvserver":        responderpolicyDSComputedString,
					},
				},
			},
		},
	}
}

// responderpolicyDataSourceSetAttrFromGet projects a NITRO responderpolicy GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func responderpolicyDataSourceSetAttrFromGet(ctx context.Context, data *ResponderpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In responderpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Appflowaction = utils.MapGetString(g, "appflowaction")
	data.Comment = utils.MapGetString(g, "comment")
	data.Logaction = utils.MapGetString(g, "logaction")
	data.Rule = utils.MapGetString(g, "rule")
	data.Undefaction = utils.MapGetString(g, "undefaction")

	// The convenience-block sets are not part of the responderpolicy GET
	// projection (they are reconciled from separate binding endpoints on the
	// resource); leave them Null on the data source.
	data.Globalbinding = types.SetNull(types.ObjectType{AttrTypes: responderpolicyGlobalbindingAttrTypes})
	data.Lbvserverbinding = types.SetNull(types.ObjectType{AttrTypes: responderpolicyLbvserverbindingAttrTypes})
	data.Csvserverbinding = types.SetNull(types.ObjectType{AttrTypes: responderpolicyCsvserverbindingAttrTypes})

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
