---
subcategory: "SSL"
---

# Resource: sslfipssimtarget_init

This resource is used to run the SIM `init` action on the target Citrix ADC FIPS appliance.

~> **WARNING:** Requires a dedicated FIPS appliance with an on-board HSM.


## Example usage

```hcl
variable "sslfipssimtarget_init_targetsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfipssimtarget_init" "tf_sslfipssimtarget_init" {
  certfile     = "ns-server.cert"
  keyvector    = "kv.key"
  targetsecret = var.sslfipssimtarget_init_targetsecret
}
```

## Argument Reference

* `certfile` - (Required) Name of and, optionally, path to the source FIPS appliance's certificate file. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `keyvector` - (Required) Name of and, optionally, path to the target FIPS appliance's key vector. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `targetsecret` - (Required, Sensitive) Name for and, optionally, path to the target FIPS appliance's secret data. The default input path for the secret data is `/nsconfig/ssl/`. The value is persisted in Terraform state. Changing this attribute forces a new resource to be created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslfipssimtarget_init resource. It is set to `sslfipssimtarget_init`.
