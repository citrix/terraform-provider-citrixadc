package systemgroup_systemuser_binding

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

// SystemgroupSystemuserBindingResourceModel describes the resource data model.
type SystemgroupSystemuserBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Groupname types.String `tfsdk:"groupname"`
	Username  types.String `tfsdk:"username"`
}

func (r *SystemgroupSystemuserBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemgroup_systemuser_binding resource.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the system group.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The system user.",
			},
		},
	}
}

func systemgroup_systemuser_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SystemgroupSystemuserBindingResourceModel) system.Systemgroupsystemuserbinding {
	tflog.Debug(ctx, "In systemgroup_systemuser_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemgroup_systemuser_binding := system.Systemgroupsystemuserbinding{}
	if !data.Groupname.IsNull() {
		systemgroup_systemuser_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Username.IsNull() {
		systemgroup_systemuser_binding.Username = data.Username.ValueString()
	}

	return systemgroup_systemuser_binding
}

func systemgroup_systemuser_bindingSetAttrFromGet(ctx context.Context, data *SystemgroupSystemuserBindingResourceModel, getResponseData map[string]interface{}) *SystemgroupSystemuserBindingResourceModel {
	tflog.Debug(ctx, "In systemgroup_systemuser_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Username.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
