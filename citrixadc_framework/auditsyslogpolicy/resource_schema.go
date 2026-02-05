package auditsyslogpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuditsyslogpolicyResourceModel describes the resource data model.
type AuditsyslogpolicyResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Rule   types.String `tfsdk:"rule"`
}

func (r *AuditsyslogpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the auditsyslogpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Syslog server action to perform when this policy matches traffic.\nNOTE: A syslog server action must be associated with a syslog audit policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy.\nMust begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the syslog policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my syslog policy\" or 'my syslog policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that defines the messages to be logged to the syslog server.",
			},
		},
	}
}

func auditsyslogpolicyGetThePayloadFromtheConfig(ctx context.Context, data *AuditsyslogpolicyResourceModel) audit.Auditsyslogpolicy {
	tflog.Debug(ctx, "In auditsyslogpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	auditsyslogpolicy := audit.Auditsyslogpolicy{}
	if !data.Action.IsNull() {
		auditsyslogpolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() {
		auditsyslogpolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() {
		auditsyslogpolicy.Rule = data.Rule.ValueString()
	}

	return auditsyslogpolicy
}

func auditsyslogpolicySetAttrFromGet(ctx context.Context, data *AuditsyslogpolicyResourceModel, getResponseData map[string]interface{}) *AuditsyslogpolicyResourceModel {
	tflog.Debug(ctx, "In auditsyslogpolicySetAttrFromGet Function")

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
