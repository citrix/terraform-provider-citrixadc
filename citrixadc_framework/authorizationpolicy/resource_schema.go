package authorizationpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authorization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthorizationpolicyResourceModel describes the resource data model.
type AuthorizationpolicyResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Action  types.String `tfsdk:"action"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Rule    types.String `tfsdk:"rule"`
}

func (r *AuthorizationpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authorizationpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action to perform if the policy matches: either allow or deny the request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new authorization policy. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authorization policy\" or 'my authorization policy').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new name of the author policy.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to perform the authentication.",
			},
		},
	}
}

func authorizationpolicyGetThePayloadFromtheConfig(ctx context.Context, data *AuthorizationpolicyResourceModel) authorization.Authorizationpolicy {
	tflog.Debug(ctx, "In authorizationpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authorizationpolicy := authorization.Authorizationpolicy{}
	if !data.Action.IsNull() {
		authorizationpolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() {
		authorizationpolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		authorizationpolicy.Newname = data.Newname.ValueString()
	}
	if !data.Rule.IsNull() {
		authorizationpolicy.Rule = data.Rule.ValueString()
	}

	return authorizationpolicy
}

func authorizationpolicySetAttrFromGet(ctx context.Context, data *AuthorizationpolicyResourceModel, getResponseData map[string]interface{}) *AuthorizationpolicyResourceModel {
	tflog.Debug(ctx, "In authorizationpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
