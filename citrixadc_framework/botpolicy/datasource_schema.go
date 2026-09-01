package botpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BotpolicyDataSourceModel is the data-source-specific model, decoupled from
// BotpolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (hits, undefhits, builtin, feature). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares,
// which is why it cannot reuse the resource model.
type BotpolicyDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"` // Required lookup key
	Newname     types.String `tfsdk:"newname"`
	Profilename types.String `tfsdk:"profilename"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/botpolicy.json). Never settable; populated from GET.
	Hits      types.Int64  `tfsdk:"hits"`
	Undefhits types.Int64  `tfsdk:"undefhits"`
	Builtin   types.List   `tfsdk:"builtin"`
	Feature   types.String `tfsdk:"feature"`
}

func BotpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this bot policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the bot policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the bot policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my bot policy\" or 'my bot policy').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the bot policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. \n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my bot policy\" or 'my bot policy').",
			},
			"profilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the bot profile to apply if the request matches this bot policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that the policy uses to determine whether to apply bot profile on the specified request.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
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
				Description: "Flag to determine if bot policy is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// botpolicyDataSourceSetAttrFromGet projects a NITRO botpolicy GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func botpolicyDataSourceSetAttrFromGet(ctx context.Context, data *BotpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In botpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Logaction = utils.MapGetString(g, "logaction")
	data.Profilename = utils.MapGetString(g, "profilename")
	data.Rule = utils.MapGetString(g, "rule")
	data.Undefaction = utils.MapGetString(g, "undefaction")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
