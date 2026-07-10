package vpnvserver_aaapreauthenticationpolicy_binding

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

// VpnvserverAaapreauthenticationpolicyBindingResourceModel describes the resource data model.
type VpnvserverAaapreauthenticationpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnvserverAaapreauthenticationpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_aaapreauthenticationpolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupextraction": schema.BoolAttribute{
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
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.",
			},
		},
	}
}

func vpnvserver_aaapreauthenticationpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverAaapreauthenticationpolicyBindingResourceModel) vpn.Vpnvserveraaapreauthenticationpolicybinding {
	tflog.Debug(ctx, "In vpnvserver_aaapreauthenticationpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_aaapreauthenticationpolicy_binding := vpn.Vpnvserveraaapreauthenticationpolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		vpnvserver_aaapreauthenticationpolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnvserver_aaapreauthenticationpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnvserver_aaapreauthenticationpolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_aaapreauthenticationpolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		vpnvserver_aaapreauthenticationpolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnvserver_aaapreauthenticationpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnvserver_aaapreauthenticationpolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnvserver_aaapreauthenticationpolicy_binding
}

// vpnvserver_aaapreauthenticationpolicy_bindingSetAttrFromGet is the RESOURCE setter.
// All attributes are RequiresReplace inputs and several of them (bindpoint,
// gotopriorityexpression, priority, secondary, groupextraction) are not reliably
// echoed back (or are returned normalized) by the binding GET response. To avoid
// "inconsistent result after apply" (Pattern 7 / Pattern 13), preserve the
// configured plan/state values; only adopt a GET value when the current model
// value is null/unknown (e.g. on import, where state carries only the ID).
func vpnvserver_aaapreauthenticationpolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverAaapreauthenticationpolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAaapreauthenticationpolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_aaapreauthenticationpolicy_bindingSetAttrFromGet Function")

	// name and policy form the identity; always adopt from the GET response.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	}

	// All remaining attributes are Optional, RequiresReplace inputs that the binding
	// GET response does not reliably echo (the SDK v2 resource likewise did not read
	// bindpoint back, and the others are returned normalized or absent). Preserving
	// the configured plan/state value verbatim avoids "inconsistent result after
	// apply" (Pattern 7 / Pattern 13). They are intentionally NOT overwritten here.

	// Set ID for the resource (legacy SDK v2 order: name,policy)
	data.Id = types.StringValue(vpnvserver_aaapreauthenticationpolicy_bindingBuildId(data))

	return data
}

// vpnvserver_aaapreauthenticationpolicy_bindingSetAttrFromGetForDatasource is the
// DATASOURCE setter. A datasource has no prior plan/state to preserve, so it
// faithfully copies every field from the GET response and sets its own ID
// (Pattern 7 datasource split).
func vpnvserver_aaapreauthenticationpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnvserverAaapreauthenticationpolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAaapreauthenticationpolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_aaapreauthenticationpolicy_bindingSetAttrFromGetForDatasource Function")

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

	data.Id = types.StringValue(vpnvserver_aaapreauthenticationpolicy_bindingBuildId(data))

	return data
}

// vpnvserver_aaapreauthenticationpolicy_bindingBuildId composes the resource ID in
// the new key:value format using the legacy SDK v2 attribute order (name,policy),
// matching resource_id_mapping.json so legacy imports resolve via ParseIdString.
func vpnvserver_aaapreauthenticationpolicy_bindingBuildId(data *VpnvserverAaapreauthenticationpolicyBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(data.Policy.ValueString())))
	return strings.Join(idParts, ",")
}
