package dnspolicy64

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Dnspolicy64DataSourceModel is the data-source-specific model, decoupled from
// Dnspolicy64ResourceModel. A data source is a pure read surface, so it exposes
// the full GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes the resource deliberately omits (hits, labeltype,
// labelname, undefhits, description). Every non-key attribute is Computed.
type Dnspolicy64DataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"` // Required lookup key
	Rule   types.String `tfsdk:"rule"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnspolicy64.json). Never settable; populated from GET.
	Hits        types.Int64  `tfsdk:"hits"`
	Labeltype   types.String `tfsdk:"labeltype"`
	Labelname   types.String `tfsdk:"labelname"`
	Undefhits   types.Int64  `tfsdk:"undefhits"`
	Description types.String `tfsdk:"description"`
}

func Dnspolicy64DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS64 action to perform when the rule evaluates to TRUE. The built in actions function as follows:\n* A default dns64 action with prefix <default prefix> and mapped and exclude are any\nYou can create custom actions by using the add dns action command in the CLI or the DNS64 > Actions > Create DNS64 Action dialog box in the Citrix ADC configuration utility.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DNS64 policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression against which DNS traffic is evaluated.\nNote:\n* On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.\n* If the expression itself includes double quotation marks, you must escape the quotations by using the  character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.\nExample: CLIENT.IP.SRC.IN_SUBENT(23.34.0.0/16)",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the policy has been hit.",
			},
			"labeltype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of policy label invocation. Possible values: [ reqvserver, resvserver, policylabel ]",
			},
			"labelname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the label to invoke if the current policy rule evaluates to TRUE.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of Undef hits.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the policy.",
			},
		},
	}
}

// dnspolicy64DataSourceSetAttrFromGet projects a NITRO dnspolicy64 GET response
// onto the data-source model. Attributes are filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func dnspolicy64DataSourceSetAttrFromGet(ctx context.Context, data *Dnspolicy64DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnspolicy64DataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Rule = utils.MapGetString(g, "rule")

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Description = utils.MapGetString(g, "description")
}
