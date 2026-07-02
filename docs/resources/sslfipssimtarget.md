---
subcategory: "SSL"
---

# Resource: sslfipssimtarget

The sslfipssimtarget resource drives the Secure Information Management (SIM) **import** action on the *target* Citrix ADC FIPS appliance. It is used during FIPS key transfer to import the secret/key material previously exported from a source FIPS appliance (using the source certificate, source secret, key vector, and target secret) into the target appliance's Hardware Security Module. This is a one-shot action resource, not a persistent configuration object.

~> **WARNING: FIPS / HSM hardware required.**
> This resource requires a dedicated FIPS appliance with an on-board Hardware Security Module. It is **not supported on non-FIPS appliances** (including VPX/CPX/standard MPX models) and will fail there.
>
> This resource invokes a FIPS SIM key-import action that loads cryptographic key material into the target appliance's FIPS HSM. Destroying the resource only removes it from Terraform state and does not undo the import performed on the appliance. Plan FIPS key transfers carefully and keep a backup of the appliance configuration.

-> **Note (one-shot action):** This is an action-only resource. The NITRO API exposes only the `enable`/`init` actions for it and provides **no GET, update, or delete endpoint**. As a result, the provider's Read is a no-op (drift cannot be detected), Update is a no-op, and Delete only removes the entry from Terraform state. All attributes force replacement when changed. Because there is no GET endpoint, no data source is published for this resource, and import is not meaningful.

## Example usage

```hcl
variable "sslfipssimtarget_sourcesecret" {
  type      = string
  sensitive = true
}

variable "sslfipssimtarget_targetsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfipssimtarget" "tf_sslfipssimtarget" {
  certfile     = "ns-server.cert"
  keyvector    = "kv.key"
  sourcesecret = var.sslfipssimtarget_sourcesecret
  targetsecret = var.sslfipssimtarget_targetsecret
}
```

## Argument Reference

* `certfile` - (Required) Name of and, optionally, path to the source FIPS appliance's certificate file. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `keyvector` - (Required) Name of and, optionally, path to the target FIPS appliance's key vector. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `sourcesecret` - (Required, Sensitive) Name of and, optionally, path to the source FIPS appliance's secret data. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `targetsecret` - (Required, Sensitive) Name for and, optionally, path to the target FIPS appliance's secret data. The default input path for the secret data is `/nsconfig/ssl/`. Changing this attribute forces a new resource to be created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslfipssimtarget. Because this is an action-only resource with no GET endpoint, it is a synthetic constant string `"sslfipssimtarget-config"`.
