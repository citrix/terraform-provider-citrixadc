package authorizationpolicy

import (
	"context"
	"fmt"

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
				// SDK v2 parity: Required (mutable via NITRO update).
				Required:    true,
				Description: "Action to perform if the policy matches: either allow or deny the request.",
			},
			"name": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew -> RequiresReplace.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the new authorization policy. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authorization policy\" or 'my authorization policy').",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it must
				// NOT force replacement - it drives an in-place rename via Update. Not
				// Computed: it is a pure user input, never echoed back by GET.
				Optional:    true,
				Description: "The new name of the author policy.",
			},
			"rule": schema.StringAttribute{
				// SDK v2 parity: Required (mutable via NITRO update).
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to perform the authentication.",
			},
		},
	}
}

func authorizationpolicyGetThePayloadFromthePlan(ctx context.Context, data *AuthorizationpolicyResourceModel) authorization.Authorizationpolicy {
	tflog.Debug(ctx, "In authorizationpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model (add/create POST)
	authorizationpolicy := authorization.Authorizationpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		authorizationpolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authorizationpolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of the
	// add payload, so it is deliberately excluded from the create POST body.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		authorizationpolicy.Rule = data.Rule.ValueString()
	}

	return authorizationpolicy
}

func authorizationpolicyGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *AuthorizationpolicyResourceModel) authorization.Authorizationpolicy {
	tflog.Debug(ctx, "In authorizationpolicyGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields.
	authorizationpolicy := authorization.Authorizationpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		authorizationpolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authorizationpolicy.Name = data.Name.ValueString()
	}
	// newname is rename-only; excluded from the update PUT body.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		authorizationpolicy.Rule = data.Rule.ValueString()
	}

	return authorizationpolicy
}

func authorizationpolicySetAttrFromGet(ctx context.Context, data *AuthorizationpolicyResourceModel, getResponseData map[string]interface{}) *AuthorizationpolicyResourceModel {
	tflog.Debug(ctx, "In authorizationpolicySetAttrFromGet Function")

	// Convert API response to model.
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	}
	// name is the user-facing key. Once a rename has happened (via newname), the live
	// object name (tracked by data.Id) diverges from the configured name, and GET
	// returns the live (new) name. Overwriting name from GET would clobber the user's
	// configured value and trigger a spurious RequiresReplace diff. So only adopt the
	// GET value when we don't already have one (e.g. on import, where state carries
	// only the ID); otherwise preserve.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	}

	return data
}

// authorizationpolicySetAttrFromGetForDatasource faithfully copies every field from
// the GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func authorizationpolicySetAttrFromGetForDatasource(ctx context.Context, data *AuthorizationpolicyResourceModel, getResponseData map[string]interface{}) *AuthorizationpolicyResourceModel {
	tflog.Debug(ctx, "In authorizationpolicySetAttrFromGetForDatasource Function")

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

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
