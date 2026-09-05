package tunneltrafficpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TunneltrafficpolicyDataSourceModel is the data-source-specific model,
// decoupled from TunneltrafficpolicyResourceModel. A data source is a pure read
// surface (Read only; no plan/apply lifecycle), so it can expose the FULL GET
// projection: the read/write attributes (as Computed outputs) AND the read-only
// attributes the resource deliberately omits (hits, counters, isdefault,
// builtin, ...). Every non-key attribute is Computed.
type TunneltrafficpolicyDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Action    types.String `tfsdk:"action"`
	Comment   types.String `tfsdk:"comment"`
	Logaction types.String `tfsdk:"logaction"`
	Rule      types.String `tfsdk:"rule"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/tunneltrafficpolicy.json). Never settable; populated from GET.
	Expressiontype     types.String `tfsdk:"expressiontype"`
	Hits               types.Int64  `tfsdk:"hits"`
	Undefhits          types.Int64  `tfsdk:"undefhits"`
	Txbytes            types.Int64  `tfsdk:"txbytes"`
	Rxbytes            types.Int64  `tfsdk:"rxbytes"`
	Clientttlb         types.Int64  `tfsdk:"clientttlb"`
	Clienttransactions types.Int64  `tfsdk:"clienttransactions"`
	Serverttlb         types.Int64  `tfsdk:"serverttlb"`
	Servertransactions types.Int64  `tfsdk:"servertransactions"`
	Isdefault          types.Bool   `tfsdk:"isdefault"`
	Builtin            types.List   `tfsdk:"builtin"`
	Feature            types.String `tfsdk:"feature"`
}

func TunneltrafficpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the built-in compression action to associate with the policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the tunnel traffic policy.\nMust begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy)'.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, against which traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n*  If the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.\n*  If the expression itself includes double quotation marks, you must escape the quotations by using the \\ character. \n*  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"expressiontype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of policy (Classic/Advanced).",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of policy UNDEF hits.",
			},
			"txbytes": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of bytes transmitted.",
			},
			"rxbytes": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of bytes received.",
			},
			"clientttlb": schema.Int64Attribute{
				Computed:    true,
				Description: "Total client TTLB value.",
			},
			"clienttransactions": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of client transactions.",
			},
			"serverttlb": schema.Int64Attribute{
				Computed:    true,
				Description: "Total server TTLB value.",
			},
			"servertransactions": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of server transactions.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default tunnelpolicy.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// tunneltrafficpolicyDataSourceSetAttrFromGet projects a NITRO tunneltrafficpolicy
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func tunneltrafficpolicyDataSourceSetAttrFromGet(ctx context.Context, data *TunneltrafficpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In tunneltrafficpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Comment = utils.MapGetString(g, "comment")
	data.Logaction = utils.MapGetString(g, "logaction")
	data.Rule = utils.MapGetString(g, "rule")

	// Read-only metadata.
	data.Expressiontype = utils.MapGetString(g, "expressiontype")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Txbytes = utils.MapGetInt64(g, "txbytes")
	data.Rxbytes = utils.MapGetInt64(g, "rxbytes")
	data.Clientttlb = utils.MapGetInt64(g, "clientttlb")
	data.Clienttransactions = utils.MapGetInt64(g, "clienttransactions")
	data.Serverttlb = utils.MapGetInt64(g, "serverttlb")
	data.Servertransactions = utils.MapGetInt64(g, "servertransactions")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
