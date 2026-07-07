package vpnvserver_vpntrafficpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnvserverVpntrafficpolicyBindingResourceModel describes the resource data model.
type VpnvserverVpntrafficpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnvserverVpntrafficpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_vpntrafficpolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				// Not echoed back by NITRO for this binding; keep Optional only
				// (no Computed) so omitting it does not produce known-after-apply churn.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// Not echoed back by NITRO for this binding; Optional only.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupextraction": schema.BoolAttribute{
				// Not reliably echoed back by NITRO when omitted; Optional only.
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Binds the authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the virtual server.",
			},
			"policy": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the policy, if any, bound to the VPN virtual server.",
			},
			"priority": schema.Int64Attribute{
				// Not reliably echoed back by NITRO for this binding; Optional only.
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				// Not reliably echoed back by NITRO when omitted; Optional only.
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.",
			},
		},
	}
}

func vpnvserver_vpntrafficpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverVpntrafficpolicyBindingResourceModel) vpn.Vpnvservervpntrafficpolicybinding {
	tflog.Debug(ctx, "In vpnvserver_vpntrafficpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_vpntrafficpolicy_binding := vpn.Vpnvservervpntrafficpolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		vpnvserver_vpntrafficpolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnvserver_vpntrafficpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnvserver_vpntrafficpolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_vpntrafficpolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		vpnvserver_vpntrafficpolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnvserver_vpntrafficpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnvserver_vpntrafficpolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnvserver_vpntrafficpolicy_binding
}

// vpnvserver_vpntrafficpolicy_bindingSetAttrFromGet is the resource-side setter.
// NITRO does NOT echo back bindpoint, gotopriorityexpression or priority for this
// binding (see the SDK v2 resource), so those user inputs are PRESERVED from the
// existing plan/state rather than overwritten (otherwise Terraform reports an
// "inconsistent result after apply"). Only the fields NITRO actually returns
// (groupextraction, name, policy, secondary) are read from the GET response.
func vpnvserver_vpntrafficpolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverVpntrafficpolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverVpntrafficpolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_vpntrafficpolicy_bindingSetAttrFromGet Function")

	// bindpoint, gotopriorityexpression, priority, groupextraction and secondary
	// are either not echoed back or returned in a normalized form (e.g. NITRO
	// returns false for an omitted bool) that does not exactly match the user's
	// config. All of these are RequiresReplace, Optional-only inputs, so we
	// PRESERVE the existing plan/state values to avoid "inconsistent result after
	// apply" and spurious replace-forcing diffs. Only adopt the identity fields
	// (name, policy) that the GET reliably returns; on import they populate the
	// model from the live object.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	// Composite ID uses the legacy SDK v2 key order: name,policy
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
	return data
}

// vpnvserver_vpntrafficpolicy_bindingSetAttrFromGetForDatasource is the datasource-side
// setter. A datasource has no prior plan/state to preserve, so it faithfully copies
// every field present in the GET response and sets the composite ID itself.
func vpnvserver_vpntrafficpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnvserverVpntrafficpolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverVpntrafficpolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_vpntrafficpolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	} else {
		data.Bindpoint = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
	} else {
		data.Groupextraction = types.BoolNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	} else {
		data.Policy = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	} else {
		data.Secondary = types.BoolNull()
	}

	// Set the composite ID (legacy key order: name,policy).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
