package nd6ravariables_onlinkipv6prefix_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Nd6ravariablesOnlinkipv6prefixBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipv6prefix": schema.StringAttribute{
				Required:    true,
				Description: "Onlink prefixes for RA messages.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "The VLAN number.",
			},
		},
	}
}
