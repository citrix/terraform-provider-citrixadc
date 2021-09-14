---
subcategory: "SSL"
---

# Resource: sslvserver_sslciphersuite_binding

The sslvserver_sslciphersuite_binding resource is used to add an ssl cipher suite to sslvserver.



## Example usage

```hcl
resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
	ciphername = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
	vservername = citrixadc_lbvserver.tf_sslvserver.name
}

resource "citrixadc_lbvserver" "tf_sslvserver" {
	name = "tf_sslvserver"
	servicetype = "SSL"
	ipv46 = "5.5.5.5"
	port = 80
}
```


## Argument Reference

* `ciphername` - (Required) The cipher group/alias/individual cipher configuration.
* `description` - (Optional) The cipher suite description.
* `vservername` - (Required) Name of the SSL virtual server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver_sslciphersuite_binding. It is the concatenation of the `vservername` and `ciphername` attributes separated by a comma.


## Import

A sslvserver_sslciphersuite_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding tf_sslvserver,TLS1.2-ECDHE-RSA-AES128-GCM-SHA256
```
