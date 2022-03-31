---
subcategory: "Authentication"
---

# Resource: authenticationvserver_vpnportaltheme_binding

The authenticationvserver_vpnportaltheme_binding resource is used to to bind authenticationvserver to vpnportaltheme resource.


## Example usage

```hcl
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_vpnportaltheme" "tf_vpnportaltheme" {
  name      = "tf_vpnportaltheme"
  basetheme = "X1"
}
resource "citrixadc_authenticationvserver_vpnportaltheme_binding" "tf_bind" {
  name = citrixadc_authenticationvserver.tf_authenticationvserver.name
  portaltheme = citrixadc_vpnportaltheme.tf_vpnportaltheme.name
}
```


## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `portaltheme` - (Required) Theme for Authentication virtual server Login portal


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationvserver_vpnportaltheme_binding. It is the concatenation of `name` and `portaltheme` attributes seperated by comma.


## Import

A authenticationvserver_vpnportaltheme_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_authenticationvserver_vpnportaltheme_binding.tf_bind tf_authenticationvserver,tf_vpnportaltheme
```
