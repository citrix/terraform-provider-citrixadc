package rnatglobal_auditsyslogpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RnatglobalAuditsyslogpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The policy Name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the policy.",
			},
		},
	}
}
