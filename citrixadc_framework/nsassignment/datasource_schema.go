package nsassignment

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsassignmentDataSourceModel is the data-source-specific model, decoupled from
// NsassignmentResourceModel. A data source is a pure read surface, so it can
// expose the read/write attributes (as Computed outputs) AND the read-only
// attributes the resource deliberately omits (hits, undefhits, referencecount).
type NsassignmentDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Add      types.String `tfsdk:"add"`
	Append   types.String `tfsdk:"append"`
	Clear    types.Bool   `tfsdk:"clear"`
	Comment  types.String `tfsdk:"comment"`
	Name     types.String `tfsdk:"name"`
	Set      types.String `tfsdk:"set"`
	Sub      types.String `tfsdk:"sub"`
	Variable types.String `tfsdk:"variable"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nsassignment.json). Never settable; populated from GET.
	Hits           types.Int64 `tfsdk:"hits"`
	Undefhits      types.Int64 `tfsdk:"undefhits"`
	Referencecount types.Int64 `tfsdk:"referencecount"`
}

func NsassignmentDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"add": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Right hand side of the assignment. The expression is evaluated and added to the left hand variable.",
			},
			"append": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Right hand side of the assignment. The expression is evaluated and appended to the left hand variable.",
			},
			"clear": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Clear the variable value. Deallocates a text value, and for a map, the text key.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Can be used to preserve information about this rewrite action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the assignment. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the assignment is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my assignment\" or my assignment).",
			},
			"set": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Right hand side of the assignment. The expression is evaluated and assigned to the left hand variable.",
			},
			"sub": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Right hand side of the assignment. The expression is evaluated and subtracted from the left hand variable.",
			},
			"variable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Left hand side of the assigment, of the form $variable-name (for a singleton variabled) or $variable-name[key-expression], where key-expression is an expression that evaluates to a text string and provides the key to select a map entry",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
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
		},
	}
}

// nsassignmentDataSourceSetAttrFromGet projects a NITRO nsassignment GET response
// onto the data-source model. Attributes are filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func nsassignmentDataSourceSetAttrFromGet(ctx context.Context, data *NsassignmentDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsassignmentDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// NITRO returns the "add" field using the JSON key "Add" (capitalised in the
	// nitro-go struct tag); fall back to the lowercase form for robustness.
	if v, ok := g["Add"]; ok && v != nil {
		data.Add = types.StringValue(utils.AnyToString(v))
	} else {
		data.Add = utils.MapGetString(g, "add")
	}
	data.Append = utils.MapGetString(g, "append")
	data.Clear = utils.MapGetBool(g, "clear")
	data.Comment = utils.MapGetString(g, "comment")
	data.Set = utils.MapGetString(g, "set")
	data.Sub = utils.MapGetString(g, "sub")
	data.Variable = utils.MapGetString(g, "variable")

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Referencecount = utils.MapGetInt64(g, "referencecount")
}
