package aaauser

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaauserResourceModel describes the resource data model.
type AaauserResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Loggedin          types.Bool   `tfsdk:"loggedin"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Username          types.String `tfsdk:"username"`
}

func (r *AaauserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaauser resource.",
			},
			"loggedin": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Show whether the user is logged in or not.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password with which the user logs on. Required for any user account that does not exist on an external authentication server.\nIf you are not using an external authentication server, all user accounts must have a password. If you are using an external authentication server, you must provide a password for local user accounts that do not exist on the authentication server.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Password with which the user logs on. Required for any user account that does not exist on an external authentication server.\nIf you are not using an external authentication server, all user accounts must have a password. If you are using an external authentication server, you must provide a password for local user accounts that do not exist on the authentication server.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a password_wo update.",
			},
			"username": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the user. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the user is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or\nsingle quotation marks (for example, \"my aaa user\" or \"my aaa user\").",
			},
		},
	}
}

func aaauserGetThePayloadFromthePlan(ctx context.Context, data *AaauserResourceModel) aaa.Aaauser {
	tflog.Debug(ctx, "In aaauserGetThePayloadFromthePlan Function")

	// Create API request body from the model
	aaauser := aaa.Aaauser{}
	if !data.Loggedin.IsNull() {
		aaauser.Loggedin = data.Loggedin.ValueBool()
	}
	if !data.Password.IsNull() {
		aaauser.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Username.IsNull() {
		aaauser.Username = data.Username.ValueString()
	}

	return aaauser
}

func aaauserGetThePayloadFromtheConfig(ctx context.Context, data *AaauserResourceModel, payload *aaa.Aaauser) {
	tflog.Debug(ctx, "In aaauserGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func aaauserSetAttrFromGet(ctx context.Context, data *AaauserResourceModel, getResponseData map[string]interface{}) *AaauserResourceModel {
	tflog.Debug(ctx, "In aaauserSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["loggedin"]; ok && val != nil {
		data.Loggedin = types.BoolValue(val.(bool))
	} else {
		data.Loggedin = types.BoolNull()
	}
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
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
