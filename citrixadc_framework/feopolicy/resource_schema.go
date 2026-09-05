package feopolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/feo"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// FeopolicyResourceModel describes the resource data model.
type FeopolicyResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Rule   types.String `tfsdk:"rule"`
}

func (r *FeopolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the feopolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "The front end optimization action that has to be performed when the rule matches.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the front end optimization policy.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "The rule associated with the front end optimization policy.",
			},
		},
	}
}

func feopolicyGetThePayloadFromthePlan(ctx context.Context, data *FeopolicyResourceModel) feo.Feopolicy {
	tflog.Debug(ctx, "In feopolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	feopolicy := feo.Feopolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		feopolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		feopolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		feopolicy.Rule = data.Rule.ValueString()
	}

	return feopolicy
}

func feopolicySetAttrFromGet(ctx context.Context, data *FeopolicyResourceModel, getResponseData map[string]interface{}) *FeopolicyResourceModel {
	tflog.Debug(ctx, "In feopolicySetAttrFromGet Function")

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
