package systemglobal_authenticationradiuspolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SystemglobalAuthenticationradiuspolicyBindingResourceModel describes the resource data model.
type SystemglobalAuthenticationradiuspolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Builtin                types.List   `tfsdk:"builtin"`
	Feature                types.String `tfsdk:"feature"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Nextfactor             types.String `tfsdk:"nextfactor"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemglobal_authenticationradiuspolicy_binding resource.",
			},
			"builtin": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"feature": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The feature to be checked while applying this config",
			},
			"globalbindtype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.",
			},
			"nextfactor": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "On success invoke label. Applicable for advanced authentication policy binding",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the  command policy.",
			},
			"priority": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The priority of the command policy.",
			},
		},
	}
}

func systemglobal_authenticationradiuspolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *SystemglobalAuthenticationradiuspolicyBindingResourceModel) system.Systemglobalauthenticationradiuspolicybinding {
	tflog.Debug(ctx, "In systemglobal_authenticationradiuspolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemglobal_authenticationradiuspolicy_binding := system.Systemglobalauthenticationradiuspolicybinding{}
	if !data.Builtin.IsNull() && !data.Builtin.IsUnknown() {
		builtin := make([]string, 0, len(data.Builtin.Elements()))
		data.Builtin.ElementsAs(ctx, &builtin, false)
		systemglobal_authenticationradiuspolicy_binding.Builtin = builtin
	}
	if !data.Feature.IsNull() && !data.Feature.IsUnknown() {
		systemglobal_authenticationradiuspolicy_binding.Feature = data.Feature.ValueString()
	}
	if !data.Globalbindtype.IsNull() && !data.Globalbindtype.IsUnknown() {
		systemglobal_authenticationradiuspolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		systemglobal_authenticationradiuspolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Nextfactor.IsNull() && !data.Nextfactor.IsUnknown() {
		systemglobal_authenticationradiuspolicy_binding.Nextfactor = data.Nextfactor.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		systemglobal_authenticationradiuspolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		systemglobal_authenticationradiuspolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return systemglobal_authenticationradiuspolicy_binding
}

// systemglobal_authenticationradiuspolicy_bindingSetAttrFromGet is the RESOURCE-side state setter.
// The NITRO GET for this binding does not faithfully echo every configured input (priority,
// gotopriorityexpression, feature, builtin, etc. are server-overridden or non-echoed), so to avoid
// "inconsistent result after apply" we preserve the values already present in the model (plan/state)
// and only resolve a Computed attribute from the GET response when the model has no value.
func systemglobal_authenticationradiuspolicy_bindingSetAttrFromGet(ctx context.Context, data *SystemglobalAuthenticationradiuspolicyBindingResourceModel, getResponseData map[string]interface{}) *SystemglobalAuthenticationradiuspolicyBindingResourceModel {
	tflog.Debug(ctx, "In systemglobal_authenticationradiuspolicy_bindingSetAttrFromGet Function")

	// policyname is the key; adopt the GET value only if the model does not already hold it
	// (covers import, where state carries only the ID).
	if data.Policyname.IsNull() || data.Policyname.IsUnknown() {
		if val, ok := getResponseData["policyname"]; ok && val != nil {
			data.Policyname = types.StringValue(fmt.Sprintf("%v", val))
		}
	}

	// builtin is Optional+Computed and may be null in the plan but returned by GET. Resolve it from
	// the GET response when the model has no value so the Computed attribute is never left unknown.
	if data.Builtin.IsNull() || data.Builtin.IsUnknown() {
		data.Builtin = builtinFromGet(ctx, getResponseData)
	}

	// The remaining Optional+Computed attributes are resolved from the GET response only when the
	// model has no value; otherwise the configured/state value is preserved to keep the apply
	// consistent.
	if data.Feature.IsNull() || data.Feature.IsUnknown() {
		data.Feature = stringFromGet(getResponseData, "feature")
	}
	if data.Globalbindtype.IsNull() || data.Globalbindtype.IsUnknown() {
		data.Globalbindtype = stringFromGet(getResponseData, "globalbindtype")
	}
	if data.Gotopriorityexpression.IsNull() || data.Gotopriorityexpression.IsUnknown() {
		data.Gotopriorityexpression = stringFromGet(getResponseData, "gotopriorityexpression")
	}
	if data.Nextfactor.IsNull() || data.Nextfactor.IsUnknown() {
		data.Nextfactor = stringFromGet(getResponseData, "nextfactor")
	}

	// priority is Required (user-driven); keep the configured value. On import it is resolved from GET.
	if data.Priority.IsNull() || data.Priority.IsUnknown() {
		if val, ok := getResponseData["priority"]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				data.Priority = types.Int64Value(intVal)
			}
		}
	}

	// Set ID for the resource (single unique attribute -> plain value).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}

// systemglobal_authenticationradiuspolicy_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// setter. Datasources have no prior plan/state to preserve, so this faithfully copies every field
// from the GET response and sets the ID.
func systemglobal_authenticationradiuspolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SystemglobalAuthenticationradiuspolicyBindingResourceModel, getResponseData map[string]interface{}) *SystemglobalAuthenticationradiuspolicyBindingResourceModel {
	tflog.Debug(ctx, "In systemglobal_authenticationradiuspolicy_bindingSetAttrFromGetForDatasource Function")

	data.Builtin = builtinFromGet(ctx, getResponseData)
	data.Feature = stringFromGet(getResponseData, "feature")
	data.Globalbindtype = stringFromGet(getResponseData, "globalbindtype")
	data.Gotopriorityexpression = stringFromGet(getResponseData, "gotopriorityexpression")
	data.Nextfactor = stringFromGet(getResponseData, "nextfactor")
	data.Policyname = stringFromGet(getResponseData, "policyname")
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}

	// Set ID for the datasource (single unique attribute -> plain value).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}

func stringFromGet(getResponseData map[string]interface{}, key string) types.String {
	if val, ok := getResponseData[key]; ok && val != nil {
		return types.StringValue(fmt.Sprintf("%v", val))
	}
	return types.StringNull()
}

func builtinFromGet(ctx context.Context, getResponseData map[string]interface{}) types.List {
	if val, ok := getResponseData["builtin"]; ok && val != nil {
		if listVal, ok := val.([]interface{}); ok {
			strs := make([]string, 0, len(listVal))
			for _, e := range listVal {
				strs = append(strs, fmt.Sprintf("%v", e))
			}
			listValue, _ := types.ListValueFrom(ctx, types.StringType, strs)
			return listValue
		}
	}
	return types.ListNull(types.StringType)
}
