---
subcategory: "QUIC"
---

# Data Source: quicparam

The quicparam data source allows you to retrieve the global QUIC parameters configured on the Citrix ADC, such as the rotation frequency of the secret used to generate QUIC address validation tokens.

Because the QUIC parameters are a singleton configuration on the appliance, no lookup attribute is required.


## Example usage

```terraform
data "citrixadc_quicparam" "example" {
}

output "quic_secret_timeout" {
  value = data.citrixadc_quicparam.example.quicsecrettimeout
}
```


## Argument Reference

This data source has no configurable arguments. The QUIC parameters form a single global configuration that is read directly from the appliance.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the quicparam data source.
* `quicsecrettimeout` - Rotation frequency, in seconds, for the secret used to generate address validation tokens that will be issued in QUIC Retry packets and QUIC NEW_TOKEN frames sent by the Citrix ADC. A value of `0` indicates that secret rotation is disabled.
