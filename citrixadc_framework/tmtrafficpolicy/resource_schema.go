package tmtrafficpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TmtrafficpolicyResourceModel describes the resource data model.
type TmtrafficpolicyResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Rule   types.String `tfsdk:"rule"`
}

func (r *TmtrafficpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the tmtrafficpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the action to apply to requests or connections that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the traffic policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named expression, or an expression, that the policy uses to determine whether to apply certain action on the current traffic.",
			},
		},
	}
}

func tmtrafficpolicyGetThePayloadFromtheConfig(ctx context.Context, data *TmtrafficpolicyResourceModel) tm.Tmtrafficpolicy {
	tflog.Debug(ctx, "In tmtrafficpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	tmtrafficpolicy := tm.Tmtrafficpolicy{}
	if !data.Action.IsNull() {
		tmtrafficpolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() {
		tmtrafficpolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() {
		tmtrafficpolicy.Rule = data.Rule.ValueString()
	}

	return tmtrafficpolicy
}

func tmtrafficpolicySetAttrFromGet(ctx context.Context, data *TmtrafficpolicyResourceModel, getResponseData map[string]interface{}) *TmtrafficpolicyResourceModel {
	tflog.Debug(ctx, "In tmtrafficpolicySetAttrFromGet Function")

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
