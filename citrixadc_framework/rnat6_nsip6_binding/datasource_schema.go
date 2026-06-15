package rnat6_nsip6_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Rnat6Nsip6BindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the RNAT6 rule to which to bind NAT IPs.",
			},
			"natip6": schema.StringAttribute{
				Required:    true,
				Description: "Nat IP Address.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this rnat rule.",
			},
		},
	}
}
