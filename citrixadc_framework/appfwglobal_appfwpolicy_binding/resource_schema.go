package appfwglobal_appfwpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwglobalAppfwpolicyBindingResourceModel describes the resource data model.
type AppfwglobalAppfwpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	State                  types.String `tfsdk:"state"`
	Type                   types.String `tfsdk:"type"`
}

func (r *AppfwglobalAppfwpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwglobal_appfwpolicy_binding resource.",
			},
			"globalbindtype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("SYSTEM_GLOBAL"),
				Description: "0",
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
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label invocation.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the policy.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "The priority of the policy.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the binding to activate or deactivate the policy. This is applicable to classic policies only.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind point to which to policy is bound.",
			},
		},
	}
}

func appfwglobal_appfwpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwglobalAppfwpolicyBindingResourceModel) appfw.Appfwglobalappfwpolicybinding {
	tflog.Debug(ctx, "In appfwglobal_appfwpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwglobal_appfwpolicy_binding := appfw.Appfwglobalappfwpolicybinding{}
	if !data.Globalbindtype.IsNull() {
		appfwglobal_appfwpolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() {
		appfwglobal_appfwpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() {
		appfwglobal_appfwpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() {
		appfwglobal_appfwpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() {
		appfwglobal_appfwpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() {
		appfwglobal_appfwpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		appfwglobal_appfwpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.State.IsNull() {
		appfwglobal_appfwpolicy_binding.State = data.State.ValueString()
	}
	if !data.Type.IsNull() {
		appfwglobal_appfwpolicy_binding.Type = data.Type.ValueString()
	}

	return appfwglobal_appfwpolicy_binding
}

func appfwglobal_appfwpolicy_bindingSetAttrFromGet(ctx context.Context, data *AppfwglobalAppfwpolicyBindingResourceModel, getResponseData map[string]interface{}) *AppfwglobalAppfwpolicyBindingResourceModel {
	tflog.Debug(ctx, "In appfwglobal_appfwpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	} else {
		data.Invoke = types.BoolNull()
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
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
