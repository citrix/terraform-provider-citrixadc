package authenticationcitrixauthaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationcitrixauthactionResourceModel describes the resource data model.
type AuthenticationcitrixauthactionResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Authentication     types.String `tfsdk:"authentication"`
	Authenticationtype types.String `tfsdk:"authenticationtype"`
	Name               types.String `tfsdk:"name"`
}

func (r *AuthenticationcitrixauthactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationcitrixauthaction resource.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Authentication needs to be disabled for searching user object without performing authentication.",
			},
			"authenticationtype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("CITRIXCONNECTOR"),
				Description: "Type of the Citrix Authentication implementation. Default implementation uses Citrix Cloud Connector.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new Citrix Authentication action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
		},
	}
}

func authenticationcitrixauthactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationcitrixauthactionResourceModel) authentication.Authenticationcitrixauthaction {
	tflog.Debug(ctx, "In authenticationcitrixauthactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationcitrixauthaction := authentication.Authenticationcitrixauthaction{}
	if !data.Authentication.IsNull() {
		authenticationcitrixauthaction.Authentication = data.Authentication.ValueString()
	}
	if !data.Authenticationtype.IsNull() {
		authenticationcitrixauthaction.Authenticationtype = data.Authenticationtype.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationcitrixauthaction.Name = data.Name.ValueString()
	}

	return authenticationcitrixauthaction
}

func authenticationcitrixauthactionSetAttrFromGet(ctx context.Context, data *AuthenticationcitrixauthactionResourceModel, getResponseData map[string]interface{}) *AuthenticationcitrixauthactionResourceModel {
	tflog.Debug(ctx, "In authenticationcitrixauthactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["authenticationtype"]; ok && val != nil {
		data.Authenticationtype = types.StringValue(val.(string))
	} else {
		data.Authenticationtype = types.StringNull()
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
