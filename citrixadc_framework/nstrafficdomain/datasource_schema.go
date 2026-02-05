package nstrafficdomain

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NstrafficdomainDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aliasname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of traffic domain  being added.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a traffic domain.",
			},
			"vmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Associate the traffic domain with a VMAC address instead of with VLANs. The Citrix ADC then sends the VMAC address of the traffic domain in all responses to ARP queries for network entities in that domain. As a result, the ADC can segregate subsequent incoming traffic for this traffic domain on the basis of the destination MAC address, because the destination MAC address is the VMAC address of the traffic domain. After creating entities on a traffic domain, you can easily manage and monitor them by performing traffic domain level operations.",
			},
		},
	}
}
