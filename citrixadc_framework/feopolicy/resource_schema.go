package feopolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/feo"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
				Required:    true,
				Description: "The name of the front end optimization policy.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "The rule associated with the front end optimization policy.",
			},
		},
	}
}

func feopolicyGetThePayloadFromtheConfig(ctx context.Context, data *FeopolicyResourceModel) feo.Feopolicy {
	tflog.Debug(ctx, "In feopolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	feopolicy := feo.Feopolicy{}
	if !data.Action.IsNull() {
		feopolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() {
		feopolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() {
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
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
