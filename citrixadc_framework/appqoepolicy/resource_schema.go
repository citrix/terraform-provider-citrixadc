package appqoepolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppqoepolicyResourceModel describes the resource data model.
type AppqoepolicyResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Rule   types.String `tfsdk:"rule"`
}

func (r *AppqoepolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appqoepolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Configured AppQoE action to trigger.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the AppQoE policy. Minimum length = 1",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression or name of a named expression, against which the request is evaluated. The policy is applied if the rule evaluates to true.",
			},
		},
	}
}

func appqoepolicyGetThePayloadFromthePlan(ctx context.Context, data *AppqoepolicyResourceModel) appqoe.Appqoepolicy {
	tflog.Debug(ctx, "In appqoepolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appqoepolicy := appqoe.Appqoepolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		appqoepolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appqoepolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		appqoepolicy.Rule = data.Rule.ValueString()
	}

	return appqoepolicy
}

func appqoepolicyGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *AppqoepolicyResourceModel) appqoe.Appqoepolicy {
	tflog.Debug(ctx, "In appqoepolicyGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields.
	// Name is the primary key and must be included in the unnamed PUT payload.
	appqoepolicy := appqoe.Appqoepolicy{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appqoepolicy.Name = data.Name.ValueString()
	}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		appqoepolicy.Action = data.Action.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		appqoepolicy.Rule = data.Rule.ValueString()
	}

	return appqoepolicy
}

func appqoepolicySetAttrFromGet(ctx context.Context, data *AppqoepolicyResourceModel, getResponseData map[string]interface{}) *AppqoepolicyResourceModel {
	tflog.Debug(ctx, "In appqoepolicySetAttrFromGet Function")

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
