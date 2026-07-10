package vpnvserver_authenticationdfapolicy_binding

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

// VpnvserverAuthenticationdfapolicyBindingResourceModel describes the resource data model.
type VpnvserverAuthenticationdfapolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnvserverAuthenticationdfapolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_authenticationdfapolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				// Not echoed back by NITRO GET; Computed would leave it unknown after
				// apply and trigger "inconsistent result". Optional-only (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// Not echoed back by NITRO GET; Optional-only (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupextraction": schema.BoolAttribute{
				// Not echoed back by NITRO GET; Optional-only (Pattern 13).
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.",
			},
		},
	}
}

func vpnvserver_authenticationdfapolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverAuthenticationdfapolicyBindingResourceModel) vpn.Vpnvserverauthenticationdfapolicybinding {
	tflog.Debug(ctx, "In vpnvserver_authenticationdfapolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_authenticationdfapolicy_binding := vpn.Vpnvserverauthenticationdfapolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		vpnvserver_authenticationdfapolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnvserver_authenticationdfapolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnvserver_authenticationdfapolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_authenticationdfapolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		vpnvserver_authenticationdfapolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnvserver_authenticationdfapolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnvserver_authenticationdfapolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnvserver_authenticationdfapolicy_binding
}

// vpnvserver_authenticationdfapolicy_bindingSetAttrFromGet is the resource-side setter.
// All schema attributes are RequiresReplace (no NITRO update endpoint), and several
// inputs (bindpoint, gotopriorityexpression, groupextraction, priority, secondary) are
// either not echoed back by NITRO or are server-normalized. Overwriting them from the
// GET response (or nulling them when absent) produces "inconsistent result after apply"
// errors and spurious diffs. So preserve the plan/state values: only adopt values from
// the GET response when the current model value is null/unknown (covers import, where
// state carries only the ID). Pattern 7 (preserve plan/state on server-overridden /
// non-echoed inputs).
func vpnvserver_authenticationdfapolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverAuthenticationdfapolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAuthenticationdfapolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_authenticationdfapolicy_bindingSetAttrFromGet Function")

	// name and policy are the identity keys; adopt from GET only if missing (import).
	if data.Name.IsNull() || data.Name.IsUnknown() {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	if data.Policy.IsNull() || data.Policy.IsUnknown() {
		if val, ok := getResponseData["policy"]; ok && val != nil {
			data.Policy = types.StringValue(val.(string))
		}
	}

	// Non-echoed / server-normalized inputs: only populate when the model has no value
	// yet (import). Otherwise preserve the configured plan/state value to avoid drift.
	if data.Bindpoint.IsNull() || data.Bindpoint.IsUnknown() {
		if val, ok := getResponseData["bindpoint"]; ok && val != nil {
			data.Bindpoint = types.StringValue(val.(string))
		}
	}
	if data.Gotopriorityexpression.IsNull() || data.Gotopriorityexpression.IsUnknown() {
		if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
			data.Gotopriorityexpression = types.StringValue(val.(string))
		}
	}
	if data.Groupextraction.IsNull() || data.Groupextraction.IsUnknown() {
		if val, ok := getResponseData["groupextraction"]; ok && val != nil {
			data.Groupextraction = types.BoolValue(val.(bool))
		}
	}
	if data.Priority.IsNull() || data.Priority.IsUnknown() {
		if val, ok := getResponseData["priority"]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				data.Priority = types.Int64Value(intVal)
			}
		}
	}
	if data.Secondary.IsNull() || data.Secondary.IsUnknown() {
		if val, ok := getResponseData["secondary"]; ok && val != nil {
			data.Secondary = types.BoolValue(val.(bool))
		}
	}

	// Set ID for the resource (legacy format: name,policy).
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// vpnvserver_authenticationdfapolicy_bindingSetAttrFromGetForDatasource is the
// datasource-side setter: it faithfully copies every field from the GET response (the
// datasource has no prior plan/state to preserve) and sets the ID. Pattern 7 split.
func vpnvserver_authenticationdfapolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnvserverAuthenticationdfapolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAuthenticationdfapolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_authenticationdfapolicy_bindingSetAttrFromGetForDatasource Function")

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
		} else {
			data.Priority = types.Int64Null()
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	} else {
		data.Secondary = types.BoolNull()
	}

	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
