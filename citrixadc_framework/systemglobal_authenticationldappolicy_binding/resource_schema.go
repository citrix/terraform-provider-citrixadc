package systemglobal_authenticationldappolicy_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SystemglobalAuthenticationldappolicyBindingResourceModel describes the resource data model.
type SystemglobalAuthenticationldappolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("SYSTEM_GLOBAL"),
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.",
			},
			"nextfactor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On success invoke label. Applicable for advanced authentication policy binding",
			},
			"policyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the  command policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the command policy.",
			},
		},
	}
}

func systemglobal_authenticationldappolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SystemglobalAuthenticationldappolicyBindingResourceModel) system.Systemglobalauthenticationldappolicybinding {
	tflog.Debug(ctx, "In systemglobal_authenticationldappolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemglobal_authenticationldappolicy_binding := system.Systemglobalauthenticationldappolicybinding{}
	if !data.Feature.IsNull() {
		systemglobal_authenticationldappolicy_binding.Feature = data.Feature.ValueString()
	}
	if !data.Globalbindtype.IsNull() {
		systemglobal_authenticationldappolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() {
		systemglobal_authenticationldappolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Nextfactor.IsNull() {
		systemglobal_authenticationldappolicy_binding.Nextfactor = data.Nextfactor.ValueString()
	}
	if !data.Policyname.IsNull() {
		systemglobal_authenticationldappolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		systemglobal_authenticationldappolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return systemglobal_authenticationldappolicy_binding
}

func systemglobal_authenticationldappolicy_bindingSetAttrFromGet(ctx context.Context, data *SystemglobalAuthenticationldappolicyBindingResourceModel, getResponseData map[string]interface{}) *SystemglobalAuthenticationldappolicyBindingResourceModel {
	tflog.Debug(ctx, "In systemglobal_authenticationldappolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Policyname.ValueString())

	return data
}
