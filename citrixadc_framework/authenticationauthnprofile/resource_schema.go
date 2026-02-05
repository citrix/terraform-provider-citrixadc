package authenticationauthnprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationauthnprofileResourceModel describes the resource data model.
type AuthenticationauthnprofileResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Authenticationdomain types.String `tfsdk:"authenticationdomain"`
	Authenticationhost   types.String `tfsdk:"authenticationhost"`
	Authenticationlevel  types.Int64  `tfsdk:"authenticationlevel"`
	Authnvsname          types.String `tfsdk:"authnvsname"`
	Name                 types.String `tfsdk:"name"`
}

func (r *AuthenticationauthnprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationauthnprofile resource.",
			},
			"authenticationdomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain for which TM cookie must to be set. If unspecified, cookie will be set for FQDN.",
			},
			"authenticationhost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hostname of the authentication vserver to which user must be redirected for authentication.",
			},
			"authenticationlevel": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication weight or level of the vserver to which this will bound. This is used to order TM vservers based on the protection required. A session that is created by authenticating against TM vserver at given level cannot be used to access TM vserver at a higher level.",
			},
			"authnvsname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authentication vserver at which authentication should be done.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the authentication profile.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the RADIUS action is added.",
			},
		},
	}
}

func authenticationauthnprofileGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationauthnprofileResourceModel) authentication.Authenticationauthnprofile {
	tflog.Debug(ctx, "In authenticationauthnprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationauthnprofile := authentication.Authenticationauthnprofile{}
	if !data.Authenticationdomain.IsNull() {
		authenticationauthnprofile.Authenticationdomain = data.Authenticationdomain.ValueString()
	}
	if !data.Authenticationhost.IsNull() {
		authenticationauthnprofile.Authenticationhost = data.Authenticationhost.ValueString()
	}
	if !data.Authenticationlevel.IsNull() {
		authenticationauthnprofile.Authenticationlevel = utils.IntPtr(int(data.Authenticationlevel.ValueInt64()))
	}
	if !data.Authnvsname.IsNull() {
		authenticationauthnprofile.Authnvsname = data.Authnvsname.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationauthnprofile.Name = data.Name.ValueString()
	}

	return authenticationauthnprofile
}

func authenticationauthnprofileSetAttrFromGet(ctx context.Context, data *AuthenticationauthnprofileResourceModel, getResponseData map[string]interface{}) *AuthenticationauthnprofileResourceModel {
	tflog.Debug(ctx, "In authenticationauthnprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["authenticationdomain"]; ok && val != nil {
		data.Authenticationdomain = types.StringValue(val.(string))
	} else {
		data.Authenticationdomain = types.StringNull()
	}
	if val, ok := getResponseData["authenticationhost"]; ok && val != nil {
		data.Authenticationhost = types.StringValue(val.(string))
	} else {
		data.Authenticationhost = types.StringNull()
	}
	if val, ok := getResponseData["authenticationlevel"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Authenticationlevel = types.Int64Value(intVal)
		}
	} else {
		data.Authenticationlevel = types.Int64Null()
	}
	if val, ok := getResponseData["authnvsname"]; ok && val != nil {
		data.Authnvsname = types.StringValue(val.(string))
	} else {
		data.Authnvsname = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
