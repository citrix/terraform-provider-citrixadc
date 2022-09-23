---
subcategory: "SSL"
---

# Resource: sslcipher

The sslcipher resource is used to create ssl ciphers.


## Example usage

```hcl
resource "citrixadc_sslcipher" "tfsslcipher" {
  ciphergroupname = "tfsslcipher"

  # ciphersuitebinding is MANDATORY attribute
  # Any change in the ciphersuitebinding will result in re-creation of the whole sslcipher resource.
  ciphersuitebinding {
    ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
    cipherpriority = 1
  }
  ciphersuitebinding {
    ciphername     = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
    cipherpriority = 2
  }
  ciphersuitebinding {
    ciphername     = "TLS1.2-ECDHE-RSA-AES-128-SHA256"
    cipherpriority = 3
  }
}
```


## Argument Reference

* `ciphergroupname` - (Required) Name of the cipher group to be created.
* `ciphersuitebinding` - (Required) A set of ciphersuites bound to this cipher group. Any change to this set will recreate the whole cipher group. Attributes documented below.

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
