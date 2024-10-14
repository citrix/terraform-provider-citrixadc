---
subcategory: "sslcipher_sslciphersuite_binding"
---

# Resource: sslcipher_sslciphersuite_binding

The sslcipher_sslciphersuite_binding resource is used to bind sslciphersuite to sslcipher.


## Example usage

```hcl

resource "citrixadc_sslcipher" "tfsslcipher" {
  ciphergroupname = "tfsslcipher"
}

resource "citrixadc_sslcipher_sslciphersuite_binding" "tf_bind" {
  ciphergroupname = citrixadc_sslcipher.tfsslcipher.ciphergroupname
  ciphername      = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
  cipherpriority  = 1
}

```


## Argument Reference

* `ciphername` - (Required) Cipher name.
* `ciphergroupname` - (Required) Name for the user-defined cipher group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the cipher group is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ciphergroup" or 'my ciphergroup'). Minimum length =  1
* `description` - (Optional) Cipher suite description.
* `cipherpriority` - (Optional) This indicates priority assigned to the particular cipher. Minimum value =  1
* `cipheroperation` - (Optional) The operation that is performed when adding the cipher-suite.  Possible cipher operations are: 	ADD - Appends the given cipher-suite to the existing one configured for the virtual server. 	REM - Removes the given cipher-suite from the existing one configured for the virtual server. 	ORD - Overrides the current configured cipher-suite for the virtual server with the given cipher-suite. Possible values: [ ADD, REM, ORD ]
* `ciphgrpals` - (Optional) A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or user defined cipher-group name. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcipher_sslciphersuite_binding. It is the concatenation of both `ciphergroupname` and `ciphername` attributes seperated by comma. Ex: `tfsslcipher,TLS1.2-ECDHE-RSA-AES128-GCM-SHA256` is `id` for above example.


## Import

A sslcipher_sslciphersuite_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslcipher_sslciphersuite_binding.tf_bind tfsslcipher,TLS1.2-ECDHE-RSA-AES128-GCM-SHA256
```
