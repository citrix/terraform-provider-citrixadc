package cachepolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CachepolicyDataSourceModel is the data-source-specific model, decoupled from
// CachepolicyResourceModel. A data source is a pure read surface, so it can
// expose the full GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type CachepolicyDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Action       types.String `tfsdk:"action"`
	Invalgroups  types.List   `tfsdk:"invalgroups"`
	Invalobjects types.List   `tfsdk:"invalobjects"`
	Newname      types.String `tfsdk:"newname"`
	Policyname   types.String `tfsdk:"policyname"`
	Rule         types.String `tfsdk:"rule"`
	Storeingroup types.String `tfsdk:"storeingroup"`
	Undefaction  types.String `tfsdk:"undefaction"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/cachepolicy.json). Never settable; from GET.
	Hits      types.Int64  `tfsdk:"hits"`
	Undefhits types.Int64  `tfsdk:"undefhits"`
	Flags     types.Int64  `tfsdk:"flags"`
	Builtin   types.List   `tfsdk:"builtin"`
	Feature   types.String `tfsdk:"feature"`
}

func CachepolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to apply to content that matches the policy.\n* CACHE or MAY_CACHE action - positive cachability policy\n* NOCACHE or MAY_NOCACHE action - negative cachability policy\n* INVAL action - Dynamic Invalidation Policy",
			},
			"invalgroups": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Content group(s) to be invalidated when the INVAL action is applied. Maximum number of content groups that can be specified is 16.",
			},
			"invalobjects": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Content groups(s) in which the objects will be invalidated if the action is INVAL.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the cache policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the policy is created.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression against which the traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"storeingroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the content group in which to store the object when the final result of policy evaluation is CACHE. The content group must exist before being mentioned here. Use the \"show cache contentgroup\" command to view the list of existing content groups.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be performed when the result of rule evaluation is undefined.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of Undef hits.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flag.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether the cache policy is built-in. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// cachepolicyDataSourceSetAttrFromGet projects a NITRO cachepolicy GET response
// onto the data-source model. Attributes are filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func cachepolicyDataSourceSetAttrFromGet(ctx context.Context, data *CachepolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cachepolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["policyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Policyname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Invalgroups = utils.MapGetStringList(g, "invalgroups")
	data.Invalobjects = utils.MapGetStringList(g, "invalobjects")
	data.Rule = utils.MapGetString(g, "rule")
	data.Storeingroup = utils.MapGetString(g, "storeingroup")
	data.Undefaction = utils.MapGetString(g, "undefaction")

	// newname is an action-only (rename) input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
