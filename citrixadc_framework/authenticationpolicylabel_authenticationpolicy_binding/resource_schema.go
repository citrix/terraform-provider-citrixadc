package authenticationpolicylabel_authenticationpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel describes the resource data model.
type AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Labelname              types.String `tfsdk:"labelname"`
	Nextfactor             types.String `tfsdk:"nextfactor"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationpolicylabel_authenticationpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authentication policy label to which to bind the policy.",
			},
			"nextfactor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On success invoke label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authentication policy to bind to the policy label.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}

func authenticationpolicylabel_authenticationpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel) authentication.Authenticationpolicylabelauthenticationpolicybinding {
	tflog.Debug(ctx, "In authenticationpolicylabel_authenticationpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationpolicylabel_authenticationpolicy_binding := authentication.Authenticationpolicylabelauthenticationpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() {
		authenticationpolicylabel_authenticationpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Labelname.IsNull() {
		authenticationpolicylabel_authenticationpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Nextfactor.IsNull() {
		authenticationpolicylabel_authenticationpolicy_binding.Nextfactor = data.Nextfactor.ValueString()
	}
	if !data.Policyname.IsNull() {
		authenticationpolicylabel_authenticationpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		authenticationpolicylabel_authenticationpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return authenticationpolicylabel_authenticationpolicy_binding
}

func authenticationpolicylabel_authenticationpolicy_bindingSetAttrFromGet(ctx context.Context, data *AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationpolicylabel_authenticationpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["nextfactor"]; ok && val != nil {
		data.Nextfactor = types.StringValue(val.(string))
	} else {
		data.Nextfactor = types.StringNull()
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
