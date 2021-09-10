---
subcategory: "Content Switching"
---

# Resource: csvserver_gslbvserver_binding

The csvserver_gslbvserver_binding resource is used to bind a gslb vserver to csvserver.


## Example usage

```hcl
resource "citrixadc_csvserver_gslbvserver_binding" "tf_csvserver_gslbvserver_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	vserver = citrixadc_gslbvserver.tf_gslbvserver.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	servicetype = "HTTP"
	targettype = "GSLB"
}

resource "citrixadc_gslbvserver" "tf_gslbvserver" {
	name = "tf_gslbvserver"
	servicetype = "HTTP"
}
```


## Argument Reference

* `vserver` - (Required) Name of the default gslb or vpn vserver bound to CS vserver of type GSLB/VPN. For Example: bind cs vserver cs1 -vserver gslb1 or bind cs vserver cs1 -vserver vpn1.
* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_gslbvserver_binding. It is the concatenation of the `name` and `vserver` attributes separated by a comma.


## Import

A csvserver_gslbvserver_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding tf_csvserver,tf_gslbvserver
```
