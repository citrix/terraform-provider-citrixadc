package systemgroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SystemgroupResourceModel describes the resource data model.
type SystemgroupResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Allowedmanagementinterface types.List   `tfsdk:"allowedmanagementinterface"`
	Daystoexpire               types.Int64  `tfsdk:"daystoexpire"`
	Groupname                  types.String `tfsdk:"groupname"`
	Promptstring               types.String `tfsdk:"promptstring"`
	Timeout                    types.Int64  `tfsdk:"timeout"`
	Warnpriorndays             types.Int64  `tfsdk:"warnpriorndays"`
}

func (r *SystemgroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemgroup resource.",
			},
			"allowedmanagementinterface": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Allowed Management interfaces of the system users in the group. By default allowed from both API and CLI interfaces. If management interface for a group is set to API, then all users under this group will not allowed to access NS through CLI. GUI interface will come under API interface",
			},
			"daystoexpire": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Password days to expire for system groups. The daystoexpire value ranges from 30 to 255.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the group. Must begin with a letter, number, hash(#) or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the group is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my group\" or 'my group').",
			},
			"promptstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:\n* %u - Will be replaced by the user name.\n* %h - Will be replaced by the hostname of the Citrix ADC.\n* %t - Will be replaced by the current time in 12-hour format.\n* %T - Will be replaced by the current time in 24-hour format.\n* %d - Will be replaced by the current date.\n* %s - Will be replaced by the state of the Citrix ADC.\n\nNote: The 63-character limit for the length of the string does not apply to the characters that replace the variables.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds.If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.",
			},
			"warnpriorndays": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of days before which password expiration warning would be thrown with respect to daystoexpire. The warnpriorndays value ranges from 5 to 40.",
			},
		},
	}
}

func systemgroupGetThePayloadFromtheConfig(ctx context.Context, data *SystemgroupResourceModel) system.Systemgroup {
	tflog.Debug(ctx, "In systemgroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemgroup := system.Systemgroup{}
	if !data.Daystoexpire.IsNull() {
		systemgroup.Daystoexpire = utils.IntPtr(int(data.Daystoexpire.ValueInt64()))
	}
	if !data.Groupname.IsNull() {
		systemgroup.Groupname = data.Groupname.ValueString()
	}
	if !data.Promptstring.IsNull() {
		systemgroup.Promptstring = data.Promptstring.ValueString()
	}
	if !data.Timeout.IsNull() {
		systemgroup.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Warnpriorndays.IsNull() {
		systemgroup.Warnpriorndays = utils.IntPtr(int(data.Warnpriorndays.ValueInt64()))
	}

	return systemgroup
}

func systemgroupSetAttrFromGet(ctx context.Context, data *SystemgroupResourceModel, getResponseData map[string]interface{}) *SystemgroupResourceModel {
	tflog.Debug(ctx, "In systemgroupSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["daystoexpire"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Daystoexpire = types.Int64Value(intVal)
		}
	} else {
		data.Daystoexpire = types.Int64Null()
	}
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["promptstring"]; ok && val != nil {
		data.Promptstring = types.StringValue(val.(string))
	} else {
		data.Promptstring = types.StringNull()
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["warnpriorndays"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Warnpriorndays = types.Int64Value(intVal)
		}
	} else {
		data.Warnpriorndays = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Groupname.ValueString())

	return data
}
