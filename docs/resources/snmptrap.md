---
subcategory: "SNMP"
---

# Resource: snmptrap

The snmptrap resource is used to create snmptrap.


## Example usage

```hcl
resource "citrixadc_snmptrap" "tf_snmptrap" {
  severity        = "Major"
  trapclass       = "specific"
  trapdestination = "192.168.2.2"
}

```


## Argument Reference

* `trapclass` - (Required) Type of trap messages that the Citrix ADC sends to the trap listener: Generic or the enterprise-specific messages defined in the MIB file.
* `trapdestination` - (Required) IPv4 or the IPv6 address of the trap listener to which the Citrix ADC is to send SNMP trap messages.
* `allpartitions` - (Optional) Send traps of all partitions to this destination.
* `communityname` - (Optional) Password (string) sent with the trap messages, so that the trap listener can authenticate them. Can include 1 to 31 uppercase or lowercase letters, numbers, and hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters.   You must specify the same community string on the trap listener device. Otherwise, the trap listener drops the trap messages.  The following requirement applies only to the Citrix ADC CLI: If the string includes one or more spaces, enclose the name in double or single quotation marks (for example, "my string" or 'my string').
* `destport` - (Optional) UDP port at which the trap listener listens for trap messages. This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.
* `severity` - (Optional) Severity level at or above which the Citrix ADC sends trap messages to this trap listener. The severity levels, in increasing order of severity, are Informational, Warning, Minor, Major, Critical. This parameter can be set for trap listeners of type SPECIFIC only. The default is to send all levels of trap messages.  Important: Trap messages are not assigned severity levels unless you specify severity levels when configuring SNMP alarms.
* `srcip` - (Optional) IPv4 or IPv6 address that the Citrix ADC inserts as the source IP address in all SNMP trap messages that it sends to this trap listener. By default this is the appliance's NSIP or NSIP6 address, but you can specify an IPv4 MIP or SNIP/SNIP6 address. In cluster setup, the default value is the individual node's NSIP, but it can be set to CLIP or Striped SNIP address. In non default partition, this parameter must be set to the SNIP/SNIP6 address.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `version` - (Optional) SNMP version, which determines the format of trap messages sent to the trap listener.  This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmptrap is the concatenation of the `trapclass` and `trapdestination` attributes separated by a comma.


## Import

A snmptrap can be imported using its name, e.g.

```shell
terraform import citrixadc_snmptrap.tf_snmptrap specific,192.168.2.2
```