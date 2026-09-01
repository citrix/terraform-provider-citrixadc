package tunnelglobal_tunneltrafficpolicy_binding

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TunnelglobalTunneltrafficpolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from the resource model. A data source is a pure read surface
// (Read only; no plan/apply lifecycle), so it can expose the full GET projection:
// the read/write attributes (as Computed outputs) AND the read-only attributes
// that the resource deliberately omits (policytype, numpol).
type TunnelglobalTunneltrafficpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Feature                types.String `tfsdk:"feature"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	State                  types.String `tfsdk:"state"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/tunnelglobal_tunneltrafficpolicy_binding.json).
	Policytype types.String `tfsdk:"policytype"`
	Numpol     types.Int64  `tfsdk:"numpol"`
}

func TunnelglobalTunneltrafficpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"feature": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The feature to be checked while applying this config",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Policy name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Current state of the binding. If the binding is enabled, the policy is active.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind point to which the policy is bound.",
			},

			// Read-only (GET-only) metadata surfaced by the data source. All Computed.
			"policytype": schema.StringAttribute{
				Computed:    true,
				Description: "Policy type (Classic/Advanced) to be bound. Used for display. Possible values: [ Classic Policy, Advanced Policy ]",
			},
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of policies bound to the bindpoint.",
			},
		},
	}
}

// tunnelglobal_tunneltrafficpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// GET response onto the data-source model using the shared utils.MapGet* helpers.
func tunnelglobal_tunneltrafficpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *TunnelglobalTunneltrafficpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In tunnelglobal_tunneltrafficpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Feature = utils.MapGetString(g, "feature")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.State = utils.MapGetString(g, "state")

	// type is a GET filter / bind point not echoed by NITRO GET - preserve the
	// config-provided value the datasource was queried with when GET omits it.
	if t := utils.MapGetString(g, "type"); !t.IsNull() {
		data.Type = t
	}

	// Read-only (GET-only) metadata.
	data.Policytype = utils.MapGetString(g, "policytype")
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// Set ID for the datasource (plain policyname, matching the resource).
	data.Id = types.StringValue(data.Policyname.ValueString())
}
