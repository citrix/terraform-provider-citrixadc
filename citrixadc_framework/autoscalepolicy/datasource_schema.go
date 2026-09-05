package autoscalepolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AutoscalepolicyDataSourceModel is the data-source-specific model, decoupled
// from AutoscalepolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime attributes that the resource deliberately
// omits. Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type AutoscalepolicyDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Action    types.String `tfsdk:"action"`
	Comment   types.String `tfsdk:"comment"`
	Logaction types.String `tfsdk:"logaction"`
	Name      types.String `tfsdk:"name"`
	Newname   types.String `tfsdk:"newname"`
	Rule      types.String `tfsdk:"rule"`

	// Read-only (GET-only) runtime attributes from the NITRO doc read-only set
	// (zion73x_readonly/autoscalepolicy.json). Never settable; populated from GET.
	Hits      types.Int64 `tfsdk:"hits"`
	Undefhits types.Int64 `tfsdk:"undefhits"`
	Priority  types.Int64 `tfsdk:"priority"`
}

func AutoscalepolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The autoscale profile associated with the policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this autoscale policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The log action associated with the autoscale policy",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the autoscale policy.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the autoscale policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The rule associated with the policy.",
			},

			// Read-only (GET-only) runtime attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of Undef hits.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}

// autoscalepolicyDataSourceSetAttrFromGet projects a NITRO autoscalepolicy GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) — no unknown->null resolution or plan preservation is
// required. The shared utils.MapGet* helpers implement that projection.
func autoscalepolicyDataSourceSetAttrFromGet(ctx context.Context, data *AutoscalepolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In autoscalepolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Comment = utils.MapGetString(g, "comment")
	data.Logaction = utils.MapGetString(g, "logaction")
	data.Rule = utils.MapGetString(g, "rule")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only runtime attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Priority = utils.MapGetInt64(g, "priority")
}
