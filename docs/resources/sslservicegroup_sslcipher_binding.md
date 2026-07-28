---
subcategory: "SSL"
---

# Resource: sslservicegroup_sslcipher_binding

This resource is used to bind an SSL cipher to an SSL service group.


## Example usage

```hcl
resource "citrixadc_sslservicegroup_sslcipher_binding" "tf_binding" {
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
  ciphername       = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "SSL"
}
```


## Argument Reference

* `servicegroupname` - (Required) The name of the SSL service group to which the cipher is bound. Changing this forces a new resource to be created.
* `ciphername` - (Required) A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or a user defined cipher-group name. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following read-only attributes are available:

* `id` - The id of the sslservicegroup_sslcipher_binding. It is the concatenation of the `servicegroupname` and `ciphername` attributes separated by a comma.
* `cipheraliasname` - The name of the cipher group/alias/name configured for the SSL service group. This is a read-only (computed) value returned by the NITRO server; the bind endpoint does not accept it.
* `description` - The description of the cipher. This is a read-only (computed) value returned by the NITRO server; the bind endpoint does not accept it.


## Import

A sslservicegroup_sslcipher_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslservicegroup_sslcipher_binding.tf_binding tf_servicegroup,TLS1.2-ECDHE-RSA-AES256-GCM-SHA384
```
