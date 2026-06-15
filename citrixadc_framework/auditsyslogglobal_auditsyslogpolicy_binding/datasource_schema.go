package auditsyslogglobal_auditsyslogpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AuditsyslogglobalAuditsyslogpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"builtin": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"feature": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The feature to be checked while applying this config",
			},
			"globalbindtype": schema.StringAttribute{
				Required:    true,
				Description: "0",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the audit syslog policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}
