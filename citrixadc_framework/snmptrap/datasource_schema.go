package snmptrap

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SnmptrapDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allpartitions": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send traps of all partitions to this destination.",
			},
			"communityname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password (string) sent with the trap messages, so that the trap listener can authenticate them. Can include 1 to 31 uppercase or lowercase letters, numbers, and hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters.  \nYou must specify the same community string on the trap listener device. Otherwise, the trap listener drops the trap messages.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the string includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my string\" or 'my string').",
			},
			"destport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "UDP port at which the trap listener listens for trap messages. This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.",
			},
			"severity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Severity level at or above which the Citrix ADC sends trap messages to this trap listener. The severity levels, in increasing order of severity, are Informational, Warning, Minor, Major, Critical. This parameter can be set for trap listeners of type SPECIFIC only. The default is to send all levels of trap messages. \nImportant: Trap messages are not assigned severity levels unless you specify severity levels when configuring SNMP alarms.",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address that the Citrix ADC inserts as the source IP address in all SNMP trap messages that it sends to this trap listener. By default this is the appliance's NSIP or NSIP6 address, but you can specify an IPv4 MIP or SNIP/SNIP6 address. In cluster setup, the default value is the individual node's NSIP, but it can be set to CLIP or Striped SNIP address. In non default partition, this parameter must be set to the SNIP/SNIP6 address.",
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
			"version": schema.StringAttribute{
				Required:    true,
				Description: "SNMP version, which determines the format of trap messages sent to the trap listener. \nThis setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.",
			},
		},
	}
}
