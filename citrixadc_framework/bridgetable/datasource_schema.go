package bridgetable

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func BridgetableDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bridgeage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value.",
			},
			"devicevlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The vlan on which to send multicast packets when the VXLAN tunnel endpoint is a muticast group address.",
			},
			"ifnum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "INTERFACE  whose entries are to be removed.",
			},
			"mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The MAC address of the target.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN  whose entries are to be removed.",
			},
			"vni": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The VXLAN VNI Network Identifier (or VXLAN Segment ID) to use to connect to the remote VXLAN tunnel endpoint.  If omitted the value specified as vxlan will be used.",
			},
			"vtep": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IP address of the destination VXLAN tunnel endpoint where the Ethernet MAC ADDRESS resides.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The VXLAN to which this address is associated.",
			},
		},
	}
}
