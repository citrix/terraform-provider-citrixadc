---
subcategory: "sslcipher_sslciphersuite_binding"
---

# Data Source: sslcipher_sslciphersuite_binding

The sslcipher_sslciphersuite_binding data source allows you to retrieve information about a specific binding between a sslcipher group and a sslciphersuite.


## Example Usage

```terraform
data "citrixadc_sslcipher_sslciphersuite_binding" "tf_bind" {
  ciphergroupname = "tfsslcipher"
  ciphername      = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
}

output "cipherpriority" {
  value = data.citrixadc_sslcipher_sslciphersuite_binding.tf_bind.cipherpriority
}

output "description" {
  value = data.citrixadc_sslcipher_sslciphersuite_binding.tf_bind.description
}
```


## Argument Reference

* `ciphergroupname` - (Required) Name of the user-defined cipher group.
* `ciphername` - (Required) Cipher name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `cipheroperation` - The operation that is performed when adding the cipher-suite. Possible cipher operations are: ADD - Appends the given cipher-suite to the existing one configured for the virtual server. REM - Removes the given cipher-suite from the existing one configured for the virtual server. ORD - Overrides the current configured cipher-suite for the virtual server with the given cipher-suite.
* `cipherpriority` - This indicates priority assigned to the particular cipher.
* `ciphgrpals` - A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or user defined cipher-group name.
* `description` - Cipher suite description.
* `id` - The id of the sslcipher_sslciphersuite_binding. It is the concatenation of both `ciphergroupname` and `ciphername` attributes separated by comma.
