---
subcategory: "Authentication"
---

# Resource: authenticationauthnprofile

The authenticationauthnprofile resource is used to create Authentication profile resource.


## Example usage

```hcl
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new_vserver"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationauthnprofile" "tf_authenticationauthnprofile" {
  name                = "tf_name"
  authnvsname         = citrixadc_authenticationvserver.tf_authenticationvserver.name
  authenticationhost  = "hostname"
  authenticationlevel = "20"
}
```


## Argument Reference

* `name` - (Required) Name for the authentication profile.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the RADIUS action is added.
* `authnvsname` - (Required) Name of the authentication vserver at which authentication should be done.
* `authenticationdomain` - (Optional) Domain for which TM cookie must to be set. If unspecified, cookie will be set for FQDN.
* `authenticationhost` - (Optional) Hostname of the authentication vserver to which user must be redirected for authentication.
* `authenticationlevel` - (Optional) Authentication weight or level of the vserver to which this will bound. This is used to order TM vservers based on the protection required. A session that is created by authenticating against TM vserver at given level cannot be used to access TM vserver at a higher level.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationauthnprofile. It has the same value as the `name` attribute.


## Import

A authenticationauthnprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile tf_name
```
