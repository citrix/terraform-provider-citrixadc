---
subcategory: "VPN"
---

# Resource: vpnvserver_vpnepaprofile_binding

Binds an advanced EPA (Endpoint Analysis) profile to a VPN virtual server on the Citrix ADC. Create this binding when you want a preauthentication EPA profile to be evaluated for connections handled by the VPN virtual server.

This binding is immutable: it can only be created (bound) or deleted (unbound). Changing any attribute forces Terraform to replace the resource.

~> **NOTE:** Current NetScaler firmware rejects `bind vpn vserver -epaprofile` with the error "There has been a design change in the support of OPSWAT specific EPA scans. EPA Profile Configuration is no longer needed." This binding is therefore non-functional on current firmware and a live apply will fail. The resource is retained for older firmware and state compatibility.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf.citrix.example.com"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}

resource "citrixadc_vpnepaprofile" "tf_vpnepaprofile" {
  name     = "tf_vpnepaprofile"
  filename = "tf_vpnepaprofile.xml"
}

resource "citrixadc_vpnvserver_vpnepaprofile_binding" "tf_binding" {
  name               = citrixadc_vpnvserver.tf_vpnvserver.name
  epaprofile         = citrixadc_vpnepaprofile.tf_vpnepaprofile.name
  epaprofileoptional = true
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server. Minimum length = 1. Changing this attribute forces a new resource to be created.
* `epaprofile` - (Required) Advanced EPA profile to bind. Changing this attribute forces a new resource to be created.
* `epaprofileoptional` - (Optional) Mark the EPA profile optional for preauthentication EPA profile. User would be shown a logon page even if the EPA profile fails to evaluate. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnepaprofile_binding. It is a system-generated identifier built from the unique attributes as a comma-separated list of `key:value` pairs (the values are URL-encoded), in the form `epaprofile:<epaprofile>,name:<name>`.


## Import

A vpnvserver_vpnepaprofile_binding can be imported using its id, which is the comma-separated list of `key:value` pairs described above, e.g.

```shell
terraform import citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding "epaprofile:tf_vpnepaprofile,name:tf.citrix.example.com"
```
