package mapbmr_bmrv4network_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func MapbmrBmrv4networkBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Basic Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Basic Mapping Rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"add network MapBmr bmr1 -natprefix 2005::/64 -EAbitLength 16 -psidoffset 6 -portsharingratio 8\" ).\n			The Basic Mapping Rule information allows a MAP BR to determine source IPv4 address from the IPv6 packet sent from MAP CE device.\n			Also it allows to determine destination IPv6 address of MAP CE before sending packets to MAP CE",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask for the IPv4 address specified in the Network parameter.",
			},
			"network": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 NAT address range of Customer Edge (CE).",
			},
		},
	}
}
