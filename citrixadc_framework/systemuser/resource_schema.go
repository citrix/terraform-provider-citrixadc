package systemuser

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SystemuserResourceModel describes the resource data model.
type SystemuserResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Allowedmanagementinterface types.List   `tfsdk:"allowedmanagementinterface"`
	Externalauth               types.String `tfsdk:"externalauth"`
	Logging                    types.String `tfsdk:"logging"`
	Maxsession                 types.Int64  `tfsdk:"maxsession"`
	Password                   types.String `tfsdk:"password"`
	Promptstring               types.String `tfsdk:"promptstring"`
	Timeout                    types.Int64  `tfsdk:"timeout"`
	Username                   types.String `tfsdk:"username"`
}

func (r *SystemuserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemuser resource.",
			},
			"allowedmanagementinterface": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Allowed Management interfaces to the system user. By default user is allowed from both API and CLI interfaces. If management interface for a user is set to API, then user is not allowed to access NS through CLI. GUI interface will come under API interface",
			},
			"externalauth": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Whether to use external authentication servers for the system user authentication or not",
			},
			"logging": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Users logging privilege",
			},
			"maxsession": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of client connection allowed per user",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Description: "Password for the system user. Can include any ASCII character.",
			},
			"promptstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:\n* %u - Will be replaced by the user name.\n* %h - Will be replaced by the hostname of the Citrix ADC.\n* %t - Will be replaced by the current time in 12-hour format.\n* %T - Will be replaced by the current time in 24-hour format.\n* %d - Will be replaced by the current date.\n* %s - Will be replaced by the state of the Citrix ADC.\n\nNote: The 63-character limit for the length of the string does not apply to the characters that replace the variables.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds. If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name for a user. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the user is added.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my user\" or 'my user').",
			},
		},
	}
}

func systemuserGetThePayloadFromtheConfig(ctx context.Context, data *SystemuserResourceModel) system.Systemuser {
	tflog.Debug(ctx, "In systemuserGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemuser := system.Systemuser{}
	if !data.Externalauth.IsNull() {
		systemuser.Externalauth = data.Externalauth.ValueString()
	}
	if !data.Logging.IsNull() {
		systemuser.Logging = data.Logging.ValueString()
	}
	if !data.Maxsession.IsNull() {
		systemuser.Maxsession = utils.IntPtr(int(data.Maxsession.ValueInt64()))
	}
	if !data.Password.IsNull() {
		systemuser.Password = data.Password.ValueString()
	}
	if !data.Promptstring.IsNull() {
		systemuser.Promptstring = data.Promptstring.ValueString()
	}
	if !data.Timeout.IsNull() {
		systemuser.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Username.IsNull() {
		systemuser.Username = data.Username.ValueString()
	}

	return systemuser
}

func systemuserSetAttrFromGet(ctx context.Context, data *SystemuserResourceModel, getResponseData map[string]interface{}) *SystemuserResourceModel {
	tflog.Debug(ctx, "In systemuserSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["externalauth"]; ok && val != nil {
		data.Externalauth = types.StringValue(val.(string))
	} else {
		data.Externalauth = types.StringNull()
	}
	if val, ok := getResponseData["logging"]; ok && val != nil {
		data.Logging = types.StringValue(val.(string))
	} else {
		data.Logging = types.StringNull()
	}
	if val, ok := getResponseData["maxsession"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxsession = types.Int64Value(intVal)
		}
	} else {
		data.Maxsession = types.Int64Null()
	}
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
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
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Username.ValueString())

	return data
}
