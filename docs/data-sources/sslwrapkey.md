---
subcategory: "SSL"
---

# Data Source: sslwrapkey

The sslwrapkey data source allows you to retrieve information about an SSL wrap key configured on the Citrix ADC, looked up by its name.

~> **WARNING: FIPS / crypto subsystem required.**
> This data source queries an SSL wrap key managed by the Citrix ADC cryptographic subsystem. It is intended for FIPS-capable appliances and the read may fail on platforms where the crypto subsystem is unavailable. Note that the wrap key's password and salt are secret and are never returned by the NITRO API.

## Example usage

```terraform
data "citrixadc_sslwrapkey" "example" {
  wrapkeyname = "mywrapkey"
}

output "sslwrapkey_name" {
  value = data.citrixadc_sslwrapkey.example.wrapkeyname
}
```

## Argument Reference

* `wrapkeyname` - (Required) Name for the wrap key to look up.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslwrapkey data source. It has the same value as the `wrapkeyname` attribute.
* `wrapkeyname` - Name of the wrap key.

Note: the `password` and `salt` attributes (and their write-only variants) are secret and are never returned by the NITRO API.
