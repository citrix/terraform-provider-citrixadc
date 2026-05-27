package authenticationadfsproxyprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationadfsproxyprofileResourceModel describes the resource data model.
type AuthenticationadfsproxyprofileResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Certkeyname       types.String `tfsdk:"certkeyname"`
	Name              types.String `tfsdk:"name"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Serverurl         types.String `tfsdk:"serverurl"`
	Username          types.String `tfsdk:"username"`
}

func (r *AuthenticationadfsproxyprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationadfsproxyprofile resource.",
			},
			"certkeyname": schema.StringAttribute{
				Required:    true,
				Description: "SSL certificate of the proxy that is registered at adfs server for trust.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the adfs proxy profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my push service\" or 'my push service').",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "This is the password of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "This is the password of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a password_wo update.",
			},
			"serverurl": schema.StringAttribute{
				Required:    true,
				Description: "Fully qualified url of the adfs server.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "This is the name of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.",
			},
		},
	}
}

func authenticationadfsproxyprofileGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationadfsproxyprofileResourceModel) authentication.Authenticationadfsproxyprofile {
	tflog.Debug(ctx, "In authenticationadfsproxyprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationadfsproxyprofile := authentication.Authenticationadfsproxyprofile{}
	if !data.Certkeyname.IsNull() && !data.Certkeyname.IsUnknown() {
		authenticationadfsproxyprofile.Certkeyname = data.Certkeyname.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationadfsproxyprofile.Name = data.Name.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		authenticationadfsproxyprofile.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Serverurl.IsNull() && !data.Serverurl.IsUnknown() {
		authenticationadfsproxyprofile.Serverurl = data.Serverurl.ValueString()
	}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		authenticationadfsproxyprofile.Username = data.Username.ValueString()
	}

	return authenticationadfsproxyprofile
}

func authenticationadfsproxyprofileGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationadfsproxyprofileResourceModel, payload *authentication.Authenticationadfsproxyprofile) {
	tflog.Debug(ctx, "In authenticationadfsproxyprofileGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func authenticationadfsproxyprofileSetAttrFromGet(ctx context.Context, data *AuthenticationadfsproxyprofileResourceModel, getResponseData map[string]interface{}) *AuthenticationadfsproxyprofileResourceModel {
	tflog.Debug(ctx, "In authenticationadfsproxyprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["certkeyname"]; ok && val != nil {
		data.Certkeyname = types.StringValue(val.(string))
	} else {
		data.Certkeyname = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["serverurl"]; ok && val != nil {
		data.Serverurl = types.StringValue(val.(string))
	} else {
		data.Serverurl = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
