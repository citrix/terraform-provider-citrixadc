package rewritepolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RewritepolicyDataSourceModel is the data-source-specific model, decoupled from
// RewritepolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (hits, undefhits, description, isdefault, builtin, feature). Every non-key
// attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type RewritepolicyDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`

	// Convenience-block sets, shared with the resource model.
	Globalbinding    types.Set `tfsdk:"globalbinding"`
	Lbvserverbinding types.Set `tfsdk:"lbvserverbinding"`
	Csvserverbinding types.Set `tfsdk:"csvserverbinding"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/rewritepolicy.json). Never settable; from GET.
	Hits        types.Int64  `tfsdk:"hits"`
	Undefhits   types.Int64  `tfsdk:"undefhits"`
	Description types.String `tfsdk:"description"`
	Isdefault   types.Bool   `tfsdk:"isdefault"`
	Builtin     types.List   `tfsdk:"builtin"`
	Feature     types.String `tfsdk:"feature"`
}

// Shared computed-only attribute definitions for the convenience-block nested
// objects (identical to the previous inline "Computed: true" declarations). They
// are referenced rather than written inline so the block bodies stay concise.
var (
	rewritepolicyDSComputedString = schema.StringAttribute{Computed: true}
	rewritepolicyDSComputedBool   = schema.BoolAttribute{Computed: true}
	rewritepolicyDSComputedInt64  = schema.Int64Attribute{Computed: true}
)

func RewritepolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the rewrite action to perform if the request or response matches this rewrite policy.\nThere are also some built-in actions which can be used. These are:\n* NOREWRITE - Send the request from the client to the server or response from the server to the client without making any changes in the message.\n* RESET - Resets the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n* DROP - Drop the request without sending a response to the user.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this rewrite policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of messagelog action to use when a request matches this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the rewrite policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite policy\" or 'my rewrite policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression against which traffic is evaluated.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character. \n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of Undef hits.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the policy.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default rewritepolicy.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if rewrite policy is built-in or not.",
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
						"globalbindtype":         rewritepolicyDSComputedString,
						"gotopriorityexpression": rewritepolicyDSComputedString,
						"invoke":                 rewritepolicyDSComputedBool,
						"labelname":              rewritepolicyDSComputedString,
						"labeltype":              rewritepolicyDSComputedString,
						"policyname":             rewritepolicyDSComputedString,
						"priority":               rewritepolicyDSComputedInt64,
						"type":                   rewritepolicyDSComputedString,
					},
				},
			},
			"lbvserverbinding": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"bindpoint":              rewritepolicyDSComputedString,
						"gotopriorityexpression": rewritepolicyDSComputedString,
						"invoke":                 rewritepolicyDSComputedBool,
						"labelname":              rewritepolicyDSComputedString,
						"labeltype":              rewritepolicyDSComputedString,
						"name":                   rewritepolicyDSComputedString,
						"priority":               rewritepolicyDSComputedInt64,
					},
				},
			},
			"csvserverbinding": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"bindpoint":              rewritepolicyDSComputedString,
						"gotopriorityexpression": rewritepolicyDSComputedString,
						"invoke":                 rewritepolicyDSComputedBool,
						"labelname":              rewritepolicyDSComputedString,
						"labeltype":              rewritepolicyDSComputedString,
						"name":                   rewritepolicyDSComputedString,
						"priority":               rewritepolicyDSComputedInt64,
						"targetlbvserver":        rewritepolicyDSComputedString,
					},
				},
			},
		},
	}
}

// rewritepolicyDataSourceSetAttrFromGet projects a NITRO rewritepolicy GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func rewritepolicyDataSourceSetAttrFromGet(ctx context.Context, data *RewritepolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In rewritepolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Comment = utils.MapGetString(g, "comment")
	data.Logaction = utils.MapGetString(g, "logaction")
	data.Rule = utils.MapGetString(g, "rule")
	data.Undefaction = utils.MapGetString(g, "undefaction")

	// The convenience-block sets are not part of the rewritepolicy GET projection
	// (they are reconciled from separate binding endpoints on the resource); leave
	// them Null on the data source.
	data.Globalbinding = types.SetNull(types.ObjectType{AttrTypes: rewritepolicyGlobalbindingAttrTypes})
	data.Lbvserverbinding = types.SetNull(types.ObjectType{AttrTypes: rewritepolicyLbvserverbindingAttrTypes})
	data.Csvserverbinding = types.SetNull(types.ObjectType{AttrTypes: rewritepolicyCsvserverbindingAttrTypes})

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Description = utils.MapGetString(g, "description")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
