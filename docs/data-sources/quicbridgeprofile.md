---
subcategory: "NS"
---

# Data Source: citrixadc_quicbridgeprofile

The `citrixadc_quicbridgeprofile` data source is used to retrieve information about a specific QUIC bridge profile configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_quicbridgeprofile" "example" {
  name = "my_quicbridgeprofile"
}
```

## Argument Reference

* `name` - (Required) Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals sign (=), and hyphen (-) characters.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the QUIC bridge profile.
* `routingalgorithm` - Routing algorithm to generate routable connection IDs.
* `serveridlength` - Length of serverid to encode/decode server information.
