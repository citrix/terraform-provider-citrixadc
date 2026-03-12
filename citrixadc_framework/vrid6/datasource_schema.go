package vrid6

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Vrid6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove all configured VMAC6 addresses from the Citrix ADC.",
			},
			"vrid6_id": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a VMAC6 address.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "In a cluster setup, assign a cluster node as the owner of this VMAC address for IP based VRRP configuration. If no owner is configured, ow ner node is displayed as ALL and one node is dynamically elected as the owner.",
			},
			"preemption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In an active-active mode configuration, make a backup VIP address the master if its priority becomes higher than that of a master VIP address bound to this VMAC address.\n             If you disable pre-emption while a backup VIP address is the master, the backup VIP address remains master until the original master VIP's priority becomes higher than that of the current master.",
			},
			"preemptiondelaytimer": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Preemption delay time in seconds, in an active-active configuration. If any high priority node will come in network, it will wait for these many seconds before becoming master.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Base priority (BP), in an active-active mode configuration, which ordinarily determines the master VIP address.",
			},
			"sharing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In an active-active mode configuration, enable the backup VIP address to process any traffic instead of dropping it.",
			},
			"trackifnumpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority by which the Effective priority will be reduced if any of the tracked interfaces goes down in an active-active configuration.",
			},
			"tracking": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The effective priority (EP) value, relative to the base priority (BP) value in an active-active mode configuration. When EP is set to a value other than None, it is EP, not BP, which determines the master VIP address.\nAvailable settings function as follows:\n* NONE - No tracking. EP = BP\n* ALL -  If the status of all virtual servers is UP, EP = BP. Otherwise, EP = 0.\n* ONE - If the status of at least one virtual server is UP, EP = BP. Otherwise, EP = 0.\n* PROGRESSIVE - If the status of all virtual servers is UP, EP = BP. If the status of all virtual servers is DOWN, EP = 0. Otherwise EP = BP (1 - K/N), where N is the total number of virtual servers associated with the VIP address and K is the number of virtual servers for which the status is DOWN.\nDefault: NONE.",
			},
		},
	}
}
