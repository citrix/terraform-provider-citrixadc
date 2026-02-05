---
subcategory: "LLDP"
---

# Data Source `lldpparam`

The lldpparam data source allows you to retrieve information about LLDP (Link Layer Discovery Protocol) parameters configuration.


## Example usage

```terraform
data "citrixadc_lldpparam" "tf_lldpparam" {
}

output "mode" {
  value = data.citrixadc_lldpparam.tf_lldpparam.mode
}

output "timer" {
  value = data.citrixadc_lldpparam.tf_lldpparam.timer
}

output "holdtimetxmult" {
  value = data.citrixadc_lldpparam.tf_lldpparam.holdtimetxmult
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `holdtimetxmult` - A multiplier for calculating the duration for which the receiving device stores the LLDP information in its database before discarding or removing it. The duration is calculated as the holdtimeTxMult (Holdtime Multiplier) parameter value multiplied by the timer (Timer) parameter value.
* `mode` - Global mode of Link Layer Discovery Protocol (LLDP) on the Citrix ADC. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels.
* `timer` - Interval, in seconds, between LLDP packet data units (LLDPDUs).  that the Citrix ADC sends to a directly connected device.

## Attribute Reference

* `id` - The id of the lldpparam. It is a system-generated identifier.
