package systemuser_systemcmdpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SystemuserSystemcmdpolicyBindingResourceModel describes the resource data model.
type SystemuserSystemcmdpolicyBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Policyname types.String `tfsdk:"policyname"`
	Priority   types.Int64  `tfsdk:"priority"`
	Username   types.String `tfsdk:"username"`
}

func (r *SystemuserSystemcmdpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemuser_systemcmdpolicy_binding resource.",
			},
			"policyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of command policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the policy.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of the system-user entry to which to bind the command policy.",
			},
		},
	}
}

func systemuser_systemcmdpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SystemuserSystemcmdpolicyBindingResourceModel) system.Systemusersystemcmdpolicybinding {
	tflog.Debug(ctx, "In systemuser_systemcmdpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemuser_systemcmdpolicy_binding := system.Systemusersystemcmdpolicybinding{}
	if !data.Policyname.IsNull() {
		systemuser_systemcmdpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		systemuser_systemcmdpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Username.IsNull() {
		systemuser_systemcmdpolicy_binding.Username = data.Username.ValueString()
	}

	return systemuser_systemcmdpolicy_binding
}

func systemuser_systemcmdpolicy_bindingSetAttrFromGet(ctx context.Context, data *SystemuserSystemcmdpolicyBindingResourceModel, getResponseData map[string]interface{}) *SystemuserSystemcmdpolicyBindingResourceModel {
	tflog.Debug(ctx, "In systemuser_systemcmdpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Username.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
