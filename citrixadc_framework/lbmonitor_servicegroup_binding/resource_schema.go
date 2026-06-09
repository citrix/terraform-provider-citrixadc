package lbmonitor_servicegroup_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbmonitorServicegroupBindingResourceModel describes the resource data model.
type LbmonitorServicegroupBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	DupState         types.String `tfsdk:"dup_state"`
	DupWeight        types.Int64  `tfsdk:"dup_weight"`
	Monitorname      types.String `tfsdk:"monitorname"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Servicename      types.String `tfsdk:"servicename"`
	State            types.String `tfsdk:"state"`
	Weight           types.Int64  `tfsdk:"weight"`
}

func (r *LbmonitorServicegroupBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbmonitor_servicegroup_binding resource.",
			},
			"dup_state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State of the monitor. The state setting for a monitor of a given type affects all monitors of that type. For example, if an HTTP monitor is enabled, all HTTP monitors on the appliance are (or remain) enabled. If an HTTP monitor is disabled, all HTTP monitors on the appliance are disabled.",
			},
			"dup_weight": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Weight to assign to the binding between the monitor and service.",
			},
			"monitorname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the monitor.",
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the service group.",
			},
			"servicename": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the service or service group.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State of the monitor. The state setting for a monitor of a given type affects all monitors of that type. For example, if an HTTP monitor is enabled, all HTTP monitors on the appliance are (or remain) enabled. If an HTTP monitor is disabled, all HTTP monitors on the appliance are disabled.",
			},
			"weight": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Weight to assign to the binding between the monitor and service.",
			},
		},
	}
}

func lbmonitor_servicegroup_bindingGetThePayloadFromthePlan(ctx context.Context, data *LbmonitorServicegroupBindingResourceModel) lb.Lbmonitorservicegroupbinding {
	tflog.Debug(ctx, "In lbmonitor_servicegroup_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbmonitor_servicegroup_binding := lb.Lbmonitorservicegroupbinding{}
	if !data.DupState.IsNull() && !data.DupState.IsUnknown() {
		lbmonitor_servicegroup_binding.Dupstate = data.DupState.ValueString()
	}
	if !data.DupWeight.IsNull() && !data.DupWeight.IsUnknown() {
		lbmonitor_servicegroup_binding.Dupweight = utils.IntPtr(int(data.DupWeight.ValueInt64()))
	}
	if !data.Monitorname.IsNull() && !data.Monitorname.IsUnknown() {
		lbmonitor_servicegroup_binding.Monitorname = data.Monitorname.ValueString()
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		lbmonitor_servicegroup_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Servicename.IsNull() && !data.Servicename.IsUnknown() {
		lbmonitor_servicegroup_binding.Servicename = data.Servicename.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		lbmonitor_servicegroup_binding.State = data.State.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		lbmonitor_servicegroup_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return lbmonitor_servicegroup_binding
}
