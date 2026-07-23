---
subcategory: "SSL"
---

# Resource: sslfipssimtarget_enable

This resource is used to run the SIM `enable` action on the target Citrix ADC FIPS appliance.

~> **WARNING:** Requires a dedicated FIPS appliance with an on-board HSM.


## Example usage

```hcl
variable "sslfipssimtarget_enable_sourcesecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfipssimtarget_enable" "tf_sslfipssimtarget_enable" {
  keyvector    = "kv.key"
  sourcesecret = var.sslfipssimtarget_enable_sourcesecret
}
```

## Argument Reference

* `keyvector` - (Required) Name of and, optionally, path to the target FIPS appliance's key vector. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `sourcesecret` - (Required, Sensitive) Name of and, optionally, path to the source FIPS appliance's secret data. `/nsconfig/ssl/` is the default path. The value is persisted in Terraform state. Changing this attribute forces a new resource to be created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslfipssimtarget_enable resource. It is set to `sslfipssimtarget_enable`.
