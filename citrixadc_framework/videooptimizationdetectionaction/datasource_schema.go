package videooptimizationdetectionaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VideooptimizationdetectionactionDataSourceModel is the data-source-specific
// model, decoupled from VideooptimizationdetectionactionResourceModel. A data
// source is a pure read surface (Read only; no plan/apply lifecycle), so it can
// expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (hits,
// referencecount, undefhits, builtin, ...). Every non-key attribute is Computed.
type VideooptimizationdetectionactionDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Comment types.String `tfsdk:"comment"`
	Newname types.String `tfsdk:"newname"`
	Type    types.String `tfsdk:"type"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/videooptimizationdetectionaction.json). Never settable; populated from GET.
	Hits           types.Int64  `tfsdk:"hits"`
	Referencecount types.Int64  `tfsdk:"referencecount"`
	Undefhits      types.Int64  `tfsdk:"undefhits"`
	Builtin        types.List   `tfsdk:"builtin"`
	Feature        types.String `tfsdk:"feature"`
}

func VideooptimizationdetectionactionDataSourceSchema() schema.Schema {
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
				Description: "Name for the video optimization detection action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the videooptimization detection action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of video optimization action. Available settings function as follows:\n* clear_text_pd - Cleartext PD type is detected.\n* clear_text_abr - Cleartext ABR is detected.\n* encrypted_abr - Encrypted ABR is detected.\n* trigger_enc_abr - Possible encrypted ABR is detected.\n* trigger_body_detection - Possible cleartext ABR is detected. Triggers body content detection.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
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
				Description: "Flag to determine whether video optimization detection action is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// videooptimizationdetectionactionDataSourceSetAttrFromGet projects a NITRO
// videooptimizationdetectionaction GET response onto the data-source model.
// Because a data source has no plan/apply reconciliation, attributes are simply
// filled from the GET (or left Null when the GET omits them). The shared
// utils.MapGet* helpers implement that projection.
func videooptimizationdetectionactionDataSourceSetAttrFromGet(ctx context.Context, data *VideooptimizationdetectionactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In videooptimizationdetectionactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Newname = utils.MapGetString(g, "newname")
	data.Type = utils.MapGetString(g, "type")

	// Read-only metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Referencecount = utils.MapGetInt64(g, "referencecount")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
