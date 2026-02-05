---
subcategory: "SSL"
---

# Data Source: sslfipskey

The sslfipskey data source allows you to retrieve information about SSL FIPS keys.

## Example usage

```terraform
data "citrixadc_sslfipskey" "demo_sslfipskey" {
  fipskeyname = "f1"
}

output "keytype" {
  value = data.citrixadc_sslfipskey.demo_sslfipskey.keytype
}

output "curve" {
  value = data.citrixadc_sslfipskey.demo_sslfipskey.curve
}
```

## Argument Reference

* `fipskeyname` - (Required) Name for the FIPS key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the FIPS key is created.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my fipskey" or 'my fipskey').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `curve` - Only p_256 (prime256v1) and P_384 (secp384r1) are supported.
* `exponent` - Exponent value for the FIPS key to be created. Available values function as follows:
 3=3 (hexadecimal)
F4=10001 (hexadecimal)
* `inform` - Input format of the key file. Available formats are:
SIM - Secure Information Management; select when importing a FIPS key. If the external FIPS key is encrypted, first decrypt it, and then import it.
PEM - Privacy Enhanced Mail; select when importing a non-FIPS key.
* `iv` - Initialization Vector (IV) to use for importing the key. Required for importing a non-FIPS key.
* `key` - Name of and, optionally, path to the key file to be imported.
 /nsconfig/ssl/ is the default path.
* `keytype` - Only RSA key and ECDSA Key are supported.
* `modulus` - Modulus, in multiples of 64, of the FIPS key to be created.
* `wrapkeyname` - Name of the wrap key to use for importing the key. Required for importing a non-FIPS key.

## Attribute Reference

* `id` - The id of the sslfipskey. It has the same value as the `fipskeyname` attribute.

## Import

A sslfipskey can be imported using its fipskeyname, e.g.

```shell
terraform import citrixadc_sslfipskey.demo_sslfipskey f1
```
