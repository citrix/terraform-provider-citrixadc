package authenticationpolicylabel_authenticationpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationpolicylabelAuthenticationpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authentication policy label to which to bind the policy.",
			},
			"nextfactor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On success invoke label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authentication policy to bind to the policy label.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}
