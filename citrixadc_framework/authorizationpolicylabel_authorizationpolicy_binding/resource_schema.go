package authorizationpolicylabel_authorizationpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/authorization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthorizationpolicylabelAuthorizationpolicyBindingResourceModel describes the resource data model.
type AuthorizationpolicylabelAuthorizationpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *AuthorizationpolicylabelAuthorizationpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authorizationpolicylabel_authorizationpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then either forward the request or response to the specified virtual server or evaluate the specified policy label.",
			},
			"invoke_labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authorization policy label to which to bind the policy.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of invocation. Available settings function as follows:\n* reqvserver - Send the request to the specified request virtual server.\n* resvserver - Send the response to the specified response virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authorization policy to bind to the policy label.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}

func authorizationpolicylabel_authorizationpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AuthorizationpolicylabelAuthorizationpolicyBindingResourceModel) authorization.Authorizationpolicylabelauthorizationpolicybinding {
	tflog.Debug(ctx, "In authorizationpolicylabel_authorizationpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authorizationpolicylabel_authorizationpolicy_binding := authorization.Authorizationpolicylabelauthorizationpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() {
		authorizationpolicylabel_authorizationpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() {
		authorizationpolicylabel_authorizationpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.InvokeLabelname.IsNull() {
		authorizationpolicylabel_authorizationpolicy_binding.Invokelabelname = data.InvokeLabelname.ValueString()
	}
	if !data.Labelname.IsNull() {
		authorizationpolicylabel_authorizationpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() {
		authorizationpolicylabel_authorizationpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() {
		authorizationpolicylabel_authorizationpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		authorizationpolicylabel_authorizationpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return authorizationpolicylabel_authorizationpolicy_binding
}

func authorizationpolicylabel_authorizationpolicy_bindingSetAttrFromGet(ctx context.Context, data *AuthorizationpolicylabelAuthorizationpolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthorizationpolicylabelAuthorizationpolicyBindingResourceModel {
	tflog.Debug(ctx, "In authorizationpolicylabel_authorizationpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	} else {
		data.Invoke = types.BoolNull()
	}
	if val, ok := getResponseData["invoke_labelname"]; ok && val != nil {
		data.InvokeLabelname = types.StringValue(val.(string))
	} else {
		data.InvokeLabelname = types.StringNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	} else {
		data.Labeltype = types.StringNull()
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
