---
subcategory: "NS"
---

# Data Source `nslicenseparameters`

The nslicenseparameters data source allows you to retrieve information about NetScaler license parameters configuration.


## Example usage

```terraform
data "citrixadc_nslicenseparameters" "tf_nslicenseparameters" {
}

output "alert1gracetimeout" {
  value = data.citrixadc_nslicenseparameters.tf_nslicenseparameters.alert1gracetimeout
}

output "heartbeatinterval" {
  value = data.citrixadc_nslicenseparameters.tf_nslicenseparameters.heartbeatinterval
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `alert1gracetimeout` - If ADC remains in grace for the configured hours then first grace alert will be raised.
* `alert2gracetimeout` - If ADC remains in grace for the configured hours then major grace alert will be raised.
* `heartbeatinterval` - Heartbeat between ADC and Licenseserver is configurable and applicable in case of pooled licensing.
* `inventoryrefreshinterval` - Inventory refresh interval between ADC and Licenseserver is configurable and applicable in case of pooled licensing.
* `licenseexpiryalerttime` - If ADC license contract expiry date is nearer then GUI/SNMP license expiry alert will be raised.

## Attribute Reference

* `id` - The id of the nslicenseparameters. It is a system-generated identifier.
