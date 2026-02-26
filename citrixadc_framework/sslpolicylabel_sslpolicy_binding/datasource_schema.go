package sslpolicylabel_sslpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslpolicylabelSslpolicyBindingDataSourceSchema() schema.Schema {
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
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Invoke policies bound to a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.",
			},
			"invoke_labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the label to invoke if the current policy rule evaluates to TRUE.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL policy label to which to bind policies.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label invocation.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL policy to bind to the policy label.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}
