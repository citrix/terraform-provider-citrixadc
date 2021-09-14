---
subcategory: "Content Switching"
---

# Resource: csvserver_vpnvserver_binding

The csvserver_vpnvserver_binding resource is used to bind a vpnvserver to csvserver.


## Example usage

```hcl
resource "citrixadc_csvserver_vpnvserver_binding" "tf_csvserver_vpnvserver_binding" {
	name = "tf_csvserver"
	vserver = "tf_vpnvserver"
}
```


## Argument Reference

* `vserver` - (Required) Name of the default gslb or vpn vserver bound to CS vserver of type GSLB/VPN. For Example: bind cs vserver cs1 -vserver gslb1 or bind cs vserver cs1 -vserver vpn1.
* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_vpnvserver_binding. It is the concatenation of the `name` and `vserver` attributes separated by a comma.


## Import

A csvserver_vpnvserver_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding tf_csvserver,tf_vpnvserver
```
