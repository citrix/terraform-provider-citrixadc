package aaauser_auditsyslogpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaauserAuditsyslogpolicyBindingResourceModel describes the resource data model.
type AaauserAuditsyslogpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
	Username               types.String `tfsdk:"username"`
}

func (r *AaauserAuditsyslogpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaauser_auditsyslogpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The policy Name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the policy.  A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("REQUEST"),
				Description: "Bindpoint to which the policy is bound.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "User account to which to bind the policy.",
			},
		},
	}
}

func aaauser_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AaauserAuditsyslogpolicyBindingResourceModel) aaa.Aaauserauditsyslogpolicybinding {
	tflog.Debug(ctx, "In aaauser_auditsyslogpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaauser_auditsyslogpolicy_binding := aaa.Aaauserauditsyslogpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() {
		aaauser_auditsyslogpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Policy.IsNull() {
		aaauser_auditsyslogpolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() {
		aaauser_auditsyslogpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() {
		aaauser_auditsyslogpolicy_binding.Type = data.Type.ValueString()
	}
	if !data.Username.IsNull() {
		aaauser_auditsyslogpolicy_binding.Username = data.Username.ValueString()
	}

	return aaauser_auditsyslogpolicy_binding
}

func aaauser_auditsyslogpolicy_bindingSetAttrFromGet(ctx context.Context, data *AaauserAuditsyslogpolicyBindingResourceModel, getResponseData map[string]interface{}) *AaauserAuditsyslogpolicyBindingResourceModel {
	tflog.Debug(ctx, "In aaauser_auditsyslogpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
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
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Username.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
