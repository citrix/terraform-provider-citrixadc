package clusterinstance

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ClusterinstanceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"backplanebasedview": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "View based on heartbeat only on bkplane interface",
			},
			"clid": schema.Int64Attribute{
				Required:    true,
				Description: "Unique number that identifies the cluster.",
			},
			"clusterproxyarp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This field controls the proxy arp feature in cluster. By default the flag is enabled.",
			},
			"deadinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of time, in seconds, after which nodes that do not respond to the heartbeats are assumed to be down.If the value is less than 3 sec, set the helloInterval parameter to 200 msec",
			},
			"dfdretainl2params": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "flag to add ext l2 header during steering. By default the flag is disabled.",
			},
			"hellointerval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in milliseconds, at which heartbeats are sent to each cluster node to check the health status.Set the value to 200 msec, if the deadInterval parameter is less than 3 sec",
			},
			"inc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is required if the cluster nodes reside on different networks.",
			},
			"nodegroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The node group in a Cluster system used for transition from L2 to L3.",
			},
			"preemption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Preempt a cluster node that is configured as a SPARE if an ACTIVE node becomes available.",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By turning on this option packets destined to a service in a cluster will not under go any steering.",
			},
			"quorumtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Quorum Configuration Choices  - \"Majority\" (recommended) requires majority of nodes to be online for the cluster to be UP. \"None\" relaxes this requirement.",
			},
			"retainconnectionsoncluster": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled.",
			},
			"secureheartbeats": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By turning on this option cluster heartbeats will have security enabled.",
			},
			"syncstatusstrictmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "strict mode for sync status of cluster. Depending on the the mode if there are any errors while applying config, sync status is displayed accordingly. By default the flag is disabled.",
			},
		},
	}
}
