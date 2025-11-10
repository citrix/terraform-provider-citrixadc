---
subcategory: "NS"
---

# Resource: nslicenseparameters

The nslicenseparameters resource is used to create licenseparameters resource.


## Example usage

```hcl
resource "citrixadc_nslicenseparameters" "tf_nslicenseparameters" {
  alert1gracetimeout = 6
  alert2gracetimeout = 240
}
```


## Argument Reference

* `alert1gracetimeout` - (Optional) If ADC remains in grace for the configured hours then first grace alert will be raised. Minimum value =  0 Maximum value =  24
* `alert2gracetimeout` - (Optional) If ADC remains in grace for the configured hours then major grace alert will be raised. Minimum value =  24 Maximum value =  720
* `heartbeatinterval` - (Optional) Heartbeat between ADC and Licenseserver is configurable and applicable in case of pooled licensing
* `inventoryrefreshinterval` - (Optional) Inventory refresh interval between ADC and Licenseserver is configurable and applicable in case of pooled licensing
* `licenseexpiryalerttime` - (Optional) If ADC license contract expiry date is nearer then GUI/SNMP license expiry alert will be raised


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nslicenseparameters. It is a unique string prefixed with "tf-nslicenseparameters-"

