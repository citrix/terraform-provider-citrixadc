---
subcategory: "SSL"
---

# Data Source: citrixadc_sslcipher

The sslcipher data source allows you to retrieve information about user-defined SSL cipher groups.

## Example usage

```terraform
data "citrixadc_sslcipher" "tf_sslcipher" {
  ciphergroupname = "tfAccsslcipher"
}

output "ciphergroupname" {
  value = data.citrixadc_sslcipher.tf_sslcipher.ciphergroupname
}

output "ciphername" {
  value = data.citrixadc_sslcipher.tf_sslcipher.ciphername
}
```

## Argument Reference

* `ciphergroupname` - (Required) Name for the user-defined cipher group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the cipher group is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `ciphername` - Cipher name.
* `cipherpriority` - This indicates priority assigned to the particular cipher.
* `ciphgrpalias` - The individual cipher name(s), a user-defined cipher group, or a system predefined cipher alias that will be added to the predefined cipher alias that will be added to the group cipherGroupName. If a cipher alias or a cipher group is specified, all the individual ciphers in the cipher alias or group will be added to the user-defined cipher group.
* `id` - The id of the sslcipher. It has the same value as the `ciphergroupname` attribute.
* `sslprofile` - Name of the profile to which cipher is attached.

## Import

A sslcipher can be imported using its ciphergroupname, e.g.

```shell
terraform import citrixadc_sslcipher.tf_sslcipher tfAccsslcipher
```
