---
subcategory: "Load Balancing"
---

# Resource: lbgroup_lbvserver_binding

The lbgroup_lbvserver_binding resource is used to add a lbvserver to lbgroup.


## Example usage

```hcl
resource "citrixadc_lbgroup_lbvserver_binding" "tf_lbvserverbinding" {
	name = citrixadc_lbgroup.tf_lbgroup.name
	vservername = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name 	    = "tf_lbvserver"
	ipv46       = "1.1.1.8"
	port        = "80"
	servicetype = "HTTP"
}

resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup"
}
```


## Argument Reference

* `vservername` - (Required) Virtual server name.
* `name` - (Required) Name for the load balancing virtual server group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lbgroup" or 'my lbgroup').


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbgroup_lbvserver_binding. It is the concatenation of the `name` and `vservername` attributes separated by a comma.


## Import

A lbgroup_lbvserver_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding tf_lbgroup,tf_lbvserver
```
