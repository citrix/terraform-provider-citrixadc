package aaapreauthenticationpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaapreauthenticationpolicyResourceModel describes the resource data model.
type AaapreauthenticationpolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Reqaction types.String `tfsdk:"reqaction"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *AaapreauthenticationpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaapreauthenticationpolicy resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the preauthentication policy. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the preauthentication policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"reqaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the action that the policy is to invoke when a connection matches the policy.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, defining connections that match the policy.",
			},
		},
	}
}

func aaapreauthenticationpolicyGetThePayloadFromthePlan(ctx context.Context, data *AaapreauthenticationpolicyResourceModel) aaa.Aaapreauthenticationpolicy {
	tflog.Debug(ctx, "In aaapreauthenticationpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	aaapreauthenticationpolicy := aaa.Aaapreauthenticationpolicy{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		aaapreauthenticationpolicy.Name = data.Name.ValueString()
	}
	if !data.Reqaction.IsNull() && !data.Reqaction.IsUnknown() {
		aaapreauthenticationpolicy.Reqaction = data.Reqaction.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		aaapreauthenticationpolicy.Rule = data.Rule.ValueString()
	}

	return aaapreauthenticationpolicy
}

func aaapreauthenticationpolicyGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *AaapreauthenticationpolicyResourceModel) aaa.Aaapreauthenticationpolicy {
	tflog.Debug(ctx, "In aaapreauthenticationpolicyGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields
	aaapreauthenticationpolicy := aaa.Aaapreauthenticationpolicy{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		aaapreauthenticationpolicy.Name = data.Name.ValueString()
	}
	if !data.Reqaction.IsNull() && !data.Reqaction.IsUnknown() {
		aaapreauthenticationpolicy.Reqaction = data.Reqaction.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		aaapreauthenticationpolicy.Rule = data.Rule.ValueString()
	}

	return aaapreauthenticationpolicy
}

func aaapreauthenticationpolicySetAttrFromGet(ctx context.Context, data *AaapreauthenticationpolicyResourceModel, getResponseData map[string]interface{}) *AaapreauthenticationpolicyResourceModel {
	tflog.Debug(ctx, "In aaapreauthenticationpolicySetAttrFromGet Function")

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
