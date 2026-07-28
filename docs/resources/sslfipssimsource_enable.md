---
subcategory: "SSL"
---

# Resource: sslfipssimsource_enable

This resource is used to run the SIM `enable` action on the source Citrix ADC FIPS appliance.

~> **WARNING:** Requires a dedicated FIPS appliance with an on-board HSM; this one-shot action is not cleanly reversible.


## Example usage

```hcl
variable "sslfipssimsource_enable_sourcesecret" {
  type      = string
  sensitive = true
}

variable "sslfipssimsource_enable_targetsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfipssimsource_enable" "tf_sslfipssimsource_enable" {
  sourcesecret = var.sslfipssimsource_enable_sourcesecret
  targetsecret = var.sslfipssimsource_enable_targetsecret
}
```

## Argument Reference

* `sourcesecret` - (Required, Sensitive) Name for and, optionally, path to the source FIPS appliance's secret data. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `targetsecret` - (Required, Sensitive) Name of and, optionally, path to the target FIPS appliance's secret data. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslfipssimsource_enable resource. It is set to `sslfipssimsource_enable`.
