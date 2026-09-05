package appfwfieldtype

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwfieldtypeDataSourceModel is the data-source-specific model, decoupled
// from AppfwfieldtypeResourceModel. A data source is a pure read surface, so it
// exposes the existing lookup/config attributes (as Computed outputs) PLUS the
// read-only (GET-only) attributes the resource deliberately omits.
type AppfwfieldtypeDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Comment    types.String `tfsdk:"comment"`
	Name       types.String `tfsdk:"name"` // Required lookup key
	Nocharmaps types.Bool   `tfsdk:"nocharmaps"`
	Priority   types.Int64  `tfsdk:"priority"`
	Regex      types.String `tfsdk:"regex"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/appfwfieldtype.json). Populated from GET; never settable.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AppfwfieldtypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment describing the type of field that this field type is intended to match.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the field type.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the field type is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my field type\" or 'my field type').",
			},
			"nocharmaps": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "will not show internal field types added as part of FieldFormat learn rules deployment",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer specifying the priority of the field type. A lower number specifies a higher priority. Field types are checked in the order of their priority numbers.",
			},
			"regex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE - format regular expression defining the characters and length allowed for this field type.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if fieldtype is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// appfwfieldtypeDataSourceSetAttrFromGet projects a NITRO appfwfieldtype GET
// response onto the data-source model. Every attribute is filled from the GET
// (or left Null when the GET omits it); id mirrors the returned name key.
func appfwfieldtypeDataSourceSetAttrFromGet(ctx context.Context, data *AppfwfieldtypeDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwfieldtypeDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Comment = utils.MapGetString(g, "comment")
	data.Nocharmaps = utils.MapGetBool(g, "nocharmaps")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Regex = utils.MapGetString(g, "regex")

	// Read-only (GET-only) attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
