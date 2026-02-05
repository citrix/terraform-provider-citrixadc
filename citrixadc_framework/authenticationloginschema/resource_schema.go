package authenticationloginschema

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationloginschemaResourceModel describes the resource data model.
type AuthenticationloginschemaResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Authenticationschema    types.String `tfsdk:"authenticationschema"`
	Authenticationstrength  types.Int64  `tfsdk:"authenticationstrength"`
	Name                    types.String `tfsdk:"name"`
	Passwdexpression        types.String `tfsdk:"passwdexpression"`
	Passwordcredentialindex types.Int64  `tfsdk:"passwordcredentialindex"`
	Ssocredentials          types.String `tfsdk:"ssocredentials"`
	Usercredentialindex     types.Int64  `tfsdk:"usercredentialindex"`
	Userexpression          types.String `tfsdk:"userexpression"`
}

func (r *AuthenticationloginschemaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationloginschema resource.",
			},
			"authenticationschema": schema.StringAttribute{
				Required:    true,
				Description: "Name of the file for reading authentication schema to be sent for Login Page UI. This file should contain xml definition of elements as per Citrix Forms Authentication Protocol to be able to render login form. If administrator does not want to prompt users for additional credentials but continue with previously obtained credentials, then \"noschema\" can be given as argument. Please note that this applies only to loginSchemas that are used with user-defined factors, and not the vserver factor.",
			},
			"authenticationstrength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight of the current authentication",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new login schema. Login schema defines the way login form is rendered. It provides a way to customize the fields that are shown to the user. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"passwdexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression for password extraction during login. This can be any relevant advanced policy expression.",
			},
			"passwordcredentialindex": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The index at which user entered password should be stored in session.",
			},
			"ssocredentials": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option indicates whether current factor credentials are the default SSO (SingleSignOn) credentials.",
			},
			"usercredentialindex": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The index at which user entered username should be stored in session.",
			},
			"userexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression for username extraction during login. This can be any relevant advanced policy expression.",
			},
		},
	}
}

func authenticationloginschemaGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationloginschemaResourceModel) authentication.Authenticationloginschema {
	tflog.Debug(ctx, "In authenticationloginschemaGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationloginschema := authentication.Authenticationloginschema{}
	if !data.Authenticationschema.IsNull() {
		authenticationloginschema.Authenticationschema = data.Authenticationschema.ValueString()
	}
	if !data.Authenticationstrength.IsNull() {
		authenticationloginschema.Authenticationstrength = utils.IntPtr(int(data.Authenticationstrength.ValueInt64()))
	}
	if !data.Name.IsNull() {
		authenticationloginschema.Name = data.Name.ValueString()
	}
	if !data.Passwdexpression.IsNull() {
		authenticationloginschema.Passwdexpression = data.Passwdexpression.ValueString()
	}
	if !data.Passwordcredentialindex.IsNull() {
		authenticationloginschema.Passwordcredentialindex = utils.IntPtr(int(data.Passwordcredentialindex.ValueInt64()))
	}
	if !data.Ssocredentials.IsNull() {
		authenticationloginschema.Ssocredentials = data.Ssocredentials.ValueString()
	}
	if !data.Usercredentialindex.IsNull() {
		authenticationloginschema.Usercredentialindex = utils.IntPtr(int(data.Usercredentialindex.ValueInt64()))
	}
	if !data.Userexpression.IsNull() {
		authenticationloginschema.Userexpression = data.Userexpression.ValueString()
	}

	return authenticationloginschema
}

func authenticationloginschemaSetAttrFromGet(ctx context.Context, data *AuthenticationloginschemaResourceModel, getResponseData map[string]interface{}) *AuthenticationloginschemaResourceModel {
	tflog.Debug(ctx, "In authenticationloginschemaSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["authenticationschema"]; ok && val != nil {
		data.Authenticationschema = types.StringValue(val.(string))
	} else {
		data.Authenticationschema = types.StringNull()
	}
	if val, ok := getResponseData["authenticationstrength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Authenticationstrength = types.Int64Value(intVal)
		}
	} else {
		data.Authenticationstrength = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["passwdexpression"]; ok && val != nil {
		data.Passwdexpression = types.StringValue(val.(string))
	} else {
		data.Passwdexpression = types.StringNull()
	}
	if val, ok := getResponseData["passwordcredentialindex"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Passwordcredentialindex = types.Int64Value(intVal)
		}
	} else {
		data.Passwordcredentialindex = types.Int64Null()
	}
	if val, ok := getResponseData["ssocredentials"]; ok && val != nil {
		data.Ssocredentials = types.StringValue(val.(string))
	} else {
		data.Ssocredentials = types.StringNull()
	}
	if val, ok := getResponseData["usercredentialindex"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Usercredentialindex = types.Int64Value(intVal)
		}
	} else {
		data.Usercredentialindex = types.Int64Null()
	}
	if val, ok := getResponseData["userexpression"]; ok && val != nil {
		data.Userexpression = types.StringValue(val.(string))
	} else {
		data.Userexpression = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
