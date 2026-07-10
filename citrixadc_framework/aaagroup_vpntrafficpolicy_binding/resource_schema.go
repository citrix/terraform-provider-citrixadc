package aaagroup_vpntrafficpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaagroupVpntrafficpolicyBindingResourceModel describes the resource data model.
type AaagroupVpntrafficpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupname              types.String `tfsdk:"groupname"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
}

func (r *AaagroupVpntrafficpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaagroup_vpntrafficpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the group that you are binding.",
			},
			"policy": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The policy name.",
			},
			"priority": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bindpoint to which the policy is bound.",
			},
		},
	}
}

func aaagroup_vpntrafficpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *AaagroupVpntrafficpolicyBindingResourceModel) aaa.Aaagroupvpntrafficpolicybinding {
	tflog.Debug(ctx, "In aaagroup_vpntrafficpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	aaagroup_vpntrafficpolicy_binding := aaa.Aaagroupvpntrafficpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		aaagroup_vpntrafficpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		aaagroup_vpntrafficpolicy_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		aaagroup_vpntrafficpolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		aaagroup_vpntrafficpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		aaagroup_vpntrafficpolicy_binding.Type = data.Type.ValueString()
	}

	return aaagroup_vpntrafficpolicy_binding
}

// aaagroup_vpntrafficpolicy_bindingSetAttrFromGet is the resource-side setter.
// Pattern 7 (preserve-state): the user-driven, RequiresReplace attributes
// (priority, type, gotopriorityexpression) are preserved from the existing
// plan/state rather than overwritten from the GET response. The NITRO GET for
// this binding does not faithfully echo every configured value (e.g. type is a
// bindpoint filter and may be absent/normalized), which would otherwise trigger
// "inconsistent result after apply" once these attributes are no longer Computed.
func aaagroup_vpntrafficpolicy_bindingSetAttrFromGet(ctx context.Context, data *AaagroupVpntrafficpolicyBindingResourceModel, getResponseData map[string]interface{}) *AaagroupVpntrafficpolicyBindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_vpntrafficpolicy_bindingSetAttrFromGet Function")

	// Preserve user-configured gotopriorityexpression, priority and type from
	// plan/state. The NITRO GET for this binding does not faithfully echo these
	// back (gotopriorityexpression is not returned at all; type is a bindpoint
	// filter that may be omitted/normalized; priority is a RequiresReplace user
	// input). They are all Optional/Required (not Computed), so overwriting them
	// from GET would risk an "inconsistent result after apply" or leave an
	// unknown value. Only the identity keys (groupname, policy) are refreshed.

	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// aaagroup_vpntrafficpolicy_bindingSetAttrFromGetForDatasource faithfully copies
// every field from the GET response (the datasource has no prior plan/state to
// preserve) and sets the composite ID.
func aaagroup_vpntrafficpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AaagroupVpntrafficpolicyBindingResourceModel, getResponseData map[string]interface{}) *AaagroupVpntrafficpolicyBindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_vpntrafficpolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
