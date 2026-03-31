package appfwpolicylabel_appfwpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwpolicylabelAppfwpolicyBindingResourceModel describes the resource data model.
type AppfwpolicylabelAppfwpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *AppfwpolicylabelAppfwpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwpolicylabel_appfwpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.",
			},
			"invoke_labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the application firewall policy label.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label to invoke if the current policy evaluates to TRUE and the invoke parameter is set. Available settings function as follows:\n* reqvserver. Invoke the unnamed policy label associated with the specified request virtual server.\n* policylabel. Invoke the specified user-defined policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the application firewall policy to bind to the policy label.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Positive integer specifying the priority of the policy. A lower number specifies a higher priority. Must be unique within a group of policies that are bound to the same bind point or label. Policies are evaluated in the order of their priority numbers.",
			},
		},
	}
}

func appfwpolicylabel_appfwpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwpolicylabelAppfwpolicyBindingResourceModel) appfw.Appfwpolicylabelappfwpolicybinding {
	tflog.Debug(ctx, "In appfwpolicylabel_appfwpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwpolicylabel_appfwpolicy_binding := appfw.Appfwpolicylabelappfwpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() {
		appfwpolicylabel_appfwpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() {
		appfwpolicylabel_appfwpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.InvokeLabelname.IsNull() {
		appfwpolicylabel_appfwpolicy_binding.Invokelabelname = data.InvokeLabelname.ValueString()
	}
	if !data.Labelname.IsNull() {
		appfwpolicylabel_appfwpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() {
		appfwpolicylabel_appfwpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() {
		appfwpolicylabel_appfwpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		appfwpolicylabel_appfwpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return appfwpolicylabel_appfwpolicy_binding
}

func appfwpolicylabel_appfwpolicy_bindingSetAttrFromGet(ctx context.Context, data *AppfwpolicylabelAppfwpolicyBindingResourceModel, getResponseData map[string]interface{}) *AppfwpolicylabelAppfwpolicyBindingResourceModel {
	tflog.Debug(ctx, "In appfwpolicylabel_appfwpolicy_bindingSetAttrFromGet Function")

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
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
