package authenticationsmartaccesspolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationsmartaccesspolicyResourceModel describes the resource data model.
type AuthenticationsmartaccesspolicyResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Action  types.String `tfsdk:"action"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Rule    types.String `tfsdk:"rule"`
}

func (r *AuthenticationsmartaccesspolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationsmartaccesspolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Smartaccess profile to use if the policy matches.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this policy.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Smartaccess policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after Smartaccess policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression.",
			},
		},
	}
}

func authenticationsmartaccesspolicyGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationsmartaccesspolicyResourceModel) authentication.Authenticationsmartaccesspolicy {
	tflog.Debug(ctx, "In authenticationsmartaccesspolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationsmartaccesspolicy := authentication.Authenticationsmartaccesspolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		authenticationsmartaccesspolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		authenticationsmartaccesspolicy.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationsmartaccesspolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		authenticationsmartaccesspolicy.Rule = data.Rule.ValueString()
	}

	return authenticationsmartaccesspolicy
}

func authenticationsmartaccesspolicySetAttrFromGet(ctx context.Context, data *AuthenticationsmartaccesspolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationsmartaccesspolicyResourceModel {
	tflog.Debug(ctx, "In authenticationsmartaccesspolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
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
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
