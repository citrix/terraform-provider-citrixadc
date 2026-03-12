package l4param

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func L4paramDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"l2connmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Layer 2 connection method based on the combination of  channel number, MAC address and VLAN. It is tuned with l2conn param of lb vserver. If l2conn of lb vserver is ON then method specified here will be used to identify a connection in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>).",
			},
			"l4switch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In L4 switch topology, always clients and servers are on the same side. Enable l4switch to allow such connections.",
			},
		},
	}
}
