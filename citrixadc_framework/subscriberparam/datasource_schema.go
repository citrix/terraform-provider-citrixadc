package subscriberparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SubscriberparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"idleaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Once idleTTL exprires on a subscriber session, Citrix ADC will take an idle action on that session. idleAction could be chosen from one of these ==>\n1. ccrTerminate: (default) send CCR-T to inform PCRF about session termination and delete the session.  \n2. delete: Just delete the subscriber session without informing PCRF.\n3. ccrUpdate: Do not delete the session and instead send a CCR-U to PCRF requesting for an updated session. !",
			},
			"idlettl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Idle Timeout, in seconds, after which Citrix ADC will take an idleAction on a subscriber session (refer to 'idleAction' arguement in 'set subscriber param' for more details on idleAction). Any data-plane or control plane activity updates the idleTimeout on subscriber session. idleAction could be to 'just delete the session' or 'delete and CCR-T' (if PCRF is configured) or 'do not delete but send a CCR-U'. \nZero value disables the idle timeout. !",
			},
			"interfacetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subscriber Interface refers to Citrix ADC interaction with control plane protocols, RADIUS and GX.\nTypes of subscriber interface: NONE, RadiusOnly, RadiusAndGx, GxOnly.\nNONE: Only static subscribers can be configured.\nRadiusOnly: GX interface is absent. Subscriber information is obtained through RADIUS Accounting messages.\nRadiusAndGx: Subscriber ID obtained through RADIUS Accounting is used to query PCRF. Subscriber information is obtained from both RADIUS and PCRF.\nGxOnly: RADIUS interface is absent. Subscriber information is queried using Subscriber IP or IP+VLAN.",
			},
			"ipv6prefixlookuplist": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The ipv6PrefixLookupList should consist of all the ipv6 prefix lengths assigned to the UE's'",
			},
			"keytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of subscriber key type IP or IPANDVLAN. IPANDVLAN option can be used only when the interfaceType is set to gxOnly.\nChanging the lookup method should result to the subscriber session database being flushed.",
			},
		},
	}
}
