package lbpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbpolicyDataSourceModel is the data-source-specific model, decoupled from
// LbpolicyResourceModel. A data source is a pure read surface (Read only), so it
// exposes the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (hits,
// undefhits, feature, builtin).
type LbpolicyDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"` // Required lookup key
	Newname     types.String `tfsdk:"newname"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/lbpolicy.json). Never settable; populated from GET.
	Hits      types.Int64  `tfsdk:"hits"`
	Undefhits types.Int64  `tfsdk:"undefhits"`
	Feature   types.String `tfsdk:"feature"`
	Builtin   types.List   `tfsdk:"builtin"`
}

func LbpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of action to use if the request matches this LB policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this LB policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the LB policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the LB policy is added.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb policy\" or 'my lb policy').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the LB policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb policy\" or 'my lb policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression against which traffic is evaluated.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Available settings function as follows:\n* NOLBACTION - Does not consider LB actions in making LB decision.\n* RESET - Reset the request and notify the user, so that the user can resend the request.\n* DROP - Drop the request without sending a response to the user.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of policy UNDEF hits.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this configuration.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether the LB policy is built-in. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]. A list of strings.",
			},
		},
	}
}

// lbpolicyDataSourceSetAttrFromGet projects a NITRO lbpolicy GET response onto
// the data-source model. Attributes are simply filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func lbpolicyDataSourceSetAttrFromGet(ctx context.Context, data *LbpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	} else {
		data.Id = data.Name
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Comment = utils.MapGetString(g, "comment")
	data.Logaction = utils.MapGetString(g, "logaction")
	data.Rule = utils.MapGetString(g, "rule")
	data.Undefaction = utils.MapGetString(g, "undefaction")

	// newname is rename-only and never returned by GET -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Feature = utils.MapGetString(g, "feature")
	data.Builtin = utils.MapGetStringList(g, "builtin")
}
