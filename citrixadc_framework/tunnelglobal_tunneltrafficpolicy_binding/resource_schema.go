package tunnelglobal_tunneltrafficpolicy_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tunnel"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// TunnelglobalTunneltrafficpolicyBindingResourceModel describes the resource data model.
type TunnelglobalTunneltrafficpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Feature                types.String `tfsdk:"feature"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	State                  types.String `tfsdk:"state"`
	Type                   types.String `tfsdk:"type"`
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the tunnelglobal_tunneltrafficpolicy_binding resource.",
			},
			"feature": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The feature to be checked while applying this config",
			},
			"globalbindtype": schema.StringAttribute{
				// Not echoed by NITRO GET - Optional only (no Computed) so it resolves
				// to null after apply instead of staying unknown (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// Not echoed by NITRO GET - Optional only (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Policy name.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Priority.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Current state of the binding. If the binding is enabled, the policy is active.",
			},
			"type": schema.StringAttribute{
				// Bind point / GET filter - not echoed by NITRO GET. Optional only
				// (no Computed) so it resolves to null after apply (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind point to which the policy is bound.",
			},
		},
	}
}

func tunnelglobal_tunneltrafficpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *TunnelglobalTunneltrafficpolicyBindingResourceModel) tunnel.Tunnelglobaltunneltrafficpolicybinding {
	tflog.Debug(ctx, "In tunnelglobal_tunneltrafficpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	tunnelglobal_tunneltrafficpolicy_binding := tunnel.Tunnelglobaltunneltrafficpolicybinding{}
	if !data.Feature.IsNull() && !data.Feature.IsUnknown() {
		tunnelglobal_tunneltrafficpolicy_binding.Feature = data.Feature.ValueString()
	}
	if !data.Globalbindtype.IsNull() && !data.Globalbindtype.IsUnknown() {
		tunnelglobal_tunneltrafficpolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		tunnelglobal_tunneltrafficpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		tunnelglobal_tunneltrafficpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		tunnelglobal_tunneltrafficpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		tunnelglobal_tunneltrafficpolicy_binding.State = data.State.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		tunnelglobal_tunneltrafficpolicy_binding.Type = data.Type.ValueString()
	}

	return tunnelglobal_tunneltrafficpolicy_binding
}

// tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGet is the RESOURCE-side state
// setter. The NITRO GET for this binding echoes back only policyname, priority, state
// and feature; it does NOT return globalbindtype, gotopriorityexpression or type. For
// those non-echoed inputs we PRESERVE the existing plan/state value (Pattern 7) instead
// of nulling it, which would otherwise trigger an "inconsistent result after apply"
// error or a perpetual diff. The ID is preserved as set by Create (plain policyname).
func tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGet(ctx context.Context, data *TunnelglobalTunneltrafficpolicyBindingResourceModel, getResponseData map[string]interface{}) *TunnelglobalTunneltrafficpolicyBindingResourceModel {
	tflog.Debug(ctx, "In tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGet Function")

	// Echoed fields - safe to adopt from the GET response.
	if val, ok := getResponseData["feature"]; ok && val != nil {
		data.Feature = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	}
	// globalbindtype, gotopriorityexpression and type are non-echoed inputs - preserve
	// whatever the plan/state already holds (do not overwrite or null them out).

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	data.Id = types.StringValue(data.Policyname.ValueString())

	return data
}

// tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter. A datasource has no prior plan/state to preserve, so it
// faithfully copies every field present in the GET response and sets its own ID
// (plain policyname, matching the resource ID format).
func tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *TunnelglobalTunneltrafficpolicyBindingResourceModel, getResponseData map[string]interface{}) *TunnelglobalTunneltrafficpolicyBindingResourceModel {
	tflog.Debug(ctx, "In tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["feature"]; ok && val != nil {
		data.Feature = types.StringValue(val.(string))
	} else {
		data.Feature = types.StringNull()
	}
	if val, ok := getResponseData["globalbindtype"]; ok && val != nil {
		data.Globalbindtype = types.StringValue(val.(string))
	} else {
		data.Globalbindtype = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	// type is not echoed by GET - retain the value the datasource was queried with.

	// Set ID for the datasource (plain policyname, matching the resource).
	data.Id = types.StringValue(data.Policyname.ValueString())

	return data
}
