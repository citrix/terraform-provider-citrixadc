package auditsyslogpolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuditsyslogpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Syslog server action to perform when this policy matches traffic.\nNOTE: A syslog server action must be associated with a syslog audit policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
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
