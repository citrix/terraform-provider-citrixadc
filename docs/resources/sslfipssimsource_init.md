---
subcategory: "SSL"
---

# Resource: sslfipssimsource_init

This resource is used to run the SIM `init` action on the source Citrix ADC FIPS appliance.

~> **WARNING:** Requires a dedicated FIPS appliance with an on-board HSM; this one-shot action is not cleanly reversible.


## Example usage

```hcl
resource "citrixadc_sslfipssimsource_init" "tf_sslfipssimsource_init" {
  certfile = "ns-server.cert"
}
```

## Argument Reference

* `certfile` - (Required) Name for and, optionally, path to the source FIPS appliance's certificate file. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslfipssimsource_init resource. It is set to `sslfipssimsource_init`.
