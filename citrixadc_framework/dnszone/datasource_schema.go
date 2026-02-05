package dnszone

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DnszoneDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dnssecoffload": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable dnssec offload for this zone.",
			},
			"keyname": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name of the public/private DNS key pair with which to sign the zone. You can sign a zone with up to four keys.",
			},
			"nsec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable nsec generation for dnssec offload.",
			},
			"proxymode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Deploy the zone in proxy mode. Enable in the following scenarios:\n* The load balanced DNS servers are authoritative for the zone and all resource records that are part of the zone.\n* The load balanced DNS servers are authoritative for the zone, but the Citrix ADC owns a subset of the resource records that belong to the zone (partial zone ownership configuration). Typically seen in global server load balancing (GSLB) configurations, in which the appliance responds authoritatively to queries for GSLB domain names but forwards queries for other domain names in the zone to the load balanced servers.\nIn either scenario, do not create the zone's Start of Authority (SOA) and name server (NS) resource records on the appliance.\nDisable if the appliance is authoritative for the zone, but make sure that you have created the SOA and NS records on the appliance before you create the zone.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Type of zone to display. Mutually exclusive with the DNS Zone (zoneName) parameter. Available settings function as follows:\n* ADNS - Display all the zones for which the Citrix ADC is authoritative.\n* PROXY - Display all the zones for which the Citrix ADC is functioning as a proxy server.\n* ALL - Display all the zones configured on the appliance.",
			},
			"zonename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the zone to create.",
			},
		},
	}
}
