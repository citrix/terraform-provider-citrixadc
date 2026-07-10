package systemglobal_authenticationldappolicy_binding

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

// SystemglobalAuthenticationldappolicyBindingResourceModel describes the resource data model.
type SystemglobalAuthenticationldappolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Builtin                types.List   `tfsdk:"builtin"`
	Feature                types.String `tfsdk:"feature"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Nextfactor             types.String `tfsdk:"nextfactor"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *SystemglobalAuthenticationldappolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemglobal_authenticationldappolicy_binding resource.",
			},
			"builtin": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.",
			},
			"nextfactor": schema.StringAttribute{
				Optional: true,
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The priority of the command policy.",
			},
		},
	}
}

func systemglobal_authenticationldappolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *SystemglobalAuthenticationldappolicyBindingResourceModel) system.Systemglobalauthenticationldappolicybinding {
	tflog.Debug(ctx, "In systemglobal_authenticationldappolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemglobal_authenticationldappolicy_binding := system.Systemglobalauthenticationldappolicybinding{}
	if !data.Builtin.IsNull() && !data.Builtin.IsUnknown() {
		builtinList := make([]string, 0, len(data.Builtin.Elements()))
		data.Builtin.ElementsAs(ctx, &builtinList, false)
		systemglobal_authenticationldappolicy_binding.Builtin = builtinList
	}
	if !data.Feature.IsNull() && !data.Feature.IsUnknown() {
		systemglobal_authenticationldappolicy_binding.Feature = data.Feature.ValueString()
	}
	if !data.Globalbindtype.IsNull() && !data.Globalbindtype.IsUnknown() {
		systemglobal_authenticationldappolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		systemglobal_authenticationldappolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Nextfactor.IsNull() && !data.Nextfactor.IsUnknown() {
		systemglobal_authenticationldappolicy_binding.Nextfactor = data.Nextfactor.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		systemglobal_authenticationldappolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		systemglobal_authenticationldappolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return systemglobal_authenticationldappolicy_binding
}

// systemglobal_authenticationldappolicy_bindingSetAttrFromGet updates the RESOURCE
// state from a GET response. The NITRO GET for this binding echoes only
// policyname, priority, feature and globalbindtype; it does NOT return
// gotopriorityexpression, nextfactor or builtin. Those write-only inputs are
// therefore preserved from the existing plan/state rather than nulled, so that
// Terraform does not see an "inconsistent result after apply" (Pattern 7/13).
func systemglobal_authenticationldappolicy_bindingSetAttrFromGet(ctx context.Context, data *SystemglobalAuthenticationldappolicyBindingResourceModel, getResponseData map[string]interface{}) *SystemglobalAuthenticationldappolicyBindingResourceModel {
	tflog.Debug(ctx, "In systemglobal_authenticationldappolicy_bindingSetAttrFromGet Function")

	// Echoed-back fields: copy from the GET response.
	if val, ok := getResponseData["feature"]; ok && val != nil {
		data.Feature = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["globalbindtype"]; ok && val != nil {
		data.Globalbindtype = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	// gotopriorityexpression, nextfactor and builtin are not returned by GET;
	// preserve the existing plan/state values (do not touch them).

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}

// systemglobal_authenticationldappolicy_bindingSetAttrFromGetForDatasource
// faithfully copies the GET response into the model for the datasource flow,
// which has no prior plan/state to preserve. Fields the GET does not return are
// set to null. It also sets data.Id since the datasource has no Create.
func systemglobal_authenticationldappolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SystemglobalAuthenticationldappolicyBindingResourceModel, getResponseData map[string]interface{}) *SystemglobalAuthenticationldappolicyBindingResourceModel {
	tflog.Debug(ctx, "In systemglobal_authenticationldappolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["feature"]; ok && val != nil {
		data.Feature = types.StringValue(val.(string))
	} else {
		data.Feature = types.StringNull()
	}
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
		} else {
			data.Priority = types.Int64Null()
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["builtin"]; ok && val != nil {
		if listVal, diags := types.ListValueFrom(ctx, types.StringType, val); !diags.HasError() {
			data.Builtin = listVal
		} else {
			data.Builtin = types.ListNull(types.StringType)
		}
	} else {
		data.Builtin = types.ListNull(types.StringType)
	}

	// Set ID for the datasource
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}
