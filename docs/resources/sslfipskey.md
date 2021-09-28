---
subcategory: "SSL"
---

# Resource: sslfipskey

The sslfipskey resource is used to create SSL fips key.


## Example usage

```hcl
resource "citrixadc_sslfipskey" "demo_sslfipskey" {
    fipskeyname = "f1"
    keytype = "ECDSA"
    curve = "P_256"
}
```


## Argument Reference

* `fipskeyname` - (Required) Name for the FIPS key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the FIPS key is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my fipskey" or 'my fipskey').
* `keytype` - (Required) Only RSA key and ECDSA Key are supported. Possible values: [ RSA, ECDSA ]
* `exponent` - (Optional) Exponent value for the FIPS key to be created. Available values function as follows: 3=3 (hexadecimal) F4=10001 (hexadecimal). Possible values: [ 3, F4 ]
* `modulus` - (Optional) Modulus, in multiples of 64, of the FIPS key to be created.
* `curve` - (Optional) Only p_256 (prime256v1) and P_384 (secp384r1) are supported. Possible values: [ P_256, P_384 ]
* `key` - (Optional) Name of and, optionally, path to the key file to be imported. /nsconfig/ssl/ is the default path.
* `inform` - (Optional) Input format of the key file. Available formats are: SIM - Secure Information Management; select when importing a FIPS key. If the external FIPS key is encrypted, first decrypt it, and then import it. PEM - Privacy Enhanced Mail; select when importing a non-FIPS key. Possible values: [ SIM, DER, PEM ]
* `wrapkeyname` - (Optional) Name of the wrap key to use for importing the key. Required for importing a non-FIPS key.
* `iv` - (Optional) Initialization Vector (IV) to use for importing the key. Required for importing a non-FIPS key.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslfipskey. It has the same value as the `fipskeyname` attribute.


## Import

A sslfipskey can be imported using its name, e.g.

```shell
terraform import citrixadc_sslfipskey.tf_sslfipskey tf_sslfipskey
```
