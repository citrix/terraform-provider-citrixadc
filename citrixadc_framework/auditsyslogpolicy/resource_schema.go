package auditsyslogpolicy

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

// AuditsyslogpolicyResourceModel describes the resource data model.
type AuditsyslogpolicyResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Action        types.String `tfsdk:"action"`
	Name          types.String `tfsdk:"name"`
	Rule          types.String `tfsdk:"rule"`
	Globalbinding types.Set    `tfsdk:"globalbinding"`
}

// AuditsyslogpolicyGlobalbindingModel describes a single systemglobal binding.
type AuditsyslogpolicyGlobalbindingModel struct {
	Feature                types.String `tfsdk:"feature"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Nextfactor             types.String `tfsdk:"nextfactor"`
	Priority               types.Int64  `tfsdk:"priority"`
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
				Optional:    true,
				Computed:    true,
				Description: "Syslog server action to perform when this policy matches traffic.\nNOTE: A syslog server action must be associated with a syslog audit policy.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the policy.\nMust begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the syslog policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my syslog policy\" or 'my syslog policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that defines the messages to be logged to the syslog server.",
			},
		},
		Blocks: map[string]schema.Block{
			"globalbinding": schema.SetNestedBlock{
				Description: "Inline systemglobal binding for the auditsyslog policy.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"feature": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The feature to be checked while applying this config.",
						},
						"globalbindtype": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The bind type for the global binding.",
						},
						"gotopriorityexpression": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.",
						},
						"nextfactor": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "On success invoke label. Applicable for advanced authentication policy binding.",
						},
						"priority": schema.Int64Attribute{
							Optional:    true,
							Computed:    true,
							Description: "The priority of the policy binding.",
						},
					},
				},
			},
		},
	}
}

func auditsyslogpolicyGetThePayloadFromthePlan(ctx context.Context, data *AuditsyslogpolicyResourceModel) audit.Auditsyslogpolicy {
	tflog.Debug(ctx, "In auditsyslogpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	auditsyslogpolicy := audit.Auditsyslogpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		auditsyslogpolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		auditsyslogpolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
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
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
