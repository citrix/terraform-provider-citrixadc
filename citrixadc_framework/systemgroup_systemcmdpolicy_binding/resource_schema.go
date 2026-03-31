package systemgroup_systemcmdpolicy_binding

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

// SystemgroupSystemcmdpolicyBindingResourceModel describes the resource data model.
type SystemgroupSystemcmdpolicyBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Groupname  types.String `tfsdk:"groupname"`
	Policyname types.String `tfsdk:"policyname"`
	Priority   types.Int64  `tfsdk:"priority"`
}

func (r *SystemgroupSystemcmdpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemgroup_systemcmdpolicy_binding resource.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the system group.",
			},
			"policyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of command policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the command policy.",
			},
		},
	}
}

func systemgroup_systemcmdpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SystemgroupSystemcmdpolicyBindingResourceModel) system.Systemgroupsystemcmdpolicybinding {
	tflog.Debug(ctx, "In systemgroup_systemcmdpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemgroup_systemcmdpolicy_binding := system.Systemgroupsystemcmdpolicybinding{}
	if !data.Groupname.IsNull() {
		systemgroup_systemcmdpolicy_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Policyname.IsNull() {
		systemgroup_systemcmdpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		systemgroup_systemcmdpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return systemgroup_systemcmdpolicy_binding
}

func systemgroup_systemcmdpolicy_bindingSetAttrFromGet(ctx context.Context, data *SystemgroupSystemcmdpolicyBindingResourceModel, getResponseData map[string]interface{}) *SystemgroupSystemcmdpolicyBindingResourceModel {
	tflog.Debug(ctx, "In systemgroup_systemcmdpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
