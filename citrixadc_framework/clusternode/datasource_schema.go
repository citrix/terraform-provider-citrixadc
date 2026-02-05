package clusternode

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ClusternodeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"backplane": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface through which the node communicates with the other nodes in the cluster. Must be specified in the three-tuple form n/c/u, where n represents the node ID and c/u refers to the interface on the appliance.",
			},
			"clearnodegroupconfig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to remove nodegroup config",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable for Passive node and node becomes passive after this timeout (in minutes)",
			},
			"force": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Node will be removed from cluster without prompting for user confirmation.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC IP (NSIP) address of the appliance to add to the cluster. Must be an IPv4 address.",
			},
			"nodegroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The default node group in a Cluster system.",
			},
			"nodeid": schema.Int64Attribute{
				Required:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Preference for selecting a node as the configuration coordinator. The node with the lowest priority value is selected as the configuration coordinator.\nWhen the current configuration coordinator goes down, the node with the next lowest priority is made the new configuration coordinator. When the original node comes back up, it will preempt the new configuration coordinator and take over as the configuration coordinator.\nNote: When priority is not configured for any of the nodes or if multiple nodes have the same priority, the cluster elects one of the nodes as the configuration coordinator.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Admin state of the cluster node. The available settings function as follows:\nACTIVE - The node serves traffic.\nSPARE - The node does not serve traffic unless an ACTIVE node goes down.\nPASSIVE - The node does not serve traffic, unless you change its state. PASSIVE state is useful during temporary maintenance activities in which you want the node to take part in the consensus protocol but not to serve traffic.",
			},
			"tunnelmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To set the tunnel mode",
			},
		},
	}
}
