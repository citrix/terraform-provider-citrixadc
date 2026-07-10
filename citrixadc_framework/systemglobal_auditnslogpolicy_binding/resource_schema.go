package systemglobal_auditnslogpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

// SystemglobalAuditnslogpolicyBindingResourceModel describes the resource data model.
type SystemglobalAuditnslogpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Builtin                types.List   `tfsdk:"builtin"`
	Feature                types.String `tfsdk:"feature"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Nextfactor             types.String `tfsdk:"nextfactor"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemglobal_auditnslogpolicy_binding resource.",
			},
			"builtin": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				// Not Computed: NITRO GET never echoes builtin, so a Computed value
				// would stay unknown after apply (Pattern 13 / Pattern 7).
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
				// Not Computed: NITRO GET never echoes gotopriorityexpression.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.",
			},
			"nextfactor": schema.StringAttribute{
				Optional: true,
				// Not Computed: NITRO GET never echoes nextfactor.
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

func systemglobal_auditnslogpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *SystemglobalAuditnslogpolicyBindingResourceModel) system.Systemglobalauditnslogpolicybinding {
	tflog.Debug(ctx, "In systemglobal_auditnslogpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemglobal_auditnslogpolicy_binding := system.Systemglobalauditnslogpolicybinding{}
	if !data.Builtin.IsNull() && !data.Builtin.IsUnknown() {
		builtin := make([]string, 0, len(data.Builtin.Elements()))
		data.Builtin.ElementsAs(ctx, &builtin, false)
		systemglobal_auditnslogpolicy_binding.Builtin = builtin
	}
	if !data.Feature.IsNull() && !data.Feature.IsUnknown() {
		systemglobal_auditnslogpolicy_binding.Feature = data.Feature.ValueString()
	}
	if !data.Globalbindtype.IsNull() && !data.Globalbindtype.IsUnknown() {
		systemglobal_auditnslogpolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		systemglobal_auditnslogpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Nextfactor.IsNull() && !data.Nextfactor.IsUnknown() {
		systemglobal_auditnslogpolicy_binding.Nextfactor = data.Nextfactor.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		systemglobal_auditnslogpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		systemglobal_auditnslogpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return systemglobal_auditnslogpolicy_binding
}

// systemglobal_auditnslogpolicy_bindingSetAttrFromGet is the RESOURCE setter.
// The NITRO GET response for this binding only echoes back policyname, priority,
// feature, globalbindtype (and a few read-only flags). It does NOT echo back
// gotopriorityexpression, nextfactor or builtin. For those non-echoed,
// Optional+Computed inputs we preserve the prior plan/state value rather than
// nulling it (Pattern 7) so an apply that configured them does not error with
// "inconsistent result after apply". The ID is set once in Create, not here.
func systemglobal_auditnslogpolicy_bindingSetAttrFromGet(ctx context.Context, data *SystemglobalAuditnslogpolicyBindingResourceModel, getResponseData map[string]interface{}) *SystemglobalAuditnslogpolicyBindingResourceModel {
	tflog.Debug(ctx, "In systemglobal_auditnslogpolicy_bindingSetAttrFromGet Function")

	// builtin: not echoed by GET - preserve existing plan/state value unless present.
	if val, ok := getResponseData["builtin"]; ok && val != nil {
		if rawList, ok := val.([]interface{}); ok {
			elems := make([]attr.Value, 0, len(rawList))
			for _, item := range rawList {
				elems = append(elems, types.StringValue(fmt.Sprintf("%v", item)))
			}
			if listVal, diags := types.ListValue(types.StringType, elems); !diags.HasError() {
				data.Builtin = listVal
			}
		}
	}
	// feature: server-assigned (e.g. SYSTEM), echoed by GET.
	if val, ok := getResponseData["feature"]; ok && val != nil {
		data.Feature = types.StringValue(val.(string))
	}
	// globalbindtype: server-assigned (e.g. SYSTEM_GLOBAL), echoed by GET.
	if val, ok := getResponseData["globalbindtype"]; ok && val != nil {
		data.Globalbindtype = types.StringValue(val.(string))
	}
	// gotopriorityexpression: not echoed by GET - preserve existing value.
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	// nextfactor: not echoed by GET - preserve existing value.
	if val, ok := getResponseData["nextfactor"]; ok && val != nil {
		data.Nextfactor = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	return data
}

// systemglobal_auditnslogpolicy_bindingSetAttrFromGetForDatasource is the DATASOURCE
// setter. The datasource has no prior plan/state to preserve, so it faithfully copies
// every field from the GET response (nulling absent ones) and sets the ID itself
// (Pattern 7).
func systemglobal_auditnslogpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SystemglobalAuditnslogpolicyBindingResourceModel, getResponseData map[string]interface{}) *SystemglobalAuditnslogpolicyBindingResourceModel {
	tflog.Debug(ctx, "In systemglobal_auditnslogpolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["builtin"]; ok && val != nil {
		if rawList, ok := val.([]interface{}); ok {
			elems := make([]attr.Value, 0, len(rawList))
			for _, item := range rawList {
				elems = append(elems, types.StringValue(fmt.Sprintf("%v", item)))
			}
			if listVal, diags := types.ListValue(types.StringType, elems); !diags.HasError() {
				data.Builtin = listVal
			} else {
				data.Builtin = types.ListNull(types.StringType)
			}
		} else {
			data.Builtin = types.ListNull(types.StringType)
		}
	} else {
		data.Builtin = types.ListNull(types.StringType)
	}
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
		}
	} else {
		data.Priority = types.Int64Null()
	}

	// Set ID for the datasource (it has no Create) - plain single-key value.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}
