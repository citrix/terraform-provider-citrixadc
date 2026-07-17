---
subcategory: "SSL"
---

# Resource: sslfipssimtarget_init

The sslfipssimtarget_init resource drives the **init** stage of the Secure Information Management (SIM) key-transfer workflow on the *target* Citrix ADC FIPS appliance. During FIPS key transfer, this action initializes the target appliance using the source appliance's certificate file, the target appliance's key vector, and the target appliance's secret data, preparing the target's Hardware Security Module to receive key material exported from a source FIPS appliance. This is a one-shot action resource, not a persistent configuration object.

~> **WARNING: FIPS / HSM hardware required.**
> This resource requires a dedicated FIPS appliance with an on-board Hardware Security Module. It is **not supported on non-FIPS appliances** (including VPX/CPX/standard MPX models) and will fail there.
>
> This resource invokes a FIPS SIM key-transfer action that operates on cryptographic key material in the target appliance's FIPS HSM. Destroying the resource only removes it from Terraform state and does not undo the action performed on the appliance. Plan FIPS key transfers carefully and keep a backup of the appliance configuration.

-> **Note (one-shot action):** This is an action-only resource. The NITRO API exposes only the `init` action for it and provides **no GET, update, or delete endpoint**. As a result, the provider's Read is a no-op (drift cannot be detected), Update is a no-op, and Delete only removes the entry from Terraform state. All attributes force replacement when changed. Because there is no GET endpoint, no data source is published for this resource, and import is not meaningful.

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

* `id` - A synthetic identifier for this action-only resource. Because there is no GET endpoint, it is a fixed string with the value `sslfipssimtarget_init`. It does not correspond to any object on the Citrix ADC.
