---
subcategory: "SSL"
---

# Resource: sslvserver\_sslcipher\_binding

This resource is used to bind a cipher, cipher group, or cipher alias to an SSL virtual server.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.44"
  port        = 443
  servicetype = "SSL"
}

resource "citrixadc_sslvserver_sslcipher_binding" "tf_binding" {
  vservername = citrixadc_lbvserver.tf_lbvserver.name
  ciphername  = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
}
```


## Argument Reference

* `vservername` - (Required) Name of the SSL virtual server. Changing this attribute forces a new resource to be created.
* `ciphername` - (Required) Name of the individual cipher, user-defined cipher group, or predefined (built-in) cipher alias. Changing this attribute forces a new resource to be created.
* `cipheraliasname` - (Optional) The name of the cipher group/alias/individual cipher bindings. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver\_sslcipher\_binding. It is the concatenation of the `vservername` and `ciphername` attributes separated by a comma.
* `description` - The cipher suite description. This is a read-only value returned by the Citrix ADC and is not configurable.


## Import

A sslvserver\_sslcipher\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslvserver_sslcipher_binding.tf_binding tf_lbvserver,TLS1.2-ECDHE-RSA-AES256-GCM-SHA384
```
