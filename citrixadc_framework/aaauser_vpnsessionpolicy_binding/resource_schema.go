package aaauser_vpnsessionpolicy_binding

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

// AaauserVpnsessionpolicyBindingResourceModel describes the resource data model.
type AaauserVpnsessionpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
	Username               types.String `tfsdk:"username"`
}

func (r *AaauserVpnsessionpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaauser_vpnsessionpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
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

func aaauser_vpnsessionpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AaauserVpnsessionpolicyBindingResourceModel) aaa.Aaauservpnsessionpolicybinding {
	tflog.Debug(ctx, "In aaauser_vpnsessionpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaauser_vpnsessionpolicy_binding := aaa.Aaauservpnsessionpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() {
		aaauser_vpnsessionpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Policy.IsNull() {
		aaauser_vpnsessionpolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() {
		aaauser_vpnsessionpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() {
		aaauser_vpnsessionpolicy_binding.Type = data.Type.ValueString()
	}
	if !data.Username.IsNull() {
		aaauser_vpnsessionpolicy_binding.Username = data.Username.ValueString()
	}

	return aaauser_vpnsessionpolicy_binding
}

func aaauser_vpnsessionpolicy_bindingSetAttrFromGet(ctx context.Context, data *AaauserVpnsessionpolicyBindingResourceModel, getResponseData map[string]interface{}) *AaauserVpnsessionpolicyBindingResourceModel {
	tflog.Debug(ctx, "In aaauser_vpnsessionpolicy_bindingSetAttrFromGet Function")

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
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Policy.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Username.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
