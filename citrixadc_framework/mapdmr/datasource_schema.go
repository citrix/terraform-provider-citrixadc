package mapdmr

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func MapdmrDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bripv6prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 prefix of Border Relay (Citrix ADC) device.MAP-T CE will send ipv6 packets to this ipv6 prefix.The DMR IPv6 prefix length SHOULD be 64 bits long by default and in any case MUST NOT exceed 96 bits",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Default Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the MAP Default Mapping Rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"add network MapDmr map1 -BRIpv6Prefix 2003::/96\").\n			Default Mapping Rule (DMR) is defined in terms of the IPv6 prefix advertised by one or more BRs, which provide external connectivity.  A typical MAP-T CE will install an IPv4 default route using this rule.  A BR will use this rule when translating all outside IPv4 source addresses to the IPv6 MAP domain.",
			},
		},
	}
}
