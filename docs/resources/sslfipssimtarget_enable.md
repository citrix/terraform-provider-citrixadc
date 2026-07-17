---
subcategory: "SSL"
---

# Resource: sslfipssimtarget_enable

The sslfipssimtarget_enable resource drives the **enable** stage of the Secure Information Management (SIM) key-transfer workflow on the *target* Citrix ADC FIPS appliance. During FIPS key transfer, this action uses the target appliance's key vector together with the source appliance's secret data to enable the target for importing the secret/key material exported from a source FIPS appliance into the target appliance's Hardware Security Module. This is a one-shot action resource, not a persistent configuration object.

~> **WARNING: FIPS / HSM hardware required.**
> This resource requires a dedicated FIPS appliance with an on-board Hardware Security Module. It is **not supported on non-FIPS appliances** (including VPX/CPX/standard MPX models) and will fail there.
>
> This resource invokes a FIPS SIM key-transfer action that operates on cryptographic key material in the target appliance's FIPS HSM. Destroying the resource only removes it from Terraform state and does not undo the action performed on the appliance. Plan FIPS key transfers carefully and keep a backup of the appliance configuration.

-> **Note (one-shot action):** This is an action-only resource. The NITRO API exposes only the `enable` action for it and provides **no GET, update, or delete endpoint**. As a result, the provider's Read is a no-op (drift cannot be detected), Update is a no-op, and Delete only removes the entry from Terraform state. All attributes force replacement when changed. Because there is no GET endpoint, no data source is published for this resource, and import is not meaningful.

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

* `id` - A synthetic identifier for this action-only resource. Because there is no GET endpoint, it is a fixed string with the value `sslfipssimtarget_enable`. It does not correspond to any object on the Citrix ADC.
