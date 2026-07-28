---
subcategory: "QUIC"
---

# Resource: quicparam

This resource is used to manage the global QUIC parameters on the Citrix ADC.


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

* `id` - The ID of the quicparam resource. It is set to `quicparam-config`.


## Import

A quicparam can be imported using its id (the fixed singleton value `quicparam-config`), e.g.

```shell
terraform import citrixadc_quicparam.tf_quicparam quicparam-config
```
