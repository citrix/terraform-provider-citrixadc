package vpnurlpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnurlpolicyDataSourceModel is the data-source-specific model, decoupled from
// VpnurlpolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (builtin, undefhits, ...). Every non-key attribute is Computed.
type VpnurlpolicyDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Action    types.String `tfsdk:"action"`
	Comment   types.String `tfsdk:"comment"`
	Logaction types.String `tfsdk:"logaction"`
	Name      types.String `tfsdk:"name"` // Required lookup key
	Newname   types.String `tfsdk:"newname"`
	Rule      types.String `tfsdk:"rule"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/vpnurlpolicy.json). Never settable; populated from GET.
	Builtin   types.List  `tfsdk:"builtin"`
	Undefhits types.Int64 `tfsdk:"undefhits"`
}

func VpnurlpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be applied by the new urlPolicy if the rule criteria are met.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of messagelog action to use when a request matches this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new urlPolicy.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the vpn urlPolicy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vpnurl policy\" or 'my vpnurl policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, specifying the traffic that matches the policy.\n\nThe following requirements apply only to the NetScaler CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the policy evaluation resulted in undefined processing.",
			},
		},
	}
}

// vpnurlpolicyDataSourceSetAttrFromGet projects a NITRO vpnurlpolicy GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func vpnurlpolicyDataSourceSetAttrFromGet(ctx context.Context, data *VpnurlpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnurlpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Comment = utils.MapGetString(g, "comment")
	data.Logaction = utils.MapGetString(g, "logaction")
	data.Rule = utils.MapGetString(g, "rule")

	// newname is a rename-only action input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
}
