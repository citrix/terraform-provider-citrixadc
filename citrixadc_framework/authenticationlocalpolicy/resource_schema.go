package authenticationlocalpolicy

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

// AuthenticationlocalpolicyResourceModel describes the resource data model.
type AuthenticationlocalpolicyResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Rule types.String `tfsdk:"rule"`
}

func (r *AuthenticationlocalpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationlocalpolicy resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the local authentication policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after local policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to perform the authentication.",
			},
		},
	}
}

func authenticationlocalpolicyGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationlocalpolicyResourceModel) authentication.Authenticationlocalpolicy {
	tflog.Debug(ctx, "In authenticationlocalpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationlocalpolicy := authentication.Authenticationlocalpolicy{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationlocalpolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		authenticationlocalpolicy.Rule = data.Rule.ValueString()
	}

	return authenticationlocalpolicy
}

func authenticationlocalpolicySetAttrFromGet(ctx context.Context, data *AuthenticationlocalpolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationlocalpolicyResourceModel {
	tflog.Debug(ctx, "In authenticationlocalpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
