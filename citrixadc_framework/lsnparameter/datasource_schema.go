package lsnparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of Citrix ADC memory to reserve for the LSN feature, in multiples of 2MB.\n\nNote: If you later reduce the value of this parameter, the amount of active memory is not reduced. Changing the configured memory limit can only increase the amount of active memory.\nThis command is deprecated, use 'set extendedmemoryparam -memlimit' instead.",
			},
			"sessionsync": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Synchronize all LSN sessions with the secondary node in a high availability (HA) deployment (global synchronization). After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary node (new primary).\n\nThe global session synchronization parameter and session synchronization parameters (group level) of all LSN groups are enabled by default.\n\nFor a group, when both the global level and the group level LSN session synchronization parameters are enabled, the primary node synchronizes information of all LSN sessions related to this LSN group with the secondary node.",
			},
			"subscrsessionremoval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LSN global setting for controlling subscriber aware session removal, when this is enabled, when ever the subscriber info is deleted from subscriber database, sessions corresponding to that subscriber will be removed. if this setting is disabled, subscriber sessions will be timed out as per the idle time out settings.",
			},
		},
	}
}
