package rnatglobal_auditsyslogpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// RnatglobalAuditsyslogpolicyBindingResourceModel describes the resource data model.
type RnatglobalAuditsyslogpolicyBindingResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Policy   types.String `tfsdk:"policy"`
	Priority types.Int64  `tfsdk:"priority"`
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rnatglobal_auditsyslogpolicy_binding resource.",
			},
			"policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The policy Name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the policy.",
			},
		},
	}
}

func rnatglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *RnatglobalAuditsyslogpolicyBindingResourceModel) network.Rnatglobalauditsyslogpolicybinding {
	tflog.Debug(ctx, "In rnatglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	rnatglobal_auditsyslogpolicy_binding := network.Rnatglobalauditsyslogpolicybinding{}
	if !data.Policy.IsNull() {
		rnatglobal_auditsyslogpolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() {
		rnatglobal_auditsyslogpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return rnatglobal_auditsyslogpolicy_binding
}

func rnatglobal_auditsyslogpolicy_bindingSetAttrFromGet(ctx context.Context, data *RnatglobalAuditsyslogpolicyBindingResourceModel, getResponseData map[string]interface{}) *RnatglobalAuditsyslogpolicyBindingResourceModel {
	tflog.Debug(ctx, "In rnatglobal_auditsyslogpolicy_bindingSetAttrFromGet Function")

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
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
