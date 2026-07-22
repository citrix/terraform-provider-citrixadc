package rnat_retainsourceportset_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RnatRetainsourceportsetBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the RNAT rule to which to bind NAT IPs.",
			},
			"retainsourceportrange": schema.StringAttribute{
				Required:    true,
				Description: "When the source port range is configured and associated with the RNAT rule, Citrix ADC will choose a port from the specified source port range configured for connection establishment at the backend servers.",
			},
		},
	}
}
