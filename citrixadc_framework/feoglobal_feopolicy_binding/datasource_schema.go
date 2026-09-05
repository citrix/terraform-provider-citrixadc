package feoglobal_feopolicy_binding

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// FeoglobalFeopolicyBindingDataSourceModel is the data-source-specific model,
// decoupled from FeoglobalFeopolicyBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares.
type FeoglobalFeopolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/feoglobal_feopolicy_binding.json). Never settable;
	// populated from GET.
	Numpol types.Int64 `tfsdk:"numpol"`
}

func FeoglobalFeopolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"globalbindtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the globally bound front end optimization policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority assigned to the policy binding.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Bindpoint to which the policy is bound.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of policies bound to the bindpoint.",
			},
		},
	}
}

// feoglobal_feopolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// feoglobal_feopolicy_binding GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func feoglobal_feopolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *FeoglobalFeopolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In feoglobal_feopolicy_bindingDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only attributes.
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// ID is the plain policyname (backward-compatible with SDK v2).
	data.Id = types.StringValue(data.Policyname.ValueString())
}
