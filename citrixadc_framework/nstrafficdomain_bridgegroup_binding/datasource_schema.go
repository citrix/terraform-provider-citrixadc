package nstrafficdomain_bridgegroup_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NstrafficdomainBridgegroupBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bridgegroup": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the configured bridge to bind to this traffic domain. More than one bridge group can be bound to a traffic domain, but the same bridge group cannot be a part of multiple traffic domains.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a traffic domain.",
			},
		},
	}
}
