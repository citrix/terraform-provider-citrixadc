---
subcategory: "SNMP"
---

# Data Source: citrixadc_snmpoption

The `citrixadc_snmpoption` data source is used to retrieve SNMP option configuration from the Citrix ADC.

## Example usage

```hcl
data "citrixadc_snmpoption" "example" {
}
```

## Argument Reference

No arguments are required for this data source.

## Attribute Reference

The following attributes are exported:

* `id` - The ID of the SNMP option resource.
* `partitionnameintrap` - Send partition name as a varbind in traps. By default the partition names are not sent as a varbind.
* `severityinfointrap` - By default, the severity level info of the trap is not mentioned in the trap message. Enable this option to send severity level of trap as one of the varbind in the trap message.
* `snmpset` - Accept SNMP SET requests sent to the Citrix ADC, and allow SNMP managers to write values to MIB objects that are configured for write access.
* `snmptraplogging` - Log any SNMP trap events (for SNMP alarms in which logging is enabled) even if no trap listeners are configured. With the default setting, SNMP trap events are logged if at least one trap listener is configured on the appliance.
* `snmptraplogginglevel` - Audit log level of SNMP trap logs. The default value is INFORMATIONAL.
