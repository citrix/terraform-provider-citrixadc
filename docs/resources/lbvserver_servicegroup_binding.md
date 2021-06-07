---
subcategory: "Load Balancing"
---

# Resource: lbvserver\_servicegroup\_binding

The lbvserver\_servicegroup\_binding resource is used to bind servicegroups to lb vservers.

If a binding between lbvserver and servicegroup is set this way the `lbvservers` option
of `resource_citrixadc_servicegroup` should not be set for the same servicegroup.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
    servicegroupname = "tf_servicegroup"
    servicetype  = "HTTP"
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}
```


## Argument Reference

* `servicegroupname` - (Required) The service group name bound to the selected load balancing virtual server.
* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver\_servicegroup\_binding. It is the concatenation of the name and servicegroupname attributes separated by a comma.


## Import

A lbvserver\_servicegroup\_binding can be imported using its id.

```shell
terraform import citrixadc_lbvserver_servicegroup_binding.tf_binding tf_lbvserver,tf_servicegroupname
```
