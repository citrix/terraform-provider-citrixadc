---
subcategory: "SSL"
---

# Resource: sslfipssimsource_enable

The sslfipssimsource_enable resource drives the Secure Information Management (SIM) **enable** action on the *source* Citrix ADC FIPS appliance. It is used during FIPS key transfer to activate the source appliance's secret/key material (the source appliance's own secret plus the target appliance's secret) so that key material can subsequently be exported to and imported onto a target FIPS appliance. This is a one-shot action resource, not a persistent configuration object.

~> **WARNING: FIPS / HSM hardware required and not cleanly reversible.**
> This resource requires a dedicated FIPS appliance with an on-board Hardware Security Module. It is **not supported on non-FIPS appliances** (including VPX/CPX/standard MPX models) and will fail there.
>
> This resource invokes a FIPS SIM key action. The action manipulates cryptographic key material on the FIPS HSM and is **not cleanly reversible** — destroying the resource does not undo the action performed on the appliance. Plan FIPS key transfers carefully and keep a backup of the appliance configuration.

-> **Note (one-shot action):** This is an action-only resource: applying it performs the `enable` action; it does not manage a persistent object, so re-applying re-runs the action. All attributes force replacement when changed.

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
