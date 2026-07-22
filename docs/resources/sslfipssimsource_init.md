---
subcategory: "SSL"
---

# Resource: sslfipssimsource_init

The sslfipssimsource_init resource drives the Secure Information Management (SIM) **init** action on the *source* Citrix ADC FIPS appliance. It is used during FIPS key transfer to initialize the source appliance for SIM by generating and registering its certificate file, which establishes the source identity used when exchanging key material with a target FIPS appliance. This is a one-shot action resource, not a persistent configuration object.

~> **WARNING: FIPS / HSM hardware required and not cleanly reversible.**
> This resource requires a dedicated FIPS appliance with an on-board Hardware Security Module. It is **not supported on non-FIPS appliances** (including VPX/CPX/standard MPX models) and will fail there.
>
> This resource invokes a FIPS SIM initialization action. The action manipulates cryptographic key material on the FIPS HSM and is **not cleanly reversible** — destroying the resource does not undo the action performed on the appliance. Plan FIPS key transfers carefully and keep a backup of the appliance configuration.

-> **Note (one-shot action):** This is an action-only resource: applying it performs the `init` action; it does not manage a persistent object, so re-applying re-runs the action. All attributes force replacement when changed.

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
