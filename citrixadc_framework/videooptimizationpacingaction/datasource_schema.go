package videooptimizationpacingaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VideooptimizationpacingactionDataSourceModel is the data-source-specific model,
// decoupled from VideooptimizationpacingactionResourceModel. A data source is a
// pure read surface, so it can expose the FULL GET projection: the read/write
// attributes (as Computed outputs) AND the read-only counters/metadata the
// resource deliberately omits (hits, referencecount, undefhits, builtin,
// feature). Every non-key attribute is Computed.
type VideooptimizationpacingactionDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Rate    types.Int64  `tfsdk:"rate"`

	// Read-only (GET-only) attributes from zion73x_readonly. Never settable.
	Hits           types.Int64  `tfsdk:"hits"`
	Referencecount types.Int64  `tfsdk:"referencecount"`
	Undefhits      types.Int64  `tfsdk:"undefhits"`
	Builtin        types.List   `tfsdk:"builtin"`
	Feature        types.String `tfsdk:"feature"`
}

func VideooptimizationpacingactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Any type of information about this video optimization detection action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the video optimization pacing action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the videooptimization pacing action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"rate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ABR Video Optimization Pacing Rate (in Kbps)",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
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
				Description: "Flag to determine whether video optimization detection action is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// videooptimizationpacingactionDataSourceSetAttrFromGet projects a NITRO
// videooptimizationpacingaction GET response onto the data-source model. Every
// attribute is filled from the GET (or left Null when the GET omits it) via the
// shared utils.MapGet* helpers.
func videooptimizationpacingactionDataSourceSetAttrFromGet(ctx context.Context, data *VideooptimizationpacingactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In videooptimizationpacingactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Comment = utils.MapGetString(g, "comment")
	data.Rate = utils.MapGetInt64(g, "rate")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Referencecount = utils.MapGetInt64(g, "referencecount")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
