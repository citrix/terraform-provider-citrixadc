---
subcategory: "SNMP"
---

# Data Source: citrixadc_snmptrap

The `citrixadc_snmptrap` data source is used to retrieve information about an SNMP trap configuration on the Citrix ADC.

## Example usage

```hcl
data "citrixadc_snmptrap" "example" {
  trapclass       = "specific"
  trapdestination = "192.168.1.100"
  version         = "V2"
  td              = 0
}
```

## Argument Reference

* `trapclass` - (Required) Type of trap messages that the Citrix ADC sends to the trap listener: Generic or the enterprise-specific messages defined in the MIB file.
* `trapdestination` - (Required) IPv4 or the IPv6 address of the trap listener to which the Citrix ADC is to send SNMP trap messages.
* `version` - (Required) SNMP version, which determines the format of trap messages sent to the trap listener. This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.
* `td` - (Required) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the SNMP trap.
* `allpartitions` - Send traps of all partitions to this destination.
* `communityname` - Password (string) sent with the trap messages, so that the trap listener can authenticate them.
* `destport` - UDP port at which the trap listener listens for trap messages. This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.
* `severity` - Severity level at or above which the Citrix ADC sends trap messages to this trap listener. The severity levels, in increasing order of severity, are Informational, Warning, Minor, Major, Critical.
* `srcip` - IPv4 or IPv6 address that the Citrix ADC inserts as the source IP address in all SNMP trap messages that it sends to this trap listener.
