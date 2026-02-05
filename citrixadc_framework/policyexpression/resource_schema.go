package policyexpression

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicyexpressionResourceModel describes the resource data model.
type PolicyexpressionResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Clientsecuritymessage types.String `tfsdk:"clientsecuritymessage"`
	Comment               types.String `tfsdk:"comment"`
	Name                  types.String `tfsdk:"name"`
	Value                 types.String `tfsdk:"value"`
}

func (r *PolicyexpressionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policyexpression resource.",
			},
			"clientsecuritymessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Message to display if the expression fails. Allowed for classic end-point check expressions only.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the expression. Displayed upon viewing the policy expression.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Unique name for the expression. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or HTTP callout.",
			},
			"value": schema.StringAttribute{
				Required:    true,
				Description: "Expression string. For example: http.req.body(100).contains(\"this\").",
			},
		},
	}
}

func policyexpressionGetThePayloadFromtheConfig(ctx context.Context, data *PolicyexpressionResourceModel) policy.Policyexpression {
	tflog.Debug(ctx, "In policyexpressionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	policyexpression := policy.Policyexpression{}
	if !data.Clientsecuritymessage.IsNull() {
		policyexpression.Clientsecuritymessage = data.Clientsecuritymessage.ValueString()
	}
	if !data.Comment.IsNull() {
		policyexpression.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		policyexpression.Name = data.Name.ValueString()
	}
	if !data.Value.IsNull() {
		policyexpression.Value = data.Value.ValueString()
	}

	return policyexpression
}

func policyexpressionSetAttrFromGet(ctx context.Context, data *PolicyexpressionResourceModel, getResponseData map[string]interface{}) *PolicyexpressionResourceModel {
	tflog.Debug(ctx, "In policyexpressionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientsecuritymessage"]; ok && val != nil {
		data.Clientsecuritymessage = types.StringValue(val.(string))
	} else {
		data.Clientsecuritymessage = types.StringNull()
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
	if val, ok := getResponseData["value"]; ok && val != nil {
		data.Value = types.StringValue(val.(string))
	} else {
		data.Value = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Name.ValueString()))

	return data
}
