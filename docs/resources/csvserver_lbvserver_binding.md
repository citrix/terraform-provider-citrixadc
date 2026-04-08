---
subcategory: "Content Switching"
---

# Resource: csvserver_lbvserver_binding

The csvserver_lbvserver_binding resource is used to bind a load balancing vserver to a content switching vserver.


## Example usage

```hcl
resource "citrixadc_csvserver_lbvserver_binding" "tf_csvserver_lbvserver_binding" {
	name      = citrixadc_csvserver.tf_csvserver.name
	lbvserver = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name        = "tf_csvserver"
	servicetype = "HTTP"
	ipv46       = "10.10.10.10"
	port        = 80
	lifecycle {
		ignore_changes = [lbvserverbinding]
	}
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver"
	servicetype = "HTTP"
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `lbvserver` - (Required) Name of the default lb vserver bound. Use this param for Default binding only. For Example: bind cs vserver cs1 -lbvserver lb1.
* `targetvserver` - (Optional) The virtual server name (created with the add lb vserver command) to which content will be switched.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_lbvserver_binding. It is the concatenation of the `name` and `lbvserver` attributes in the format `name,lbvserver`.


## Import

A csvserver_lbvserver_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding tf_csvserver,tf_lbvserver
```
