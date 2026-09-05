package authenticationradiuspolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationradiuspolicyResourceModel describes the resource data model.
type AuthenticationradiuspolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Reqaction types.String `tfsdk:"reqaction"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *AuthenticationradiuspolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationradiuspolicy resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				// SDK v2 marked name ForceNew; NITRO cannot rename the policy after creation.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the RADIUS authentication policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after RADIUS policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"reqaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the RADIUS action to perform if the policy matches.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the RADIUS server.",
			},
		},
	}
}

func authenticationradiuspolicyGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationradiuspolicyResourceModel) authentication.Authenticationradiuspolicy {
	tflog.Debug(ctx, "In authenticationradiuspolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationradiuspolicy := authentication.Authenticationradiuspolicy{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationradiuspolicy.Name = data.Name.ValueString()
	}
	if !data.Reqaction.IsNull() && !data.Reqaction.IsUnknown() {
		authenticationradiuspolicy.Reqaction = data.Reqaction.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		authenticationradiuspolicy.Rule = data.Rule.ValueString()
	}

	return authenticationradiuspolicy
}

func authenticationradiuspolicySetAttrFromGet(ctx context.Context, data *AuthenticationradiuspolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationradiuspolicyResourceModel {
	tflog.Debug(ctx, "In authenticationradiuspolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["reqaction"]; ok && val != nil {
		data.Reqaction = types.StringValue(val.(string))
	} else {
		data.Reqaction = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
