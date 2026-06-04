package aaaglobal_authenticationnegotiateaction_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaaglobalAuthenticationnegotiateactionBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"windowsprofile": schema.StringAttribute{
				Required:    true,
				Description: "Name of the negotiate profile to be bound.",
			},
		},
	}
}