---
subcategory: "SSL"
---

# Resource: sslprofile_sslcipher_binding

The sslprofile_sslcipher_binding resource is used to create bindings between sslprofiles and sslciphers.

~> If you are using this resource to bind sslciphers to a sslprofile
do not define the `cipherbindings` attribute in the sslprofile resource.

~> The attribute `remove_existing_sslcipher_binding` should be `true` if you want to delete all the existing bindings, and bind the new sslcipher to sslprofile.


## Example usage

```hcl
resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tf_sslprofile"
}

resource "citrixadc_sslprofile_sslcipher_binding" "tf_binding" {
    name                              = citrixadc_sslprofile.tf_sslprofile.name
    ciphername                        = "HIGH"
    cipherpriority                    = 10
    remove_existing_sslcipher_binding = true
}
```


## Argument Reference

* `ciphername` - (Required) Name of the cipher.
* `name` - (Required) Name of the SSL profile.
* `cipherpriority` - (Optional) cipher priority.
* `remove_existing_sslcipher_binding` - (Optional) If you want to unbind all the existing sslcipher bindings, then set this as true else false. Default is `false`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile_sslcipher_binding. It has is the conatenation of the `name` and `ciphername` attributes.


## Import

A sslprofile_sslcipher_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslprofile_sslcipher_binding.tf_binding tf_sslprofile,HIGH
```
