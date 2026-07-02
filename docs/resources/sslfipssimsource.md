---
subcategory: "SSL"
---

# Resource: sslfipssimsource

The sslfipssimsource resource drives the Secure Information Management (SIM) **export** action on the *source* Citrix ADC FIPS appliance. It is used during FIPS key transfer to export the source appliance's secret/key material (protected by a target secret and the source/target certificate) so that it can be securely imported onto a target FIPS appliance. This is a one-shot action resource, not a persistent configuration object.

~> **WARNING: FIPS / HSM hardware required and not cleanly reversible.**
> This resource requires a dedicated FIPS appliance with an on-board Hardware Security Module. It is **not supported on non-FIPS appliances** (including VPX/CPX/standard MPX models) and will fail there.
>
> This resource invokes a FIPS SIM key-export action. The action manipulates cryptographic key material on the FIPS HSM and is **not cleanly reversible** — destroying the resource only removes it from Terraform state and does not undo the export performed on the appliance. Plan FIPS key transfers carefully and keep a backup of the appliance configuration.

-> **Note (one-shot action):** This is an action-only resource. The NITRO API exposes only the `enable`/`init` actions for it and provides **no GET, update, or delete endpoint**. As a result, the provider's Read is a no-op (drift cannot be detected), Update is a no-op, and Delete only removes the entry from Terraform state. All attributes force replacement when changed. Because there is no GET endpoint, no data source is published for this resource, and import is not meaningful.

## Example usage

```hcl
variable "sslfipssimsource_sourcesecret" {
  type      = string
  sensitive = true
}

variable "sslfipssimsource_targetsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfipssimsource" "tf_sslfipssimsource" {
  certfile     = "ns-server.cert"
  sourcesecret = var.sslfipssimsource_sourcesecret
  targetsecret = var.sslfipssimsource_targetsecret
}
```

## Argument Reference

* `certfile` - (Required) Name for and, optionally, path to the source FIPS appliance's certificate file. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `sourcesecret` - (Required, Sensitive) Name for and, optionally, path to the source FIPS appliance's secret data. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `targetsecret` - (Required, Sensitive) Name of and, optionally, path to the target FIPS appliance's secret data. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslfipssimsource. Because this is an action-only resource with no GET endpoint, it is a synthetic constant string `"sslfipssimsource-config"`.
