---
subcategory: "VPN"
---

# Resource: vpnglobal_sharefileserver_binding

The vpnglobal_sharefileserver_binding resource is used to bind sharefileserver to vpnglobal congiguration.


## Example usage

```hcl
resource "citrixadc_vpnglobal_sharefileserver_binding" "tf_bind" {
  sharefile = "3.4.5.2:8080"
}
```


## Argument Reference

* `sharefile` - (Required) Configured Sharefile server, in the format IP:PORT / FQDN:PORT
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_sharefileserver_binding. It has the same value as the `sharefile` attribute.


## Import

A vpnglobal_sharefileserver_binding can be imported using its sharefile, e.g.

```shell
terraform import citrixadc_vpnglobal_sharefileserver_binding.tf_bind 3.4.5.2:8080
```
