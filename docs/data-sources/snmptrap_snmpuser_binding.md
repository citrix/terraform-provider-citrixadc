---
subcategory: "SNMP"
---

# Data Source: snmptrap_snmpuser_binding

The snmptrap_snmpuser_binding data source allows you to retrieve information about an snmp trap snmp user binding.

## Example Usage

```terraform
data "citrixadc_snmptrap_snmpuser_binding" "tf_binding" {
  trapclass       = "generic"
  trapdestination = "10.50.50.11"
  username        = "tf_snmpuser_ds"
  td              = 0
  version         = "V3"
}

output "securitylevel" {
  value = data.citrixadc_snmptrap_snmpuser_binding.tf_binding.securitylevel
}

output "username" {
  value = data.citrixadc_snmptrap_snmpuser_binding.tf_binding.username
}
```

## Argument Reference

* `td` - (Required) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `trapclass` - (Required) Type of trap messages that the Citrix ADC sends to the trap listener: Generic or the enterprise-specific messages defined in the MIB file.
* `trapdestination` - (Required) IPv4 or the IPv6 address of the trap listener to which the Citrix ADC is to send SNMP trap messages.
* `username` - (Required) Name of the SNMP user that will send the SNMPv3 traps.
* `version` - (Required) SNMP version, which determines the format of trap messages sent to the trap listener. This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `securitylevel` - Security level of the SNMPv3 trap.
* `id` - The id of the snmptrap_snmpuser_binding. It is a system-generated identifier.
