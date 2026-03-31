package systemuser_nspartition_binding

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

// SystemuserNspartitionBindingResourceModel describes the resource data model.
type SystemuserNspartitionBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Partitionname types.String `tfsdk:"partitionname"`
	Username      types.String `tfsdk:"username"`
}

func (r *SystemuserNspartitionBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemuser_nspartition_binding resource.",
			},
			"partitionname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Partition to bind to the system user.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of the system-user entry to which to bind the command policy.",
			},
		},
	}
}

func systemuser_nspartition_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SystemuserNspartitionBindingResourceModel) system.Systemusernspartitionbinding {
	tflog.Debug(ctx, "In systemuser_nspartition_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemuser_nspartition_binding := system.Systemusernspartitionbinding{}
	if !data.Partitionname.IsNull() {
		systemuser_nspartition_binding.Partitionname = data.Partitionname.ValueString()
	}
	if !data.Username.IsNull() {
		systemuser_nspartition_binding.Username = data.Username.ValueString()
	}

	return systemuser_nspartition_binding
}

func systemuser_nspartition_bindingSetAttrFromGet(ctx context.Context, data *SystemuserNspartitionBindingResourceModel, getResponseData map[string]interface{}) *SystemuserNspartitionBindingResourceModel {
	tflog.Debug(ctx, "In systemuser_nspartition_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["partitionname"]; ok && val != nil {
		data.Partitionname = types.StringValue(val.(string))
	} else {
		data.Partitionname = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("partitionname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Partitionname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Username.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
