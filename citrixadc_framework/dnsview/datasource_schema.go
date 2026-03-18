package dnsview

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnsviewDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"viewname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DNS view.",
			},
		},
	}
}
