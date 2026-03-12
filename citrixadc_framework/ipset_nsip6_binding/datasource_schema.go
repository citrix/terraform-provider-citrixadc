package ipset_nsip6_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IpsetNsip6BindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "One or more IP addresses bound to the IP set.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the IP set to which to bind IP addresses.",
			},
		},
	}
}
