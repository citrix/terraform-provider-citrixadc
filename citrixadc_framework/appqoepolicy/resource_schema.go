package appqoepolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
				Description: "Configured AppQoE action to trigger",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "0",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression or name of a named expression, against which the request is evaluated. The policy is applied if the rule evaluates to true.",
			},
		},
	}
}

func appqoepolicyGetThePayloadFromtheConfig(ctx context.Context, data *AppqoepolicyResourceModel) appqoe.Appqoepolicy {
	tflog.Debug(ctx, "In appqoepolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appqoepolicy := appqoe.Appqoepolicy{}
	if !data.Action.IsNull() {
		appqoepolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() {
		appqoepolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() {
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
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
