package tmglobal_tmsessionpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// TmglobalTmsessionpolicyBindingResourceModel describes the resource data model.
type TmglobalTmsessionpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Feature                types.String `tfsdk:"feature"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *TmglobalTmsessionpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the tmglobal_tmsessionpolicy_binding resource.",
			},
			"feature": schema.StringAttribute{
				// Read-only: 'feature' is returned by GET but is NOT a valid
				// 'bind tm global' / add-payload argument (NitroValidator: erroneously-included read-only property).
				Computed:    true,
				Description: "The feature to be checked while applying this config",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the policy.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The priority of the policy.",
			},
		},
	}
}

func tmglobal_tmsessionpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *TmglobalTmsessionpolicyBindingResourceModel) tm.Tmglobaltmsessionpolicybinding {
	tflog.Debug(ctx, "In tmglobal_tmsessionpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	tmglobal_tmsessionpolicy_binding := tm.Tmglobaltmsessionpolicybinding{}
	// 'feature' is read-only (returned by GET only); it is NOT a valid add/bind arg, so it is excluded from the PUT payload.
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		tmglobal_tmsessionpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		tmglobal_tmsessionpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		tmglobal_tmsessionpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return tmglobal_tmsessionpolicy_binding
}

func tmglobal_tmsessionpolicy_bindingSetAttrFromGet(ctx context.Context, data *TmglobalTmsessionpolicyBindingResourceModel, getResponseData map[string]interface{}) *TmglobalTmsessionpolicyBindingResourceModel {
	tflog.Debug(ctx, "In tmglobal_tmsessionpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["feature"]; ok && val != nil {
		data.Feature = types.StringValue(val.(string))
	} else {
		data.Feature = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
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
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}
