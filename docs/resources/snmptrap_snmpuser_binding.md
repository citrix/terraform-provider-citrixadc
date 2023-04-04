---
subcategory: "SNMP"
---

# Resource: snmptrap_snmpuser_binding

The snmptrap_snmpuser_binding resource is used to bind snmp user to snmp trap reosurce.


## Example usage

```hcl
resource "citrixadc_snmptrap" "tf_snmptrap" {
  trapclass       = "generic"
  trapdestination = "10.10.10.10"
  version         = "V3"
}
resource "citrixadc_snmpuser" "tf_snmpuser" {
  name       = "tf_snmpuser"
  group      = "all_group"
  authtype   = "SHA"
  authpasswd = "secretpassword"
  privtype   = "AES"
  privpasswd = "secretpassword"
}
resource "citrixadc_snmptrap_snmpuser_binding" "tf_binding" {
  trapclass       = citrixadc_snmptrap.tf_snmptrap.trapclass
  trapdestination = citrixadc_snmptrap.tf_snmptrap.trapdestination
  username        = citrixadc_snmpuser.tf_snmpuser.name
  securitylevel   = "authPriv"
}

```


## Argument Reference

* `trapclass` - (Required) Type of trap messages that the Citrix ADC sends to the trap listener: Generic or the enterprise-specific messages defined in the MIB file. Possible values: [ generic, specific ]
* `trapdestination` - (Required) IPv4 or the IPv6 address of the trap listener to which the Citrix ADC is to send SNMP trap messages. Minimum length =  1
* `username` - (Required) Name of the SNMP user that will send the SNMPv3 traps.
* `securitylevel` - (Optional) Security level of the SNMPv3 trap. Possible values: [ noAuthNoPriv, authNoPriv, authPriv ]
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `version` - (Optional) SNMP version, which determines the format of trap messages sent to the trap listener.  This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages. Possible values: [ V1, V2, V3 ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmptrap_snmpuser_binding. It is the concatenation of `trapclass` , `trapdestination` and `username` attribute values seperated by comma(",").


## Import

A snmptrap_snmpuser_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_snmptrap_snmpuser_binding.tf_binding generic,10.10.10.10,tf_snmpuser
```
