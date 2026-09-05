---
subcategory: "VPN"
---

# Resource: vpnvserver_vpnsecureprivateaccessprofile_binding

The vpnvserver_vpnsecureprivateaccessprofile_binding resource is used to bind a Secure Private Access profile to a vpnvserver.


## Example usage

```hcl
resource "citrixadc_vpnsecureprivateaccessprofile" "tf_spa_profile" {
  name = "tf_spa_profile"
  url  = "https://www.citrix.com"
}

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}

resource "citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding" "tf_bind" {
  name                       = citrixadc_vpnvserver.tf_vpnvserver.name
  secureprivateaccessprofile = citrixadc_vpnsecureprivateaccessprofile.tf_spa_profile.name
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `secureprivateaccessprofile` - (Required) Name of the Secure Private Access profile bound to the vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnsecureprivateaccessprofile_binding. It is the concatenation of the `name` and `secureprivateaccessprofile` attributes separated by a comma.


## Import

A vpnvserver_vpnsecureprivateaccessprofile_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind tf_vserver,tf_spa_profile
```
