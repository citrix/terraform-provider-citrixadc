---
subcategory: "SSL"
---

# Resource: sslcipher

The sslcipher resource is used to create ssl ciphers.


## Example usage

```hcl
resource "citrixadc_sslcipher" "tfsslcipher" {
  ciphergroupname = "tfsslcipher"
}
```


## Argument Reference

* `ciphergroupname` - (Required) Name of the cipher group to be created.
* `ciphersuitebinding` - (Optional) A set of ciphersuites bound to this cipher group. Any change to this set will recreate the whole cipher group. Attributes documented below. (Deprecates soon)

!>
[**DEPRECATED**] Please use `sslcipher_sslciphersuite_binding` to bind `sslciphersuite` to `sslcipher` insted of this resource. The support for binding `sslciphersuite` to `sslcipher` in `sslcipher` resource will get deprecated soon.


A ciphersuitebinding supports the following:

* `ciphername` - (Required) Cipher name.
* `cipherpriority` - (Optional) This indicates priority assigned to the particular cipher.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcipher. It has the same value as the `ciphergroupname` attribute.


## Import

A sslcipher can be imported using its ciphergroupname, e.g.

```shell
terraform import citrixadc_sslcipher.tf_sslcipher tf_sslcipher
```
