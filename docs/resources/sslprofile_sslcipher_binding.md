---
subcategory: "SSL"
---

# Resource: sslprofile_sslcipher_binding

The sslprofile_sslcipher_binding resource is used to create bindings between sslprofiles and sslciphers.

~> If you are using this resource to bind sslciphers to a sslprofile
do not define the `cipherbindings` attribute in the sslprofile resource.


## Example usage

```hcl
resource "citrixadc_sslprofile_sslcipher_binding" "tf_binding" {
    name = "tf_sslprofile"
    ciphername = "HIGH"
    cipherpriority = 10
}
```


## Argument Reference

* `name` - (Required) Name of the SSL profile.
* `ciphername` - (Required) Name of the cipher.
* `cipherpriority` - (Optional) Cipher priority.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile_sslcipher_binding. It is the concatenation of the `name` and `ciphername` attributes separated by a comma.


## Import

A sslprofile_sslcipher_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslprofile_sslcipher_binding.tf_binding tf_sslprofile,HIGH
```
