package cmppolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CmppolicyDataSourceModel is the data-source-specific model, decoupled from
// CmppolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (reqaction, counters, description, builtin, feature, isdefault). Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type CmppolicyDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Newname   types.String `tfsdk:"newname"`
	Resaction types.String `tfsdk:"resaction"`
	Rule      types.String `tfsdk:"rule"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/cmppolicy.json). Never settable; populated from GET.
	Reqaction          types.String `tfsdk:"reqaction"`
	Hits               types.Int64  `tfsdk:"hits"`
	Txbytes            types.Int64  `tfsdk:"txbytes"`
	Rxbytes            types.Int64  `tfsdk:"rxbytes"`
	Clientttlb         types.Int64  `tfsdk:"clientttlb"`
	Clienttransactions types.Int64  `tfsdk:"clienttransactions"`
	Serverttlb         types.Int64  `tfsdk:"serverttlb"`
	Servertransactions types.Int64  `tfsdk:"servertransactions"`
	Description        types.String `tfsdk:"description"`
	Builtin            types.List   `tfsdk:"builtin"`
	Feature            types.String `tfsdk:"feature"`
	Isdefault          types.Bool   `tfsdk:"isdefault"`
}

func CmppolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the HTTP compression policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nCan be changed after the policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp policy\" or 'my cmp policy').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the compression policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nChoose a name that reflects the function that the policy performs.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp policy\" or 'my cmp policy').",
			},
			"resaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The built-in or user-defined compression action to apply to the response when the policy matches a request or response.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that determines which HTTP requests or responses match the compression policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"reqaction": schema.StringAttribute{
				Computed:    true,
				Description: "The compression action to be performed on requests.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"txbytes": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of bytes transferred.",
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
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the policy.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if compression policy is builtin or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default policy.",
			},
		},
	}
}

// cmppolicyDataSourceSetAttrFromGet projects a NITRO cmppolicy GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func cmppolicyDataSourceSetAttrFromGet(ctx context.Context, data *CmppolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cmppolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Resaction = utils.MapGetString(g, "resaction")
	data.Rule = utils.MapGetString(g, "rule")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.Reqaction = utils.MapGetString(g, "reqaction")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Txbytes = utils.MapGetInt64(g, "txbytes")
	data.Rxbytes = utils.MapGetInt64(g, "rxbytes")
	data.Clientttlb = utils.MapGetInt64(g, "clientttlb")
	data.Clienttransactions = utils.MapGetInt64(g, "clienttransactions")
	data.Serverttlb = utils.MapGetInt64(g, "serverttlb")
	data.Servertransactions = utils.MapGetInt64(g, "servertransactions")
	data.Description = utils.MapGetString(g, "description")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
}
