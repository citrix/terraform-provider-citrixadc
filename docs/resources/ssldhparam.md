---
subcategory: "SSL"
---

# Resource: ssldhparam

The ssldhparam resource is used to configure ssl DH parameters.


## Example usage

```hcl
// Make sure the dhfile name does not exist in the target ADC
// citrixadc_sslparam does not support UPDATE operation, so to change any attributes here, first delete the dhfile, if present
resource "citrixadc_ssldhparam" "foo" {
    dhfile = "/nsconfig/ssl/tfAcc_dhfile"
    bits   = "512"
    gen    = "2"
}
```


## Argument Reference

* `dhfile` - (Required) Name of and, optionally, path to the DH key file. /nsconfig/ssl/ is the default path.
* `bits` - (Required) Size, in bits, of the DH key being generated.
* `gen` - (Optional) Random number required for generating the DH key. Required as part of the DH key generation algorithm. Possible values: [ 2, 5 ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ssldhparam. It has the same value as the `dhfile` attribute.
