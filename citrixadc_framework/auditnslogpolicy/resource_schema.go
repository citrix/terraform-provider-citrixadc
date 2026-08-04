package auditnslogpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuditnslogpolicyResourceModel describes the resource data model.
type AuditnslogpolicyResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Rule   types.String `tfsdk:"rule"`
}

func (r *AuditnslogpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the auditnslogpolicy resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Nslog server action that is performed when this policy matches.\nNOTE: An nslog server action must be associated with an nslog audit policy.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the policy.\nMust begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the nslog policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my nslog policy\" or 'my nslog policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that defines the messages to be logged to the nslog server.",
			},
		},
	}
}

func auditnslogpolicyGetThePayloadFromthePlan(ctx context.Context, data *AuditnslogpolicyResourceModel) audit.Auditnslogpolicy {
	tflog.Debug(ctx, "In auditnslogpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	auditnslogpolicy := audit.Auditnslogpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		auditnslogpolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		auditnslogpolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		auditnslogpolicy.Rule = data.Rule.ValueString()
	}

	return auditnslogpolicy
}

func auditnslogpolicySetAttrFromGet(ctx context.Context, data *AuditnslogpolicyResourceModel, getResponseData map[string]interface{}) *AuditnslogpolicyResourceModel {
	tflog.Debug(ctx, "In auditnslogpolicySetAttrFromGet Function")

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
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
