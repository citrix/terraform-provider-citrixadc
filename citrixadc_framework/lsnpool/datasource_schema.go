package lsnpool

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnpoolDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"maxportrealloctmq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of ports for which the port reallocation timeout applies for each NAT IP address. In other words, the maximum deallocated-port queue size for which the reallocation timeout applies for each NAT IP address.\n\nWhen the queue size is full, the next port deallocated is reallocated immediately for a new LSN session.",
			},
			"nattype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of NAT IP address and port allocation (from the LSN pools bound to an LSN group) for subscribers (of the LSN client entity bound to the LSN group):\n\nAvailable options function as follows:\n\n* Deterministic - Allocate a NAT IP address and a block of ports to each subscriber (of the LSN client bound to the LSN group). The Citrix ADC sequentially allocates NAT resources to these subscribers. The Citrix ADC ADC assigns the first block of ports (block size determined by the port block size parameter of the LSN group) on the beginning NAT IP address to the beginning subscriber IP address. The next range of ports is assigned to the next subscriber, and so on, until the NAT address does not have enough ports for the next subscriber. In this case, the first port block on the next NAT address is used for the subscriber, and so on.  Because each subscriber now receives a deterministic NAT IP address and a block of ports, a subscriber can be identified without any need for logging. For a connection, a subscriber can be identified based only on the NAT IP address and port, and the destination IP address and port.\n \n* Dynamic - Allocate a random NAT IP address and a port from the LSN NAT pool for a subscriber's connection. If port block allocation is enabled (in LSN pool) and a port block size is specified (in the LSN group), the Citrix ADC allocates a random NAT IP address and a block of ports for a subscriber when it initiates a connection for the first time. The ADC allocates this NAT IP address and a port (from the allocated block of ports) for different connections from this subscriber. If all the ports are allocated (for different subscriber's connections) from the subscriber's allocated port block, the ADC allocates a new random port block for the subscriber.\nOnly LSN Pools and LSN groups with the same NAT type settings can be bound together. Multiples LSN pools can be bound to an LSN group. A maximum of 16 LSN pools can be bound to an LSN group.",
			},
			"poolname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN pool. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN pool is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn pool1\" or 'lsn pool1').",
			},
			"portblockallocation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allocate a random NAT port block, from the available NAT port pool of an NAT IP address, for each subscriber when the NAT allocation is set as Dynamic NAT. For any connection initiated from a subscriber, the Citrix ADC allocates a NAT port from the subscriber's allocated NAT port block to create the LSN session.\n\nYou must set the port block size in the bound LSN group. For a subscriber, if all the ports are allocated from the subscriber's allocated port block, the Citrix ADC allocates a new random port block for the subscriber.\n\nFor Deterministic NAT, this parameter is enabled by default, and you cannot disable it.",
			},
			"portrealloctimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The waiting time, in seconds, between deallocating LSN NAT ports (when an LSN mapping is removed) and reallocating them for a new LSN session. This parameter is necessary in order to prevent collisions between old and new mappings and sessions. It ensures that all established sessions are broken instead of redirected to a different subscriber. This is not applicable for ports used in:\n* Deterministic NAT\n* Address-Dependent filtering and Address-Port-Dependent filtering\n* Dynamic NAT with port block allocation\nIn these cases, ports are immediately reallocated.",
			},
		},
	}
}
