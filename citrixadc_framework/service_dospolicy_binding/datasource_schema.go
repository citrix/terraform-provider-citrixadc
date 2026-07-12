package service_dospolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ServiceDospolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the service to which to bind a policy or monitor.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the policy bound to the specified service.",
			},
		},
	}
}
