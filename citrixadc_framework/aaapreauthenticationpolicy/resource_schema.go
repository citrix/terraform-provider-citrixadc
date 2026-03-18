package aaapreauthenticationpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
				Required:    true,
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

func aaapreauthenticationpolicyGetThePayloadFromtheConfig(ctx context.Context, data *AaapreauthenticationpolicyResourceModel) aaa.Aaapreauthenticationpolicy {
	tflog.Debug(ctx, "In aaapreauthenticationpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaapreauthenticationpolicy := aaa.Aaapreauthenticationpolicy{}
	if !data.Name.IsNull() {
		aaapreauthenticationpolicy.Name = data.Name.ValueString()
	}
	if !data.Reqaction.IsNull() {
		aaapreauthenticationpolicy.Reqaction = data.Reqaction.ValueString()
	}
	if !data.Rule.IsNull() {
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
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
