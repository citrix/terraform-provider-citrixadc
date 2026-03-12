package systemcmdpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemcmdpolicyResourceModel describes the resource data model.
type SystemcmdpolicyResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Action     types.String `tfsdk:"action"`
	Cmdspec    types.String `tfsdk:"cmdspec"`
	Policyname types.String `tfsdk:"policyname"`
}

func (r *SystemcmdpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemcmdpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action to perform when a request matches the policy.",
			},
			"cmdspec": schema.StringAttribute{
				Required:    true,
				Description: "Regular expression specifying the data that matches the policy.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for a command policy. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the policy is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
		},
	}
}

func systemcmdpolicyGetThePayloadFromtheConfig(ctx context.Context, data *SystemcmdpolicyResourceModel) system.Systemcmdpolicy {
	tflog.Debug(ctx, "In systemcmdpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemcmdpolicy := system.Systemcmdpolicy{}
	if !data.Action.IsNull() {
		systemcmdpolicy.Action = data.Action.ValueString()
	}
	if !data.Cmdspec.IsNull() {
		systemcmdpolicy.Cmdspec = data.Cmdspec.ValueString()
	}
	if !data.Policyname.IsNull() {
		systemcmdpolicy.Policyname = data.Policyname.ValueString()
	}

	return systemcmdpolicy
}

func systemcmdpolicySetAttrFromGet(ctx context.Context, data *SystemcmdpolicyResourceModel, getResponseData map[string]interface{}) *SystemcmdpolicyResourceModel {
	tflog.Debug(ctx, "In systemcmdpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["cmdspec"]; ok && val != nil {
		data.Cmdspec = types.StringValue(val.(string))
	} else {
		data.Cmdspec = types.StringNull()
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Policyname.ValueString())

	return data
}
