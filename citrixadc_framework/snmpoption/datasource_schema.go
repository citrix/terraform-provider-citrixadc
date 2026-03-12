package snmpoption

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SnmpoptionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"partitionnameintrap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send partition name as a varbind in traps. By default the partition names are not sent as a varbind.",
			},
			"severityinfointrap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By default, the severity level info of the trap is not mentioned in the trap message. Enable this option to send severity level of trap as one of the varbind in the trap message.",
			},
			"snmpset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Accept SNMP SET requests sent to the Citrix ADC, and allow SNMP managers to write values to MIB objects that are configured for write access.",
			},
			"snmptraplogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log any SNMP trap events (for SNMP alarms in which logging is enabled) even if no trap listeners are configured. With the default setting, SNMP trap events are logged if at least one trap listener is configured on the appliance.",
			},
			"snmptraplogginglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audit log level of SNMP trap logs. The default value is INFORMATIONAL.",
			},
		},
	}
}
