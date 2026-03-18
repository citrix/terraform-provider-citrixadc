package aaaglobal_aaapreauthenticationpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaaglobalAaapreauthenticationpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "Name of the policy to be unbound.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the bound policy",
			},
		},
	}
}
