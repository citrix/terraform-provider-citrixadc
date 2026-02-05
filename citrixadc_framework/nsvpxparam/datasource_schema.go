package nsvpxparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsvpxparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cpuyield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting applicable in virtual appliances, is to affect the cpu yield(relinquishing the cpu resources) in any hypervised environment.\n\n* There are 3 options for the behavior:\n1. YES - Allow the Virtual Appliance to yield its vCPUs periodically, if there is no data traffic.\n2. NO - Virtual Appliance will not yield the vCPU.\n3. DEFAULT - Restores the default behaviour, according to the license.\n\n* Its behavior in different scenarios:\n1. As this setting is node specific only, it will not be propagated to other nodes, when executed on Cluster(CLIP) and HA(Primary).\n2. In cluster setup, use '-ownerNode' to specify ID of the cluster node.\n3. This setting is a system wide implementation and not granular to vCPUs.\n4. No effect on the management PE.",
			},
			"kvmvirtiomultiqueue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting applicable on KVM VPX with virtio NICs, is to configure multiple queues for all virtio interfaces.\n\n* There are 2 options for this behavior:\n1. YES - Allows VPX to use multiple queues for each virtio interface as configured through the KVM Hypervisor.\n2. NO - Each virtio interface within VPX will use a single queue for transmit and receive.\n\n* Its behavior in different scenarios:\n1. As this setting is node specific only, it will not be propagated to other nodes, when executed on Cluster(CLIP) and HA(Primary).\n2. In cluster setup, use '-ownerNode' to specify ID of the cluster node.",
			},
			"masterclockcpu1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is deprecated.",
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the cluster node for which you are setting the cpuyield and/or KVMVirtioMultiqueue. It can be configured only through the cluster IP address.",
			},
		},
	}
}
