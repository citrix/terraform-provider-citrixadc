---
subcategory: "VPN"
---

# Resource: vpnglobal_staserver_binding

The vpnglobal_staserver_binding resource is used to bind staserver to vpnglobal configuration.


## Example usage

```hcl
resource "citrixadc_vpnglobal_staserver_binding" "tf_bind" {
  staserver      = "http://www.example.com/"
  staaddresstype = "IPV4"
}
```


## Argument Reference

* `staserver` - (Required) Configured Secure Ticketing Authority (STA) server.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `staaddresstype` - (Optional) Type of the STA server address(ipv4/v6).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_staserver_binding. It has the same value as the `staserver` attribute.


## Import

A vpnglobal_staserver_binding can be imported using its staserver, e.g.

```shell
terraform import citrixadc_vpnglobal_staserver_binding.tf_bind http://www.example.com/
```
