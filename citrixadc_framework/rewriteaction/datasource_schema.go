package rewriteaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RewriteactionDataSourceModel is the data-source-specific model, decoupled from
// RewriteactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (hits, undefhits, referencecount, description, isdefault, builtin, feature).
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type RewriteactionDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Comment           types.String `tfsdk:"comment"`
	Newname           types.String `tfsdk:"newname"`
	Refinesearch      types.String `tfsdk:"refinesearch"`
	Search            types.String `tfsdk:"search"`
	Stringbuilderexpr types.String `tfsdk:"stringbuilderexpr"`
	Target            types.String `tfsdk:"target"`
	Type              types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/rewriteaction.json). Never settable; from GET.
	Hits           types.Int64  `tfsdk:"hits"`
	Undefhits      types.Int64  `tfsdk:"undefhits"`
	Referencecount types.Int64  `tfsdk:"referencecount"`
	Description    types.String `tfsdk:"description"`
	Isdefault      types.Bool   `tfsdk:"isdefault"`
	Builtin        types.List   `tfsdk:"builtin"`
	Feature        types.String `tfsdk:"feature"`
}

func RewriteactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Can be used to preserve information about this rewrite action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the user-defined rewrite action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite action\" or 'my rewrite action').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the rewrite action. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite action\" or 'my rewrite action').",
			},
			"refinesearch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify additional criteria to refine the results of the search. \nAlways starts with the \"extend(m,n)\" operation, where 'm' specifies number of bytes to the left of selected data and 'n' specifies number of bytes to the right of selected data to extend the selected area.\nYou can use refineSearch only on body expressions, and for the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types.\nExample: -refineSearch 'EXTEND(10, 20).REGEX_SELECT(re~0x[0-9a-zA-Z]+~).",
			},
			"search": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Search facility that is used to match multiple strings in the request or response. Used in the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types.",
			},
			"stringbuilderexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that specifies the content to insert into the request or response at the specified location, or that replaces the specified string.",
			},
			"target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that specifies which part of the request or response to rewrite.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of user-defined rewrite action.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action has been taken.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action resulted in UNDEF.",
			},
			"referencecount": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of references to the action.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the action.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default rewriteaction.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether rewrite action is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// rewriteactionDataSourceSetAttrFromGet projects a NITRO rewriteaction GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func rewriteactionDataSourceSetAttrFromGet(ctx context.Context, data *RewriteactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In rewriteactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Refinesearch = utils.MapGetString(g, "refinesearch")
	data.Search = utils.MapGetString(g, "search")
	data.Stringbuilderexpr = utils.MapGetString(g, "stringbuilderexpr")
	data.Target = utils.MapGetString(g, "target")
	data.Type = utils.MapGetString(g, "type")

	// newname is a rename-only (?action=rename) input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Referencecount = utils.MapGetInt64(g, "referencecount")
	data.Description = utils.MapGetString(g, "description")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
