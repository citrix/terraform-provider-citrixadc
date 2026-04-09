package systemuser

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SystemuserResourceModel describes the resource data model.
type SystemuserResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Allowedmanagementinterface types.List   `tfsdk:"allowedmanagementinterface"`
	Externalauth               types.String `tfsdk:"externalauth"`
	Hashedpassword             types.String `tfsdk:"hashedpassword"`
	Logging                    types.String `tfsdk:"logging"`
	Maxsession                 types.Int64  `tfsdk:"maxsession"`
	Password                   types.String `tfsdk:"password"`
	PasswordWo                 types.String `tfsdk:"password_wo"`
	PasswordWoVersion          types.Int64  `tfsdk:"password_wo_version"`
	Promptstring               types.String `tfsdk:"promptstring"`
	Timeout                    types.Int64  `tfsdk:"timeout"`
	Username                   types.String `tfsdk:"username"`
	Cmdpolicybinding           types.Set    `tfsdk:"cmdpolicybinding"`
}

// CmdpolicyBindingModel describes a single cmdpolicy binding.
type CmdpolicyBindingModel struct {
	Policyname types.String `tfsdk:"policyname"`
	Priority   types.Int64  `tfsdk:"priority"`
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
				Computed:    true,
				Description: "Allowed Management interfaces to the system user. By default user is allowed from both API and CLI interfaces. If management interface for a user is set to API, then user is not allowed to access NS through CLI. GUI interface will come under API interface",
			},
			"externalauth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether to use external authentication servers for the system user authentication or not",
			},
			"hashedpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hashed password for the system user, as returned by the NITRO API. Used to detect out-of-band password changes.",
			},
			"logging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Users logging privilege",
			},
			"maxsession": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of client connection allowed per user",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password for the system user. Can include any ASCII character.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Password for the system user. Can include any ASCII character.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a password_wo update.",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for a user. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the user is added.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my user\" or 'my user').",
			},
		},
		Blocks: map[string]schema.Block{
			"cmdpolicybinding": schema.SetNestedBlock{
				Description: "Inline command policy bindings for the system user.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
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
					},
				},
			},
		},
	}
}

func systemuserGetThePayloadFromthePlan(ctx context.Context, data *SystemuserResourceModel) system.Systemuser {
	tflog.Debug(ctx, "In systemuserGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemuser := system.Systemuser{}
	if !data.Allowedmanagementinterface.IsNull() {
		var allowedmanagementinterfaceList []string
		data.Allowedmanagementinterface.ElementsAs(ctx, &allowedmanagementinterfaceList, false)
		systemuser.Allowedmanagementinterface = allowedmanagementinterfaceList
	}
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
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
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

func systemuserGetThePayloadFromtheConfig(ctx context.Context, data *SystemuserResourceModel, payload *system.Systemuser) {
	tflog.Debug(ctx, "In systemuserGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func systemuserSetAttrFromGet(ctx context.Context, data *SystemuserResourceModel, getResponseData map[string]interface{}) *SystemuserResourceModel {
	tflog.Debug(ctx, "In systemuserSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["allowedmanagementinterface"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Allowedmanagementinterface = listValue
		} else {
			data.Allowedmanagementinterface = types.ListNull(types.StringType)
		}
	} else {
		data.Allowedmanagementinterface = types.ListNull(types.StringType)
	}
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
	// The NITRO API returns the hashed password in the "password" response field.
	// Track it in hashedpassword to detect out-of-band password changes.
	if val, ok := getResponseData["password"]; ok && val != nil {
		newHashedPassword := val.(string)
		oldHashedPassword := data.Hashedpassword.ValueString()

		// If the hash changed on the ADC but the Terraform config password didn't change,
		// something external changed it — clear password in state to force drift detection.
		if oldHashedPassword != "" && oldHashedPassword != newHashedPassword {
			if !data.Password.IsNull() && data.Password.ValueString() != "" {
				// Password was set in config but hash changed externally — clear to trigger re-apply
				data.Password = types.StringNull()
			}
		}
		data.Hashedpassword = types.StringValue(newHashedPassword)
	} else {
		data.Hashedpassword = types.StringNull()
	}
	// password_wo is not returned by NITRO API (write-only/ephemeral) - retain from config
	// password_wo_version is not returned by NITRO API (version tracker) - retain from config
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
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Username.ValueString()))

	return data
}
