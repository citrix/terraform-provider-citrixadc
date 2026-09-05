package icaaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IcaactionDataSourceModel is the data-source-specific model, decoupled from
// IcaactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (hits, referencecount, undefhits, builtin, feature, isdefault). Every non-key
// attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type IcaactionDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"` // Required lookup key
	Accessprofilename  types.String `tfsdk:"accessprofilename"`
	Latencyprofilename types.String `tfsdk:"latencyprofilename"`
	Newname            types.String `tfsdk:"newname"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/icaaction.json). Never settable; populated from GET.
	Hits           types.Int64  `tfsdk:"hits"`
	Referencecount types.Int64  `tfsdk:"referencecount"`
	Undefhits      types.Int64  `tfsdk:"undefhits"`
	Builtin        types.List   `tfsdk:"builtin"`
	Feature        types.String `tfsdk:"feature"`
	Isdefault      types.Bool   `tfsdk:"isdefault"`
}

func IcaactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"accessprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ica accessprofile to be associated with this action.",
			},
			"latencyprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ica latencyprofile to be associated with this action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the ICA action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica action\" or 'my ica action').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the ICA action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#),period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks ( for example, \"my ica action\" or 'my ica action').",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action has been taken.",
			},
			"referencecount": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of references to the action.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action resulted in UNDEF.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that the ICA action is a built-in (SYSTEM INTERNAL) type. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default action.",
			},
		},
	}
}

// icaactionDataSourceSetAttrFromGet projects a NITRO icaaction GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func icaactionDataSourceSetAttrFromGet(ctx context.Context, data *IcaactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In icaactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Accessprofilename = utils.MapGetString(g, "accessprofilename")
	data.Latencyprofilename = utils.MapGetString(g, "latencyprofilename")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Referencecount = utils.MapGetInt64(g, "referencecount")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
}
