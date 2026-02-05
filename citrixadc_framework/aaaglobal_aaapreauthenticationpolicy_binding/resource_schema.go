package aaaglobal_aaapreauthenticationpolicy_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaaglobalAaapreauthenticationpolicyBindingResourceModel describes the resource data model.
type AaaglobalAaapreauthenticationpolicyBindingResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Policy   types.String `tfsdk:"policy"`
	Priority types.Int64  `tfsdk:"priority"`
}

func (r *AaaglobalAaapreauthenticationpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaaglobal_aaapreauthenticationpolicy_binding resource.",
			},
			"policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy to be unbound.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the bound policy",
			},
		},
	}
}

func aaaglobal_aaapreauthenticationpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AaaglobalAaapreauthenticationpolicyBindingResourceModel) aaa.Aaaglobalaaapreauthenticationpolicybinding {
	tflog.Debug(ctx, "In aaaglobal_aaapreauthenticationpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaaglobal_aaapreauthenticationpolicy_binding := aaa.Aaaglobalaaapreauthenticationpolicybinding{}
	if !data.Policy.IsNull() {
		aaaglobal_aaapreauthenticationpolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() {
		aaaglobal_aaapreauthenticationpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return aaaglobal_aaapreauthenticationpolicy_binding
}

func aaaglobal_aaapreauthenticationpolicy_bindingSetAttrFromGet(ctx context.Context, data *AaaglobalAaapreauthenticationpolicyBindingResourceModel, getResponseData map[string]interface{}) *AaaglobalAaapreauthenticationpolicyBindingResourceModel {
	tflog.Debug(ctx, "In aaaglobal_aaapreauthenticationpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
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

	// Set ID for the resource
	data.Id = types.StringValue(data.Policy.ValueString())

	return data
}
