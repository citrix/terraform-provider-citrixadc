package authenticationnoauthaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationnoauthactionResourceModel describes the resource data model.
type AuthenticationnoauthactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Name                       types.String `tfsdk:"name"`
}

func (r *AuthenticationnoauthactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationnoauthaction resource.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the group that is added to user sessions that match current policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new no-authentication action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
		},
	}
}

func authenticationnoauthactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationnoauthactionResourceModel) authentication.Authenticationnoauthaction {
	tflog.Debug(ctx, "In authenticationnoauthactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationnoauthaction := authentication.Authenticationnoauthaction{}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationnoauthaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationnoauthaction.Name = data.Name.ValueString()
	}

	return authenticationnoauthaction
}

func authenticationnoauthactionSetAttrFromGet(ctx context.Context, data *AuthenticationnoauthactionResourceModel, getResponseData map[string]interface{}) *AuthenticationnoauthactionResourceModel {
	tflog.Debug(ctx, "In authenticationnoauthactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
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
