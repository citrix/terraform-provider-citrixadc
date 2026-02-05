package arpparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ArpparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"spoofvalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "enable/disable arp spoofing validation",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time-out value (aging time) for the dynamically learned ARP entries, in seconds. The new value applies only to ARP entries that are dynamically learned after the new value is set. Previously existing ARP entries expire after the previously configured aging time.",
			},
		},
	}
}
