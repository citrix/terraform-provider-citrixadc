package mapbmr

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func MapbmrDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"eabitlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The Embedded Address (EA) bit field encodes the CE-specific IPv4 address and port information.  The EA bit field, which is unique for a\n			          given Rule IPv6 prefix.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Basic Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Basic Mapping Rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"add network MapBmr bmr1 -natprefix 2005::/64 -EAbitLength 16 -psidoffset 6 -portsharingratio 8\" ).\n			The Basic Mapping Rule information allows a MAP BR to determine source IPv4 address from the IPv6 packet sent from MAP CE device.\n			Also it allows to determine destination IPv6 address of MAP CE before sending packets to MAP CE",
			},
			"psidlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Length of Port Set IdentifierPort Set Identifier(PSID) in Embedded Address (EA) bits",
			},
			"psidoffset": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Start bit position  of Port Set Identifier(PSID) value in Embedded Address (EA) bits.",
			},
			"ruleipv6prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 prefix of Customer Edge(CE) device.MAP-T CE will send ipv6 packets with this ipv6 prefix as source ipv6 address prefix",
			},
		},
	}
}
