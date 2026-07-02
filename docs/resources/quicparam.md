---
subcategory: "QUIC"
---

# Resource: quicparam

Configures the global QUIC parameters on the Citrix ADC. QUIC uses address validation tokens (carried in Retry packets and NEW_TOKEN frames) to verify a client's source address; this resource lets you control how frequently the secret used to generate those tokens is rotated, so you can balance token security against the overhead of rotation.

This is a singleton settings resource: the QUIC parameters always exist on the appliance, so creating this resource updates the existing configuration and deleting it simply removes the resource from Terraform state without changing the appliance.


## Example usage

```hcl
resource "citrixadc_quicparam" "tf_quicparam" {
  quicsecrettimeout = 3600
}
```

### Disabling secret rotation

Set `quicsecrettimeout` to `0` when secret rotation is not desired:

```hcl
resource "citrixadc_quicparam" "tf_quicparam" {
  quicsecrettimeout = 0
}
```


## Argument Reference

* `quicsecrettimeout` - (Optional) Rotation frequency, in seconds, for the secret used to generate address validation tokens that will be issued in QUIC Retry packets and QUIC NEW_TOKEN frames sent by the Citrix ADC. A value of `0` can be configured if secret rotation is not desired. Defaults to `3600`. Minimum value = `0` Maximum value = `31536000`


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the quicparam resource. Because this is a singleton, it is a fixed string with the value `"quicparam-config"`.
