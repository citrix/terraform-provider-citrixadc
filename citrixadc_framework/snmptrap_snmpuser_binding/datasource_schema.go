package snmptrap_snmpuser_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SnmptrapSnmpuserBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"securitylevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Security level of the SNMPv3 trap.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"trapclass": schema.StringAttribute{
				Required:    true,
				Description: "Type of trap messages that the Citrix ADC sends to the trap listener: Generic or the enterprise-specific messages defined in the MIB file.",
			},
			"trapdestination": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 or the IPv6 address of the trap listener to which the Citrix ADC is to send SNMP trap messages.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SNMP user that will send the SNMPv3 traps.",
			},
			"version": schema.StringAttribute{
				Required:    true,
				Description: "SNMP version, which determines the format of trap messages sent to the trap listener. \nThis setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.",
			},
		},
	}
}
