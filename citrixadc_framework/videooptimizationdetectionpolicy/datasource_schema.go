package videooptimizationdetectionpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VideooptimizationdetectionpolicyDataSourceModel is the data-source-specific
// model, decoupled from VideooptimizationdetectionpolicyResourceModel. A data
// source is a pure read surface (Read only; no plan/apply lifecycle), so it can
// expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (hits,
// undefhits, builtin, ...). Every non-key attribute is Computed.
type VideooptimizationdetectionpolicyDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Newname     types.String `tfsdk:"newname"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/videooptimizationdetectionpolicy.json). Never settable; populated from GET.
	Hits      types.Int64  `tfsdk:"hits"`
	Undefhits types.Int64  `tfsdk:"undefhits"`
	Builtin   types.List   `tfsdk:"builtin"`
	Feature   types.String `tfsdk:"feature"`
}

func VideooptimizationdetectionpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the videooptimization detection action to perform if the request matches this videooptimization detection policy. Built-in actions should be used. These are:\n* DETECT_CLEARTEXT_PD - Cleartext PD is detected and increment related counters.\n* DETECT_CLEARTEXT_ABR - Cleartext ABR is detected and increment related counters.\n* DETECT_ENCRYPTED_ABR - Encrypted ABR is detected and increment related counters.\n* TRIGGER_ENC_ABR_DETECTION - This is potentially encrypted ABR. Internal traffic heuristics algorithms will further process traffic to confirm detection.\n* TRIGGER_CT_ABR_BODY_DETECTION -  This is potentially cleartext ABR. Internal traffic heuristics algorithms will further process traffic to confirm detection.\n* RESET - Reset the client connection by closing it.\n* DROP - Drop the connection without sending a response.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this videooptimization detection policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the videooptimization detection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.Can be modified, removed or renamed.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the videooptimization detection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that determines which request or response match the video optimization detection policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.",
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
				Description: "Flag to determine if videooptimization detection policy is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// videooptimizationdetectionpolicyDataSourceSetAttrFromGet projects a NITRO
// videooptimizationdetectionpolicy GET response onto the data-source model.
// Because a data source has no plan/apply reconciliation, attributes are simply
// filled from the GET (or left Null when the GET omits them). The shared
// utils.MapGet* helpers implement that projection.
func videooptimizationdetectionpolicyDataSourceSetAttrFromGet(ctx context.Context, data *VideooptimizationdetectionpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In videooptimizationdetectionpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Comment = utils.MapGetString(g, "comment")
	data.Logaction = utils.MapGetString(g, "logaction")
	data.Newname = utils.MapGetString(g, "newname")
	data.Rule = utils.MapGetString(g, "rule")
	data.Undefaction = utils.MapGetString(g, "undefaction")

	// Read-only metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
