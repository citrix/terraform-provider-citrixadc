---
subcategory: "SSL"
---

# Resource: sslfipssimsource_init

The sslfipssimsource_init resource drives the Secure Information Management (SIM) **init** action on the *source* Citrix ADC FIPS appliance. It is used during FIPS key transfer to initialize the source appliance for SIM by generating and registering its certificate file, which establishes the source identity used when exchanging key material with a target FIPS appliance. This is a one-shot action resource, not a persistent configuration object.

~> **WARNING: FIPS / HSM hardware required and not cleanly reversible.**
> This resource requires a dedicated FIPS appliance with an on-board Hardware Security Module. It is **not supported on non-FIPS appliances** (including VPX/CPX/standard MPX models) and will fail there.
>
> This resource invokes a FIPS SIM initialization action. The action manipulates cryptographic key material on the FIPS HSM and is **not cleanly reversible** — destroying the resource only removes it from Terraform state and does not undo the action performed on the appliance. Plan FIPS key transfers carefully and keep a backup of the appliance configuration.

-> **Note (one-shot action):** This is an action-only resource. The NITRO API exposes only the `init` action for it and provides **no GET, update, or delete endpoint**. As a result, the provider's Read is a no-op (drift cannot be detected), Update is a no-op, and Delete only removes the entry from Terraform state. All attributes force replacement when changed. Because there is no GET endpoint, no data source is published for this resource, and import is not meaningful.

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

* `id` - A synthetic identifier for this action-only resource. Because there is no GET endpoint, it is a fixed string with the value `sslfipssimsource_init`. It does not correspond to any object on the Citrix ADC.
