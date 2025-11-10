---
subcategory: "SNMP"
---

# Resource: snmpopoption

The snmpopoption resource is used to create snmpopoption.


## Example usage

```hcl
resource "citrixadc_opoption" "tf_opoption" {
  snmpset              = "ENABLED"
  snmptraplogging      = "ENABLED"
  partitionnameintrap  = "ENABLED"
  snmptraplogginglevel = "WARNING"
}

```


## Argument Reference

* `partitionnameintrap` - (Optional) Send partition name as a varbind in traps. By default the partition names are not sent as a varbind.
* `snmpset` - (Optional) Accept SNMP SET requests sent to the Citrix ADC, and allow SNMP managers to write values to MIB objects that are configured for write access.
* `snmptraplogging` - (Optional) Log any SNMP trap events (for SNMP alarms in which logging is enabled) even if no trap listeners are configured. With the default setting, SNMP trap events are logged if at least one trap listener is configured on the appliance.
* `snmptraplogginglevel` - (Optional) Audit log level of SNMP trap logs. The default value is INFORMATIONAL.
* `severityinfointrap` - (Optional) By default, the severity level info of the trap is not mentioned in the trap message. Enable this option to send severity level of trap as one of the varbind in the trap message.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmpopoption. It is a unique string prefixed with "tf-snmpoption-".
