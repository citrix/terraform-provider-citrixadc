---
subcategory: "Metrics"
---

# Resource: metricsprofile_vpnvserver_binding

Associates a VPN virtual server with a metrics profile so that the metrics defined in the profile are collected for that vserver. Create one binding for each vpnvserver you want the metrics profile to monitor.


## Example usage

```hcl
resource "citrixadc_metricsprofile_vpnvserver_binding" "tf_binding" {
  name       = "http_metrics"
  entityname = "vpn_vserver1"
  entitytype = "vpnvserver"
}
```


## Argument Reference

* `name` - (Required) Name of the metrics profile to which the entity is bound. Must begin with an ASCII alphabetic or underscore (`_`) character, and must contain only ASCII alphanumeric, underscore, hash (`#`), period (`.`), space, colon (`:`), at (`@`), equals (`=`), and hyphen (`-`) characters. Changing this value forces a new resource to be created.
* `entityname` - (Required) Name of the entity (VPN virtual server) bound to the metrics profile. Changing this value forces a new resource to be created.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is always `vpnvserver`. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the `metricsprofile_vpnvserver_binding` resource. It is a composite of the unique attributes expressed as comma-separated `key:value` pairs, in the form `entityname:<entityname>,entitytype:<entitytype>,name:<name>` (each value is URL-encoded). For example: `entityname:vpn_vserver1,entitytype:vpnvserver,name:http_metrics`.


## Import

A metricsprofile_vpnvserver_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_metricsprofile_vpnvserver_binding.tf_binding "entityname:vpn_vserver1,entitytype:vpnvserver,name:http_metrics"
```
